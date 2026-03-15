package model

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	Id           uuid.UUID
	Name         string
	CreatorId    uuid.UUID // Session.id
	IsReadOnly   bool
	ChatTtlMs    int64 // Zero Value - no ttl
	MessageTtlMs int64 // Zero Value - no ttl
	MaxMessages  int64 // Zero Value - no ttl
	CreatedAt    time.Time
	ExpiresAt    *time.Time
}
