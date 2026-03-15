package memory

import (
	"context"
	"github.com/aniats/messenger/internal/model"
)

func (s *Storage) AddMember(_ context.Context, member model.ChatMember) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.members[member.ChatId] = append(s.members[member.ChatId], member.SessionId)
	return nil
}
