package app

import (
	"log/slog"
	"messenger/internal/app/grpc"
)

type App struct {
	GRPCServer *grpc.Server
}

func New(
	log *slog.Logger,
	grpcPort int,
) *App {

	// todo: storage would be here later

	grpcApp := grpc.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
