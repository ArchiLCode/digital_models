package app

import (
	"log/slog"
	grpcapp "profile/internal/app/grpc"
	"profile/internal/services/profile"
	"profile/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string) *App {
	storage, err := postgres.New(storagePath)
	if err != nil {
		panic(err)
	}
	profileService := profile.New(log, storage, storage, storage, storage, storage)
	grpcApp := grpcapp.New(log, profileService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
