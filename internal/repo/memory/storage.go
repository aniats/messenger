package memory

import (
	"github.com/aniats/messenger/internal/model"
	"github.com/google/uuid"
	"sync"
)

type Storage struct {
	mu       sync.RWMutex
	sessions map[uuid.UUID]model.Session
	chats    map[uuid.UUID]model.Chat
	members  map[uuid.UUID][]uuid.UUID // chatID -> []sessionID
	// todo: messages
}

func New() *Storage {
	return &Storage{
		sessions: make(map[uuid.UUID]model.Session),
		chats:    make(map[uuid.UUID]model.Chat),
		members:  make(map[uuid.UUID][]uuid.UUID),
	}
}
