package postgres

import (
	"auth/internal/domain/models"
	"auth/internal/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

// SaveUser сохраняет пользователя в базе данных PostgreSQL.
func (s *Storage) SaveUser(ctx context.Context, id uuid.UUID, name, lastName, email, goal string, passHash []byte) error {
	const op = "storage.postgres.SaveUser"

	// Подготавливаем SQL-запрос с использованием плейсхолдеров PostgreSQL ($1, $2)
	query := `
		INSERT INTO users(id, name, last_name, email, goal, pass_hash)
		VALUES($1, $2, $3, $4, $5, $6)
	`

	// Выполняем запрос и сканируем результат в переменную id
	_, err := s.db.ExecContext(ctx, query, id, name, lastName, email, goal, passHash)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key`) { // Код ошибки для нарушения уникальности
			return fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// User возвращает пользователя по email из базы данных PostgreSQL.
func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.postgres.User"

	// SQL-запрос с использованием плейсхолдеров PostgreSQL ($1)
	query := `
		SELECT id, name, last_name, email, goal, pass_hash
		FROM users
		WHERE email = $1
	`

	// Выполняем запрос и сканируем результат в структуру models.User
	var user models.User
	err := s.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Goal, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// IsAdmin проверяет, является ли пользователь администратором.
func (s *Storage) IsAdmin(ctx context.Context, userID string) (bool, error) {
	const op = "storage.postgres.IsAdmin"

	query := `
		SELECT is_admin
		FROM admins
		WHERE user_id = $1
	`

	// Выполняем запрос и сканируем результат в переменную isAdmin
	var isAdmin bool
	err := s.db.QueryRowContext(ctx, query, userID).Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}

// App возвращает приложение по id из базы данных PostgreSQL.
func (s *Storage) App(ctx context.Context, id int) (models.App, error) {
	const op = "storage.postgres.App"

	query := `
		SELECT id, name, secret
		FROM apps
		WHERE id = $1
	`

	// Выполняем запрос и сканируем результат в структуру models.App
	var app models.App
	err := s.db.QueryRowContext(ctx, query, id).Scan(&app.ID, &app.Name, &app.Secret)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}

		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}

// SaveRefreshToken сохраняет refresh token в бд
func (s *Storage) SaveRefreshToken(ctx context.Context, tokenID, userID string, token string, expiresAt time.Time) error {
	const op = "storage.postgres.SaveRefreshToken"
	query := `
		INSERT INTO refresh_tokens (id, user_id, token, expires_at) 
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.ExecContext(ctx, query, tokenID, userID, token, expiresAt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteRefreshToken(ctx context.Context, tokenID string) error {
	const op = "storage.postgres.DeleteRefreshToken"

	query := `
		DELETE FROM refresh_tokens WHERE token = $1
	`
	_, err := s.db.ExecContext(ctx, query, tokenID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) FindRefreshByUserID(ctx context.Context, userID string) (string, error) {
	const op = "storage.postgres.FindRefreshByUserID"

	query := `
		SELECT token FROM refresh_tokens WHERE user_id = $1
	`
	token := ""
	err := s.db.QueryRowContext(ctx, query, userID).Scan(&token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (s *Storage) FindUserByRefreshTokenID(ctx context.Context, tokenID string) (models.User, error) {
	const op = "storage.postgres.FindUserByRefreshTokenID"

	// Получаем id пользователя
	query := `
		SELECT user_id
		FROM refresh_tokens
		WHERE token = $1
	`

	var userID string
	err := s.db.QueryRowContext(ctx, query, tokenID).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	query = `
		SELECT id, name, email, goal, pass_hash
		FROM users
		WHERE id = $1
	`

	// Выполняем запрос и сканируем результат в структуру models.User
	var user models.User
	err = s.db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Goal, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) UpdateToken(ctx context.Context, tokenID, userID string, expiresAt time.Time) error {
	const op = "storage.postgres.UpdateTokens"

	query := `
		UPDATE refresh_tokens 
        SET token = $1, expires_at = $2 
        WHERE user_id = $3
	`
	_, err := s.db.ExecContext(ctx, query, tokenID, expiresAt, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
