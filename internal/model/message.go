package model

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id        uuid.UUID
	ChatId    uuid.UUID // Chat.id
	SenderId  uuid.UUID // Session.id
	Text      string
	SentAt    time.Time
	ExpiresAt time.Time
}
