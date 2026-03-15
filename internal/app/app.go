package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/aniats/messenger/internal/app/grpc"
	"github.com/aniats/messenger/internal/repo/memory"
	"github.com/aniats/messenger/internal/server"
	"github.com/aniats/messenger/internal/service"
)

type App struct {
	grpcServer *grpc.Server
}

func New(log *slog.Logger, grpcCfg grpc.Config, sessionTTL time.Duration) *App {
	storage := memory.New()

	svc := service.New(storage, sessionTTL)

	handler := server.New(svc)

	grpcApp := grpc.New(log, grpcCfg, handler)

	return &App{
		grpcServer: grpcApp,
	}
}

func (a *App) Start(_ context.Context) error {
	if err := a.grpcServer.Run(); err != nil {
		return fmt.Errorf("failed to start grpc server: %w", err)
	}

	return nil
}

func (a *App) Stop() {
	a.grpcServer.Stop()
}
