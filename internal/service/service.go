package service

import (
	"context"
	"time"

	"github.com/aniats/messenger/internal/model"
	"github.com/google/uuid"
)

type SessionStorage interface {
	CreateSession(ctx context.Context, session model.Session) error
	GetSession(ctx context.Context, id uuid.UUID) (model.Session, error)
	ExtendSession(ctx context.Context, id uuid.UUID, newExpiresAt time.Time) error
}

type ChatStorage interface {
	CreateChat(ctx context.Context, chat model.Chat) error
	ListAllChats(ctx context.Context) ([]model.Chat, error)
	ListChatsBySession(ctx context.Context, sessionID uuid.UUID) ([]model.Chat, error)
}

type MemberStorage interface {
	AddMember(ctx context.Context, member model.ChatMember) error
}

type Storage interface {
	SessionStorage
	ChatStorage
	MemberStorage
}

type Service struct {
	sessions   SessionStorage
	chats      ChatStorage
	members    MemberStorage
	sessionTTL time.Duration
}

func New(storage Storage, sessionTTL time.Duration) *Service {
	return &Service{
		sessions:   storage,
		chats:      storage,
		members:    storage,
		sessionTTL: sessionTTL,
	}
}
