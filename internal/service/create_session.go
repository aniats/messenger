package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aniats/messenger/internal/model"
	"github.com/google/uuid"
)

func (s *Service) CreateSession(ctx context.Context) (uuid.UUID, error) {
	now := time.Now()

	session := model.Session{
		Id:        uuid.New(),
		CreatedAt: now,
		ExpiresAt: now.Add(s.sessionTTL),
	}

	if err := s.sessions.CreateSession(ctx, session); err != nil {
		return uuid.Nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session.Id, nil
}

func (s *Service) ValidateSession(ctx context.Context, sessionID uuid.UUID) (model.Session, error) {
	session, err := s.sessions.GetSession(ctx, sessionID)
	if err != nil {
		return model.Session{}, fmt.Errorf("session not found: %w", err)
	}

	if time.Now().After(session.ExpiresAt) {
		return model.Session{}, fmt.Errorf("session expired")
	}

	newExpiresAt := time.Now().Add(s.sessionTTL)
	if err := s.sessions.ExtendSession(ctx, sessionID, newExpiresAt); err != nil {
		return model.Session{}, fmt.Errorf("failed to extend session: %w", err)
	}

	session.ExpiresAt = newExpiresAt
	return session, nil
}
