package auth

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shuaidd/wecom-core/pkg/cache"
)

func TestMemoryCache_SetAndGet(t *testing.T) {
	c := NewMemoryCache()
	ctx := context.Background()

	key := "test_key"
	token := "test_token"
	expireAt := time.Now().Add(1 * time.Hour)

	// Set
	err := c.Set(ctx, key, token, expireAt)
	require.NoError(t, err)

	// Get
	gotToken, gotExpireAt, err := c.Get(ctx, key)
	require.NoError(t, err)
	assert.Equal(t, token, gotToken)
	assert.Equal(t, expireAt, gotExpireAt)
}

func TestMemoryCache_Get_NotFound(t *testing.T) {
	c := NewMemoryCache()
	ctx := context.Background()

	_, _, err := c.Get(ctx, "non_existent_key")
	assert.Equal(t, cache.ErrCacheNotFound, err)
}

func TestMemoryCache_Get_Expired(t *testing.T) {
	c := NewMemoryCache()
	ctx := context.Background()

	key := "test_key"
	token := "test_token"
	// Set expiration in the past
	expireAt := time.Now().Add(-1 * time.Hour)

	err := c.Set(ctx, key, token, expireAt)
	require.NoError(t, err)

	_, _, err = c.Get(ctx, key)
	assert.Equal(t, cache.ErrCacheExpired, err)
}

func TestMemoryCache_Delete(t *testing.T) {
	c := NewMemoryCache()
	ctx := context.Background()

	key := "test_key"
	token := "test_token"
	expireAt := time.Now().Add(1 * time.Hour)

	// Set
	err := c.Set(ctx, key, token, expireAt)
	require.NoError(t, err)

	// Delete
	err = c.Delete(ctx, key)
	require.NoError(t, err)

	// Get should return not found
	_, _, err = c.Get(ctx, key)
	assert.Equal(t, cache.ErrCacheNotFound, err)
}

func TestMemoryCache_Concurrent(t *testing.T) {
	c := NewMemoryCache()
	ctx := context.Background()

	// Test concurrent writes and reads
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			key := "key"
			token := "token"
			expireAt := time.Now().Add(1 * time.Hour)

			// Set
			err := c.Set(ctx, key, token, expireAt)
			assert.NoError(t, err)

			// Get
			_, _, err = c.Get(ctx, key)
			assert.NoError(t, err)

			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}
