package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"

	"github.com/aniats/messenger/internal/server"
	pb "github.com/aniats/messenger/pkg/messenger/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Port int
	Host string
}

type Server struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	cfg        Config
}

func New(log *slog.Logger, cfg Config, handler *server.Server) *Server {
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	pb.RegisterMessengerServer(gRPCServer, handler)

	return &Server{
		log:        log,
		GRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

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

func (s *Server) Stop() {
	log := s.log.With(slog.String("op", "messenger.Stop"))

	log.Info("stopping gRPC server", slog.Int("port", s.cfg.Port))

	s.GRPCServer.GracefulStop()
}
