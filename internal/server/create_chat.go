package server

import (
	"context"
	"github.com/aniats/messenger/internal/server/repackers"

	pb "github.com/aniats/messenger/pkg/messenger/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	sessionID, err := repackers.ParseSessionID(req.GetSessionId())
	if err != nil {
		return nil, err
	}

	params := repackers.CreateChatRequestToParams(sessionID, req)

	chatID, err := s.svc.CreateChat(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create chat: %v", err)
	}

	return &pb.CreateChatResponse{
		ChatId: chatID.String(),
	}, nil
}
