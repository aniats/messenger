package memory

import (
	"context"
	"github.com/aniats/messenger/internal/model"
)

func (s *Storage) ListAllChats(_ context.Context) ([]model.Chat, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	chats := make([]model.Chat, 0, len(s.chats))
	for _, chat := range s.chats {
		chats = append(chats, chat)
	}

	return chats, nil
}
