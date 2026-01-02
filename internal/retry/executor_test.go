package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalErrors "github.com/shuaidd/wecom-core/internal/errors"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

func TestExecutor_Do_Success(t *testing.T) {
	policy := NewPolicy(3, 10*time.Millisecond, 100*time.Millisecond)
	executor := NewExecutor(policy, logger.NewNoopLogger())

	callCount := 0
	err := executor.Do(context.Background(), func() error {
		callCount++
		return nil
	})

	require.NoError(t, err)
	assert.Equal(t, 1, callCount, "should succeed on first attempt")
}

func TestExecutor_Do_RetryableError(t *testing.T) {
	policy := NewPolicy(3, 10*time.Millisecond, 100*time.Millisecond)
	executor := NewExecutor(policy, logger.NewNoopLogger())

	callCount := 0
	err := executor.Do(context.Background(), func() error {
		callCount++
		if callCount < 3 {
			// Return retriable error (token expired)
			return internalErrors.New(internalErrors.ErrCodeAccessTokenExpired, "token expired")
		}
		return nil
	})

	require.NoError(t, err)
	assert.Equal(t, 3, callCount, "should retry until success")
}

func TestExecutor_Do_NonRetriableError(t *testing.T) {
	policy := NewPolicy(3, 10*time.Millisecond, 100*time.Millisecond)
	executor := NewExecutor(policy, logger.NewNoopLogger())

	callCount := 0
	expectedErr := internalErrors.New(internalErrors.ErrCodeInvalidParameter, "invalid parameter")

	err := executor.Do(context.Background(), func() error {
		callCount++
		return expectedErr
	})

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, 1, callCount, "should not retry non-retriable error")
}

func TestExecutor_Do_MaxRetriesReached(t *testing.T) {
	policy := NewPolicy(3, 10*time.Millisecond, 100*time.Millisecond)
	executor := NewExecutor(policy, logger.NewNoopLogger())

	callCount := 0
	expectedErr := internalErrors.New(internalErrors.ErrCodeAccessTokenExpired, "token expired")

	err := executor.Do(context.Background(), func() error {
		callCount++
		return expectedErr
	})

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	// MaxRetries=3 means: original attempt + 3 retries = 4 total attempts
	assert.Equal(t, 4, callCount, "should attempt 1 + MaxRetries times")
}

func TestExecutor_Do_ContextCanceled(t *testing.T) {
	policy := NewPolicy(10, 100*time.Millisecond, 1*time.Second)
	executor := NewExecutor(policy, logger.NewNoopLogger())

	ctx, cancel := context.WithCancel(context.Background())

	callCount := 0
	err := executor.Do(ctx, func() error {
		callCount++
		if callCount == 2 {
			// Cancel context on second attempt
			cancel()
		}
		return internalErrors.New(internalErrors.ErrCodeAccessTokenExpired, "token expired")
	})

	require.Error(t, err)
	assert.Equal(t, context.Canceled, err)
	// Should stop retrying when context is canceled
	assert.LessOrEqual(t, callCount, 3)
}

func TestExecutor_Do_StandardError(t *testing.T) {
	policy := NewPolicy(3, 10*time.Millisecond, 100*time.Millisecond)
	executor := NewExecutor(policy, logger.NewNoopLogger())

	callCount := 0
	expectedErr := errors.New("standard error")

	err := executor.Do(context.Background(), func() error {
		callCount++
		return expectedErr
	})

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, 1, callCount, "should not retry standard error")
}
