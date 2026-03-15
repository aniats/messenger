package model

import (
	"github.com/google/uuid"
	"time"
)

type ChatMember struct {
	ChatId    uuid.UUID // Chat.id
	SessionId uuid.UUID // Session.id
	JoinedAt  time.Time
}
