// Package grpc provides the gRPC server implementation.
package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Config holds the gRPC server configuration.
type Config struct {
	Port int
	Host string
}

// Server wraps the gRPC server with logging and configuration.
type Server struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	cfg        Config
}

// New creates a new gRPC server with the given logger and configuration.
func New(log *slog.Logger, cfg Config) *Server {
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	return &Server{
		log:        log,
		GRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

// Run starts listening and serving gRPC requests.
func (s *Server) Run() error {
	log := s.log.With(slog.String("op", "messenger.Run"))

	listener, err := net.Listen("tcp", net.JoinHostPort(s.cfg.Host, strconv.Itoa(s.cfg.Port)))
	if err != nil {
		return fmt.Errorf("failed to net listen: %w", err)
	}

	log.Info("grpc server started", slog.String("addr", listener.Addr().String()))

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(s.GRPCServer, healthServer)

	reflection.Register(s.GRPCServer)

	if err := s.GRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to grpc server serve: %w", err)
	}

	return nil
}

// Stop gracefully stops the gRPC server.
func (s *Server) Stop() {
	log := s.log.With(slog.String("op", "messenger.Stop"))

	log.Info("stopping gRPC server", slog.Int("port", s.cfg.Port))

	s.GRPCServer.GracefulStop()
}
