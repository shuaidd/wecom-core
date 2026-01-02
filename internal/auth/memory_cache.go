package auth

import (
	"context"
	"sync"
	"time"

	"github.com/shuaidd/wecom-core/pkg/cache"
)

// cacheItem 缓存项
type cacheItem struct {
	token    string
	expireAt time.Time
}

// MemoryCache 内存缓存实现
type MemoryCache struct {
	mu    sync.RWMutex
	items map[string]*cacheItem
}

// NewMemoryCache 创建内存缓存
func NewMemoryCache() cache.Cache {
	return &MemoryCache{
		items: make(map[string]*cacheItem),
	}
}

// Get 获取缓存的token
func (c *MemoryCache) Get(ctx context.Context, key string) (token string, expireAt time.Time, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return "", time.Time{}, cache.ErrCacheNotFound
	}

	// 检查是否过期
	if time.Now().After(item.expireAt) {
		return "", time.Time{}, cache.ErrCacheExpired
	}

	return item.token, item.expireAt, nil
}

// Set 设置缓存的token
func (c *MemoryCache) Set(ctx context.Context, key string, token string, expireAt time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &cacheItem{
		token:    token,
		expireAt: expireAt,
	}

	return nil
}

// Delete 删除缓存的token
func (c *MemoryCache) Delete(ctx context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	return nil
}
