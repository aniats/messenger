package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/aniats/messenger/internal/repo/memory"
	"github.com/aniats/messenger/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestService(ttl time.Duration) *service.Service {
	return service.New(memory.New(), ttl)
}

func TestCreateSession(t *testing.T) {
	svc := newTestService(time.Hour)

	id, err := svc.CreateSession(context.Background())

	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)
}

func TestValidateSession_OK(t *testing.T) {
	svc := newTestService(time.Hour)
	ctx := context.Background()

	id, _ := svc.CreateSession(ctx)

	session, err := svc.ValidateSession(ctx, id)

	require.NoError(t, err)
	assert.Equal(t, id, session.Id)
	assert.True(t, session.ExpiresAt.After(time.Now()))
}

func TestValidateSession_NotFound(t *testing.T) {
	svc := newTestService(time.Hour)

	_, err := svc.ValidateSession(context.Background(), uuid.New())

	assert.Error(t, err)
}

func TestValidateSession_Expired(t *testing.T) {
	svc := newTestService(time.Nanosecond)
	ctx := context.Background()

	id, _ := svc.CreateSession(ctx)
	time.Sleep(time.Millisecond)

	_, err := svc.ValidateSession(ctx, id)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "expired")
}
