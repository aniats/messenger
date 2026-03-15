package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aniats/messenger/internal/model"
	"github.com/google/uuid"
)

type CreateChatParams struct {
	SessionID    uuid.UUID
	Name         string
	IsReadOnly   bool
	ChatTtlMs    int64
	MessageTtlMs int64
	MaxMessages  int64
}

func (s *Service) CreateChat(ctx context.Context, params CreateChatParams) (uuid.UUID, error) {
	if _, err := s.ValidateSession(ctx, params.SessionID); err != nil {
		return uuid.Nil, fmt.Errorf("invalid session: %w", err)
	}

	now := time.Now()

	chat := model.Chat{
		Id:           uuid.New(),
		Name:         params.Name,
		CreatorId:    params.SessionID,
		IsReadOnly:   params.IsReadOnly,
		ChatTtlMs:    params.ChatTtlMs,
		MessageTtlMs: params.MessageTtlMs,
		MaxMessages:  params.MaxMessages,
		CreatedAt:    now,
	}

	// When chat will expire, the whole chat
	if params.ChatTtlMs > 0 {
		expiresAt := now.Add(time.Duration(params.ChatTtlMs) * time.Millisecond)
		chat.ExpiresAt = &expiresAt
	}

	if err := s.chats.CreateChat(ctx, chat); err != nil {
		return uuid.Nil, fmt.Errorf("failed to create chat: %w", err)
	}

	member := model.ChatMember{
		ChatId:    chat.Id,
		SessionId: params.SessionID,
		JoinedAt:  now,
	}

	if err := s.members.AddMember(ctx, member); err != nil {
		return uuid.Nil, fmt.Errorf("failed to add creator as member: %w", err)
	}

	return chat.Id, nil
}
