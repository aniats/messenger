package server

import (
	"github.com/aniats/messenger/internal/service"
	pb "github.com/aniats/messenger/pkg/messenger/v1"
)

type Server struct {
	pb.UnimplementedMessengerServer
	svc *service.Service
}

func New(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}
