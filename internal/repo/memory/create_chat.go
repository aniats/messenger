package memory

import (
	"context"
	"github.com/aniats/messenger/internal/model"
)

func (s *Storage) CreateChat(_ context.Context, chat model.Chat) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.chats[chat.Id] = chat
	return nil
}
