package auth

import (
	"auth/internal/domain/models"
	"auth/internal/lib/jwt"
	"auth/internal/lib/logger/sl"
	"auth/internal/storage"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"strings"
	"time"
)

type Auth struct {
	log             *slog.Logger
	usrSaver        UserSaver
	usrProvider     UserProvider
	appProvider     AppProvider
	refTokenManager RefreshTokenManager
	tokenTTL        time.Duration
	refreshTTL      time.Duration
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppId       = errors.New("invalid identifier")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

var testAppId = 1

type UserSaver interface {
	SaveUser(ctx context.Context, id uuid.UUID, name, lastName, email, goal string, passHash []byte) error
}

type RefreshTokenManager interface {
	SaveRefreshToken(ctx context.Context, tokenID, userID string, token string, expiresAt time.Time) error
	DeleteRefreshToken(ctx context.Context, tokenID string) error
	UpdateToken(ctx context.Context, tokenID, userID string, expiresAt time.Time) error
	FindRefreshByUserID(ctx context.Context, userID string) (string, error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	FindUserByRefreshTokenID(ctx context.Context, tokenID string) (models.User, error)
	IsAdmin(ctx context.Context, userID string) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// sNew возвращает новый экземпляр службы Auth
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	refreshTokenManager RefreshTokenManager,
	tokenTTL time.Duration,
	refreshTTL time.Duration,
) *Auth {
	return &Auth{
		usrSaver:        userSaver,
		usrProvider:     userProvider,
		log:             log,
		appProvider:     appProvider,
		refTokenManager: refreshTokenManager,
		tokenTTL:        tokenTTL,
		refreshTTL:      refreshTTL,
	}
}

// Login проверяет, существует ли пользователь с указанными учетными данными в системе, и возвращает access token и refresh token.
//
// Если пользователь существует, но пароль неверен, возвращается ошибка.
// Если пользователь не существует, возвращает ошибку.
func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appID int,
) (string, string, error) {
	const op = "Auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email),
	)

	log.Info("attempting to login user")

	// Ищем user
	user, err := a.usrProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user logged in successfully")

	// Генерируем новый Access token
	token, err := jwt.NewToken(user, app, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	// Делаем проверку нет ли старого токена
	oldRefreshToken, err := a.refTokenManager.FindRefreshByUserID(ctx, user.ID.String())
	if err != nil {
		a.log.Error("failed to find token", sl.Err(err))
	}
	if oldRefreshToken != "" {
		// Удаляем старый refresh token, в случае если он есть
		err = a.refTokenManager.DeleteRefreshToken(ctx, oldRefreshToken)
		if err != nil {
			a.log.Warn("failed to delete old token", sl.Err(err))

			return "", "", fmt.Errorf("%s: %w", op, err)
		}
	}

	// Генерируем новый refresh token
	refreshToken, err := jwt.GenerateRefresh(user, app, a.refreshTTL)
	if err != nil {
		a.log.Error("failed to generate refresh token", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	// Создаем уникальный uuid для нашего refresh токена
	idRef := uuid.New()

	// Сохранение refresh токена в базу данных
	err = a.refTokenManager.SaveRefreshToken(ctx, idRef.String(), user.ID.String(), refreshToken, time.Now().Add(a.refreshTTL))
	if err != nil {
		a.log.Error("failed to save refresh token", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("successfully logged in")

	return token, refreshToken, nil
}

// Register регистрирует нового пользователя в системе и возвращает идентификатор пользователя.
// Если пользователь с данным именем пользователя уже существует, возвращается ошибка.
func (a *Auth) Register(ctx context.Context, name, lastName, email, goal string, pass string) (uuid.UUID, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("register user")

	id := uuid.New()

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	err = a.usrSaver.SaveUser(ctx, id, name, lastName, email, goal, passHash)
	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "user already exists") {
			log.Warn("user already exists", sl.Err(err))

			return uuid.Nil, fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("failed to save user", sl.Err(err))

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registred")

	return id, nil
}

// IsAdmin проверяет является ли пользователь администратором.
func (a *Auth) IsAdmin(ctx context.Context, userID string) (bool, error) {
	const op = "auth.IsAdmin"

	log := a.log.With(
		slog.String("op", op),
		slog.String("user_id", userID))

	log.Info("checking if user is admin")

	isAdmin, err := a.usrProvider.IsAdmin(ctx, userID)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			log.Warn("app not found", sl.Err(err))

			return false, fmt.Errorf("%s: %w", op, ErrInvalidAppId)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checking if user is admin", slog.Bool("isAdmin", isAdmin))

	return isAdmin, nil
}

// Logout удаляет refresh token
func (a *Auth) Logout(ctx context.Context, accessToken string) error {
	const op = "auth.Logout"

	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("logout user")

	app, err := a.appProvider.App(ctx, testAppId)

	claims, err := jwt.ParseAccessToken(accessToken, app)
	if err != nil {
		log.Error("failed to parse access token", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	user, err := a.usrProvider.User(ctx, claims.Email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, err := a.refTokenManager.FindRefreshByUserID(ctx, user.ID.String())
	if err != nil {
		a.log.Error("failed to find token", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	} else if refreshToken == "" {
		a.log.Warn("user not found", sl.Err(storage.ErrUserNotFound))
		return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	err = a.refTokenManager.DeleteRefreshToken(ctx, refreshToken)
	if err != nil {
		//a.log.Error("failed to delete refresh token", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("refresh token deleted")

	return nil
}

// UpdateTokens Обновляет access и refresh токены
func (a *Auth) UpdateTokens(ctx context.Context, refreshTokenID string, appID int) (string, string, error) {
	const op = "auth.RefreshToken"

	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("updating tokens")

	user, err := a.usrProvider.FindUserByRefreshTokenID(ctx, refreshTokenID)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	// Генерируем новый Access token
	token, err := jwt.NewToken(user, app, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	// Генерируем новый refresh token
	refreshToken, err := jwt.GenerateRefresh(user, app, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate refresh token", sl.Err(err))

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	// Обновляем refresh token и expiresAt
	err = a.refTokenManager.UpdateToken(ctx, refreshToken, user.ID.String(), time.Now().Add(a.refreshTTL))
	if err != nil {
		a.log.Error("failed to update refresh token", sl.Err(err))
	}

	log.Info("successfully update token")

	return token, refreshToken, nil
}

func (a *Auth) Session(ctx context.Context, accessToken string) (string, string, string, string, bool, error) {
	const op = "auth.Session"

	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("getting session")

	app, err := a.appProvider.App(ctx, testAppId)

	var claims *models.Claims

	claims, err = jwt.ParseAccessToken(accessToken, app)
	if err != nil {
		log.Error("failed to parse access token", sl.Err(err))
		return "", "", "", "", false, fmt.Errorf("%s: %w", op, err)
	}

	user, err := a.usrProvider.User(ctx, claims.Email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", "", "", "", false, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return "", "", "", "", false, fmt.Errorf("%s: %w", op, err)
	}

	isAdmin, err := a.IsAdmin(ctx, claims.ID.String())
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			a.log.Warn("app not found", sl.Err(err))
			return "", "", "", "", false, fmt.Errorf("%s: %w", op, ErrInvalidAppId)
		}
		a.log.Error("failed to get user", sl.Err(err))
		return "", "", "", "", false, fmt.Errorf("%s: %w", op, err)
	}

	return user.ID.String(), user.Name, user.LastName, user.Goal.String(), isAdmin, nil
}
