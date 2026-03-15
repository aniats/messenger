package service

import (
	"context"
	"fmt"

	"github.com/aniats/messenger/internal/model"
	"github.com/google/uuid"
)

type ListChatsParams struct {
	SessionID uuid.UUID
	OnlyMy    bool
}

func (s *Service) ListChats(ctx context.Context, params ListChatsParams) ([]model.Chat, error) {
	if _, err := s.ValidateSession(ctx, params.SessionID); err != nil {
		return nil, fmt.Errorf("invalid session: %w", err)
	}

	var chats []model.Chat
	var err error

	if params.OnlyMy {
		chats, err = s.chats.ListChatsBySession(ctx, params.SessionID)
	} else {
		chats, err = s.chats.ListAllChats(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to list chats: %w", err)
	}

	return chats, nil
}
