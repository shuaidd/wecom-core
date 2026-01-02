package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New(40014, "invalid access_token")

	assert.Equal(t, 40014, err.Code)
	assert.Equal(t, "invalid access_token", err.Message)
	assert.Nil(t, err.Cause)
}

func TestWrap(t *testing.T) {
	cause := errors.New("network error")
	err := Wrap(cause, 10001, "system busy")

	assert.Equal(t, 10001, err.Code)
	assert.Equal(t, "system busy", err.Message)
	assert.Equal(t, cause, err.Cause)
}

func TestError_Error(t *testing.T) {
	t.Run("without cause", func(t *testing.T) {
		err := New(40014, "invalid access_token")
		assert.Equal(t, "wecom error [40014]: invalid access_token", err.Error())
	})

	t.Run("with cause", func(t *testing.T) {
		cause := errors.New("network error")
		err := Wrap(cause, 10001, "system busy")
		assert.Equal(t, "wecom error [10001]: system busy (cause: network error)", err.Error())
	})
}

func TestError_Unwrap(t *testing.T) {
	cause := errors.New("network error")
	err := Wrap(cause, 10001, "system busy")

	assert.Equal(t, cause, err.Unwrap())
}

func TestError_Is(t *testing.T) {
	err1 := New(40014, "invalid access_token")
	err2 := New(40014, "different message")
	err3 := New(42001, "access_token expired")

	assert.True(t, err1.Is(err2), "same error code should match")
	assert.False(t, err1.Is(err3), "different error code should not match")
}

func TestIsWecomError(t *testing.T) {
	t.Run("wecom error", func(t *testing.T) {
		err := New(40014, "invalid access_token")
		assert.True(t, IsWecomError(err))
	})

	t.Run("standard error", func(t *testing.T) {
		err := errors.New("standard error")
		assert.False(t, IsWecomError(err))
	})
}

func TestGetErrorCode(t *testing.T) {
	t.Run("wecom error", func(t *testing.T) {
		err := New(40014, "invalid access_token")
		assert.Equal(t, 40014, GetErrorCode(err))
	})

	t.Run("standard error", func(t *testing.T) {
		err := errors.New("standard error")
		assert.Equal(t, 0, GetErrorCode(err))
	})
}

func TestIsTokenExpired(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "access_token expired",
			err:      New(ErrCodeAccessTokenExpired, "access_token expired"),
			expected: true,
		},
		{
			name:     "invalid access_token",
			err:      New(ErrCodeInvalidAccessToken, "invalid access_token"),
			expected: true,
		},
		{
			name:     "other error",
			err:      New(ErrCodeSystemBusy, "system busy"),
			expected: false,
		},
		{
			name:     "standard error",
			err:      errors.New("standard error"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsTokenExpired(tt.err))
		})
	}
}

func TestIsRateLimited(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "rate limited",
			err:      New(ErrCodeAPIFreqLimit, "api freq out of limit"),
			expected: true,
		},
		{
			name:     "other error",
			err:      New(ErrCodeSystemBusy, "system busy"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsRateLimited(tt.err))
		})
	}
}

func TestIsSystemBusy(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "system busy",
			err:      New(ErrCodeSystemBusy, "system busy"),
			expected: true,
		},
		{
			name:     "other error",
			err:      New(ErrCodeAPIFreqLimit, "api freq out of limit"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsSystemBusy(tt.err))
		})
	}
}

func TestIsRetriable(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "token expired - retriable",
			err:      New(ErrCodeAccessTokenExpired, "access_token expired"),
			expected: true,
		},
		{
			name:     "rate limited - retriable",
			err:      New(ErrCodeAPIFreqLimit, "api freq out of limit"),
			expected: true,
		},
		{
			name:     "system busy - retriable",
			err:      New(ErrCodeSystemBusy, "system busy"),
			expected: true,
		},
		{
			name:     "invalid parameter - not retriable",
			err:      New(ErrCodeInvalidParameter, "invalid parameter"),
			expected: false,
		},
		{
			name:     "standard error - not retriable",
			err:      errors.New("standard error"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsRetriable(tt.err))
		})
	}
}
