package profile

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"profile/internal/domain/models"
	"profile/internal/lib/logger/sl"
	"strings"
)

type Profile struct {
	log             *slog.Logger
	profSaver       ProfileSaver
	profProvider    ProfileProvider
	appProvider     AppProvider
	profManager     ProfileManager
	catalogProvider CatalogProvider
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppId       = errors.New("invalid identifier")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

type ProfileSaver interface {
	CreateProfile(ctx context.Context, name, lastName, userID string) error
}

type ProfileProvider interface {
	GetProfileByID(ctx context.Context, uuid string) (models.GetProfileByIDResponse, error)
}

type CatalogProvider interface {
	GetCatalog(ctx context.Context, req models.GetCatalogRequest) (models.GetCatalogResponse, error)
}

type ProfileManager interface {
	ProfileDeleter
	ProfileUpdater
}

type ProfileDeleter interface {
	DeleteProfile(ctx context.Context, uuid string) error
}

type ProfileUpdater interface {
	UpdateProfile(ctx context.Context, req models.UpdateProfileRequest) error
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New возвращает новый экземпляр службы Auth
func New(
	log *slog.Logger,
	profileSaver ProfileSaver,
	profileProvider ProfileProvider,
	appProvider AppProvider,
	profileManager ProfileManager,
	catalogProvider CatalogProvider,
) *Profile {
	return &Profile{
		profSaver:       profileSaver,
		profProvider:    profileProvider,
		log:             log,
		appProvider:     appProvider,
		profManager:     profileManager,
		catalogProvider: catalogProvider,
	}
}

// CreateProfile создает профиль нового пользователя в системе и возвращает идентификатор пользователя.
// Если пользователь с данным именем пользователя уже существует, возвращается ошибка.
func (p *Profile) CreateProfile(ctx context.Context, userID, name, lastName string) error {
	const op = "profile.CreateProfile"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("create profile")

	err := p.profSaver.CreateProfile(ctx, name, lastName, userID)
	if err != nil {
		if strings.Contains(err.Error(), "profile already exists") {
			log.Warn("profile already exists", sl.Err(err))

			return fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("failed to save profile", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("profile created")

	return nil
}

func (p *Profile) UpdateProfile(ctx context.Context, req models.UpdateProfileRequest) error {
	const op = "profile.UpdateProfile"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("update profile")

	err := p.profManager.UpdateProfile(ctx, req)
	if err != nil {
		log.Error("failed to update profile", sl.Err(err))
	}

	log.Info("profile updated")

	return nil
}

func (p *Profile) DeleteProfile(ctx context.Context, id string) error {
	const op = "profile.DeleteProfile"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("delete profile")

	err := p.profManager.DeleteProfile(ctx, id)
	if err != nil {
		log.Error("failed to delete profile", sl.Err(err))
	}

	log.Info("profile deleted")

	return nil
}

func (p *Profile) GetProfileByID(ctx context.Context, id string) (models.GetProfileByIDResponse, error) {
	const op = "profile.GetProfileByID"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("get profile")

	resp, err := p.profProvider.GetProfileByID(ctx, id)
	if err != nil {
		log.Error("failed to get profile", sl.Err(err))
		return models.GetProfileByIDResponse{}, err
	}

	log.Info("profile returned")

	return resp, nil
}

func (p *Profile) GetCatalog(ctx context.Context, req models.GetCatalogRequest) (models.GetCatalogResponse, error) {
	const op = "profile.GetCatalog"
	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("get catalog")

	resp, err := p.catalogProvider.GetCatalog(ctx, req)
	if err != nil {
		log.Error("failed to get catalog", sl.Err(err))
		return models.GetCatalogResponse{}, err
	}
	log.Info("catalog returned")
	return resp, nil
}
