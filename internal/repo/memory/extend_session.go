package memory

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *Storage) ExtendSession(_ context.Context, id uuid.UUID, newExpiresAt time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, ok := s.sessions[id]
	if !ok {
		return fmt.Errorf("session not found: %s", id)
	}

	session.ExpiresAt = newExpiresAt
	s.sessions[id] = session
	return nil
}
