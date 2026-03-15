// Package app provides the main application container.
package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aniats/messenger/internal/app/grpc"
)

// App is the main application container.
type App struct {
	grpcServer *grpc.Server
}

// New creates a new application with the given logger and gRPC configuration.
func New(log *slog.Logger, grpcCfg grpc.Config) *App {
	grpcApp := grpc.New(log, grpcCfg)

	return &App{
		grpcServer: grpcApp,
	}
}

// Start starts the application.
func (a *App) Start(_ context.Context) error {
	if err := a.grpcServer.Run(); err != nil {
		return fmt.Errorf("failed to start grpc server: %w", err)
	}

	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() {
	a.grpcServer.Stop()
}
