package memory

import (
	"context"
	"github.com/aniats/messenger/internal/model"
	"github.com/google/uuid"
)

func (s *Storage) ListChatsBySession(_ context.Context, sessionID uuid.UUID) ([]model.Chat, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var chats []model.Chat
	for chatID, memberIDs := range s.members {
		for _, memberID := range memberIDs {
			if memberID == sessionID {
				if chat, ok := s.chats[chatID]; ok {
					chats = append(chats, chat)
				}
				break
			}
		}
	}

	return chats, nil
}
