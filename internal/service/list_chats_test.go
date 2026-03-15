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

func TestListChats_All(t *testing.T) {
	svc := newTestService(time.Hour)
	ctx := context.Background()

	sessionID, _ := svc.CreateSession(ctx)
	svc.CreateChat(ctx, service.CreateChatParams{SessionID: sessionID, Name: "chat-1"})
	svc.CreateChat(ctx, service.CreateChatParams{SessionID: sessionID, Name: "chat-2"})

	chats, err := svc.ListChats(ctx, service.ListChatsParams{
		SessionID: sessionID,
		OnlyMy:    false,
	})

	require.NoError(t, err)
	assert.Len(t, chats, 2)
}

func TestListChats_OnlyMy(t *testing.T) {
	svc := newTestService(time.Hour)
	ctx := context.Background()

	session1, _ := svc.CreateSession(ctx)
	session2, _ := svc.CreateSession(ctx)

	svc.CreateChat(ctx, service.CreateChatParams{SessionID: session1, Name: "s1-chat"})
	svc.CreateChat(ctx, service.CreateChatParams{SessionID: session2, Name: "s2-chat"})

	chats, err := svc.ListChats(ctx, service.ListChatsParams{
		SessionID: session1,
		OnlyMy:    true,
	})

	require.NoError(t, err)
	assert.Len(t, chats, 1)
	assert.Equal(t, "s1-chat", chats[0].Name)
}

func TestListChats_InvalidSession(t *testing.T) {
	svc := newTestService(time.Hour)

	_, err := svc.ListChats(context.Background(), service.ListChatsParams{
		SessionID: uuid.New(),
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid session")
}
