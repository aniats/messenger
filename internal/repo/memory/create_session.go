package memory

import (
	"context"
	"github.com/aniats/messenger/internal/model"
)

func (s *Storage) CreateSession(_ context.Context, session model.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.Id] = session
	return nil
}
