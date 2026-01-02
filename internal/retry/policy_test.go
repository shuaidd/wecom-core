package retry

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPolicy(t *testing.T) {
	policy := NewPolicy(3, 1*time.Second, 30*time.Second)

	assert.Equal(t, 3, policy.MaxRetries)
	assert.Equal(t, 1*time.Second, policy.InitialBackoff)
	assert.Equal(t, 30*time.Second, policy.MaxBackoff)
}

func TestPolicy_Backoff(t *testing.T) {
	policy := NewPolicy(5, 1*time.Second, 30*time.Second)

	tests := []struct {
		name     string
		attempt  int
		expected time.Duration
	}{
		{
			name:     "attempt 0",
			attempt:  0,
			expected: 1 * time.Second, // 1 * 2^0 = 1
		},
		{
			name:     "attempt 1",
			attempt:  1,
			expected: 2 * time.Second, // 1 * 2^1 = 2
		},
		{
			name:     "attempt 2",
			attempt:  2,
			expected: 4 * time.Second, // 1 * 2^2 = 4
		},
		{
			name:     "attempt 3",
			attempt:  3,
			expected: 8 * time.Second, // 1 * 2^3 = 8
		},
		{
			name:     "attempt 4",
			attempt:  4,
			expected: 16 * time.Second, // 1 * 2^4 = 16
		},
		{
			name:     "attempt 5 - capped at max",
			attempt:  5,
			expected: 30 * time.Second, // 1 * 2^5 = 32, but capped at 30
		},
		{
			name:     "attempt 10 - capped at max",
			attempt:  10,
			expected: 30 * time.Second, // would be 1024, but capped at 30
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backoff := policy.Backoff(tt.attempt)
			assert.Equal(t, tt.expected, backoff)
		})
	}
}
