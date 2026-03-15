// Package main is the entry point for the messenger application.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/aniats/messenger/internal/app"
	"github.com/aniats/messenger/internal/app/grpc"
	"github.com/aniats/messenger/internal/config"
	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func setupLogger(env string, logLevel string) *slog.Logger {
	level := parseLogLevel(logLevel)

	var handler slog.Handler

	switch env {
	case envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:       level,
			AddSource:   false,
			ReplaceAttr: nil,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:       level,
			AddSource:   false,
			ReplaceAttr: nil,
		})
	}

	return slog.New(handler)
}

func onEnv(env string, fn func()) {
	if os.Getenv("ENV") == env {
		fn()
	}
}

func main() {
	onEnv(envLocal, func() {
		if err := godotenv.Load("configs/.env"); err != nil {
			slog.Warn("failed to load .env file", slog.String("error", err.Error()))
		}
	})

	cfg := config.Parse()

	log := setupLogger(cfg.Env, cfg.LogLevel)

	application := app.New(log, grpc.Config{
		Port: cfg.Port,
		Host: "",
	})

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	go func() {
		if err := application.Start(ctx); err != nil {
			log.Error("application start failed", slog.String("error", err.Error()))
			cancel()
		}
	}()

	<-ctx.Done()

	application.Stop()
	log.Info("Gracefully stopped")
}
