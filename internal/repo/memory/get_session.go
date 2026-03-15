package memory

import (
	"context"
	"fmt"
	"github.com/aniats/messenger/internal/model"
	"github.com/google/uuid"
)

func (s *Storage) GetSession(_ context.Context, id uuid.UUID) (model.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.sessions[id]
	if !ok {
		return model.Session{}, fmt.Errorf("session not found: %s", id)
	}

	return session, nil
}
