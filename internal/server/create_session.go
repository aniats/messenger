package server

import (
	"context"
	pb "github.com/aniats/messenger/pkg/messenger/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSession(ctx context.Context, _ *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	sessionID, err := s.svc.CreateSession(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session: %v", err)
	}

	return &pb.CreateSessionResponse{
		SessionId: sessionID.String(),
	}, nil
}
