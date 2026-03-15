package model

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id        uuid.UUID
	CreatedAt time.Time
	ExpiresAt time.Time
}
