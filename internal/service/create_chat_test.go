package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/aniats/messenger/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateChat(t *testing.T) {
	svc := newTestService(time.Hour)
	ctx := context.Background()

	sessionID, _ := svc.CreateSession(ctx)

	chatID, err := svc.CreateChat(ctx, service.CreateChatParams{
		SessionID: sessionID,
		Name:      "test-chat",
	})

	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, chatID)
}

func TestCreateChat_WithTTL(t *testing.T) {
	svc := newTestService(time.Hour)
	ctx := context.Background()

	sessionID, _ := svc.CreateSession(ctx)

	chatID, err := svc.CreateChat(ctx, service.CreateChatParams{
		SessionID: sessionID,
		Name:      "ttl-chat",
		ChatTtlMs: 60000,
	})

	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, chatID)
}

func TestCreateChat_InvalidSession(t *testing.T) {
	svc := newTestService(time.Hour)

	_, err := svc.CreateChat(context.Background(), service.CreateChatParams{
		SessionID: uuid.New(),
		Name:      "test",
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid session")
}
