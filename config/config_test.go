package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cfg := New()

		assert.Equal(t, "https://qyapi.weixin.qq.com", cfg.BaseURL)
		assert.Equal(t, 30*time.Second, cfg.Timeout)
		assert.Equal(t, 3, cfg.MaxRetries)
		assert.Equal(t, 1*time.Second, cfg.InitialBackoff)
		assert.Equal(t, 30*time.Second, cfg.MaxBackoff)
		assert.NotNil(t, cfg.Logger)
		assert.Nil(t, cfg.Cache)
	})

	t.Run("with options", func(t *testing.T) {
		cfg := New(
			WithCorpID("test_corp_id"),
			WithCorpSecret("test_secret"),
			WithTimeout(60*time.Second),
			WithRetry(5),
			WithBackoff(2*time.Second, 60*time.Second),
		)

		assert.Equal(t, "test_corp_id", cfg.CorpID)
		assert.Equal(t, "test_secret", cfg.CorpSecret)
		assert.Equal(t, 60*time.Second, cfg.Timeout)
		assert.Equal(t, 5, cfg.MaxRetries)
		assert.Equal(t, 2*time.Second, cfg.InitialBackoff)
		assert.Equal(t, 60*time.Second, cfg.MaxBackoff)
	})
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		wantErr error
	}{
		{
			name: "valid config",
			cfg: &Config{
				CorpID:     "test_corp_id",
				CorpSecret: "test_secret",
				Timeout:    30 * time.Second,
				MaxRetries: 3,
			},
			wantErr: nil,
		},
		{
			name: "missing corp id",
			cfg: &Config{
				CorpSecret: "test_secret",
				Timeout:    30 * time.Second,
				MaxRetries: 3,
			},
			wantErr: ErrMissingCorpID,
		},
		{
			name: "missing corp secret",
			cfg: &Config{
				CorpID:     "test_corp_id",
				Timeout:    30 * time.Second,
				MaxRetries: 3,
			},
			wantErr: ErrMissingCorpSecret,
		},
		{
			name: "invalid timeout",
			cfg: &Config{
				CorpID:     "test_corp_id",
				CorpSecret: "test_secret",
				Timeout:    0,
				MaxRetries: 3,
			},
			wantErr: ErrInvalidTimeout,
		},
		{
			name: "invalid max retries",
			cfg: &Config{
				CorpID:     "test_corp_id",
				CorpSecret: "test_secret",
				Timeout:    30 * time.Second,
				MaxRetries: -1,
			},
			wantErr: ErrInvalidMaxRetries,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWithCorpID(t *testing.T) {
	cfg := New(WithCorpID("test_id"))
	assert.Equal(t, "test_id", cfg.CorpID)
}

func TestWithCorpSecret(t *testing.T) {
	cfg := New(WithCorpSecret("test_secret"))
	assert.Equal(t, "test_secret", cfg.CorpSecret)
}

func TestWithBaseURL(t *testing.T) {
	cfg := New(WithBaseURL("https://custom.url"))
	assert.Equal(t, "https://custom.url", cfg.BaseURL)
}

func TestWithTimeout(t *testing.T) {
	cfg := New(WithTimeout(60 * time.Second))
	assert.Equal(t, 60*time.Second, cfg.Timeout)
}

func TestWithRetry(t *testing.T) {
	cfg := New(WithRetry(5))
	assert.Equal(t, 5, cfg.MaxRetries)
}

func TestWithBackoff(t *testing.T) {
	cfg := New(WithBackoff(2*time.Second, 60*time.Second))
	assert.Equal(t, 2*time.Second, cfg.InitialBackoff)
	assert.Equal(t, 60*time.Second, cfg.MaxBackoff)
}
