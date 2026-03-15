package server

import (
	"context"

	"github.com/aniats/messenger/internal/server/repackers"
	pb "github.com/aniats/messenger/pkg/messenger/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListChats(ctx context.Context, req *pb.ListChatsRequest) (*pb.ListChatsResponse, error) {
	sessionID, err := repackers.ParseSessionID(req.GetSessionId())
	if err != nil {
		return nil, err
	}

	params := repackers.ListChatsRequestToParams(sessionID, req)

	chats, err := s.svc.ListChats(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list chats: %v", err)
	}

	return &pb.ListChatsResponse{
		Chats: repackers.ChatsToProto(chats),
	}, nil
}
