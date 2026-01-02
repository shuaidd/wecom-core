package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/shuaidd/wecom-core/pkg/cache"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

const (
	// TokenExpireOffset token 提前刷新时间（秒）
	// 提前 5 分钟刷新 token，避免在使用时过期
	TokenExpireOffset = 300
)

// TokenResponse token响应
type TokenResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// TokenManager Token管理器
type TokenManager struct {
	// corpID 企业ID
	corpID string
	// corpSecret 应用凭证密钥
	corpSecret string
	// baseURL API基础URL
	baseURL string
	// cache 缓存
	cache cache.Cache
	// logger 日志记录器
	logger logger.Logger
	// httpClient HTTP客户端
	httpClient *http.Client
	// refreshLock 刷新锁
	refreshLock sync.Mutex
}

// NewTokenManager 创建Token管理器
func NewTokenManager(corpID, corpSecret, baseURL string, c cache.Cache, log logger.Logger) *TokenManager {
	if c == nil {
		c = NewMemoryCache()
	}

	return &TokenManager{
		corpID:     corpID,
		corpSecret: corpSecret,
		baseURL:    baseURL,
		cache:      c,
		logger:     log,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetToken 获取token
func (tm *TokenManager) GetToken(ctx context.Context) (string, error) {
	// 1. 从缓存获取
	token, expireAt, err := tm.cache.Get(ctx, tm.cacheKey())
	if err == nil && time.Now().Before(expireAt) {
		tm.logger.Debug("Token retrieved from cache",
			logger.F("expire_at", expireAt))
		return token, nil
	}

	// 2. 加锁刷新（防止并发重复获取）
	tm.refreshLock.Lock()
	defer tm.refreshLock.Unlock()

	// 3. Double-check（可能已被其他协程刷新）
	token, expireAt, err = tm.cache.Get(ctx, tm.cacheKey())
	if err == nil && time.Now().Before(expireAt) {
		tm.logger.Debug("Token retrieved from cache after lock",
			logger.F("expire_at", expireAt))
		return token, nil
	}

	// 4. 调用 API 获取 token
	tm.logger.Info("Fetching new token from API")
	token, expiresIn, err := tm.fetchTokenFromAPI(ctx)
	if err != nil {
		tm.logger.Error("Failed to fetch token",
			logger.F("error", err))
		return "", err
	}

	// 5. 缓存 token（提前 5 分钟过期）
	expireAt = time.Now().Add(time.Duration(expiresIn-TokenExpireOffset) * time.Second)
	if err := tm.cache.Set(ctx, tm.cacheKey(), token, expireAt); err != nil {
		tm.logger.Warn("Failed to cache token",
			logger.F("error", err))
	}

	tm.logger.Info("Token refreshed successfully",
		logger.F("expires_in", expiresIn),
		logger.F("expire_at", expireAt))

	return token, nil
}

// RefreshToken 强制刷新token（用于 token 失效重试）
func (tm *TokenManager) RefreshToken(ctx context.Context) error {
	tm.refreshLock.Lock()
	defer tm.refreshLock.Unlock()

	tm.logger.Info("Force refreshing token")

	token, expiresIn, err := tm.fetchTokenFromAPI(ctx)
	if err != nil {
		tm.logger.Error("Failed to refresh token",
			logger.F("error", err))
		return err
	}

	expireAt := time.Now().Add(time.Duration(expiresIn-TokenExpireOffset) * time.Second)
	if err := tm.cache.Set(ctx, tm.cacheKey(), token, expireAt); err != nil {
		tm.logger.Warn("Failed to cache refreshed token",
			logger.F("error", err))
	}

	tm.logger.Info("Token force refreshed successfully",
		logger.F("expires_in", expiresIn),
		logger.F("expire_at", expireAt))

	return nil
}

// fetchTokenFromAPI 从API获取token
func (tm *TokenManager) fetchTokenFromAPI(ctx context.Context) (token string, expiresIn int, err error) {
	url := fmt.Sprintf("%s/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		tm.baseURL, tm.corpID, tm.corpSecret)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := tm.httpClient.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("failed to fetch token: %w", err)
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", 0, fmt.Errorf("failed to decode response: %w", err)
	}

	if tokenResp.ErrCode != 0 {
		return "", 0, fmt.Errorf("failed to get token: errcode=%d, errmsg=%s",
			tokenResp.ErrCode, tokenResp.ErrMsg)
	}

	if tokenResp.AccessToken == "" {
		return "", 0, errors.New("empty access_token in response")
	}

	return tokenResp.AccessToken, tokenResp.ExpiresIn, nil
}

// cacheKey 获取缓存key
func (tm *TokenManager) cacheKey() string {
	return fmt.Sprintf("wecom:token:%s", tm.corpID)
}
