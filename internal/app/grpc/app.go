package grpc

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type Server struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *Server {
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	return &Server{
		log:        log,
		GRPCServer: gRPCServer,
		port:       port,
	}
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "messenger.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(s.GRPCServer, healthServer)

	reflection.Register(s.GRPCServer)

	if err := s.GRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Stop() {
	const op = "messenger.Stop"

	s.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", s.port))

	s.GRPCServer.GracefulStop()
}
