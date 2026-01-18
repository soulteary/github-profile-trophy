package common

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	memoryCache     = make(map[string]*cacheEntry)
	memoryCacheMu   sync.RWMutex
	redisClient     *redis.Client
	redisClientOnce sync.Once
)

type cacheEntry struct {
	Data      []byte
	ExpiresAt time.Time
}

// InitRedis initializes Redis client if enabled
func InitRedis() {
	enableRedis := os.Getenv("ENABLE_REDIS") == "true"
	if !enableRedis {
		return
	}

	redisClientOnce.Do(func() {
		redisHost := os.Getenv("REDIS_HOST")
		if redisHost == "" {
			redisHost = "localhost"
		}
		redisPort := os.Getenv("REDIS_PORT")
		if redisPort == "" {
			redisPort = "6379"
		}
		redisUsername := os.Getenv("REDIS_USERNAME")
		redisPassword := os.Getenv("REDIS_PASSWORD")

		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisHost + ":" + redisPort,
			Username: redisUsername,
			Password: redisPassword,
		})
	})
}

// GetCache retrieves cached data
func GetCache(key string) ([]byte, bool) {
	// Try memory cache first
	memoryCacheMu.RLock()
	if entry, ok := memoryCache[key]; ok {
		if time.Now().Before(entry.ExpiresAt) {
			memoryCacheMu.RUnlock()
			return entry.Data, true
		}
		// Expired, remove it
		delete(memoryCache, key)
	}
	memoryCacheMu.RUnlock()

	// Try Redis cache
	if redisClient != nil {
		ctx := context.Background()
		val, err := redisClient.Get(ctx, key).Result()
		if err == nil {
			return []byte(val), true
		}
	}

	return nil, false
}

// SetCache stores data in cache
func SetCache(key string, data []byte, ttl time.Duration) {
	// Set in memory cache
	memoryCacheMu.Lock()
	memoryCache[key] = &cacheEntry{
		Data:      data,
		ExpiresAt: time.Now().Add(ttl),
	}
	memoryCacheMu.Unlock()

	// Set in Redis cache
	if redisClient != nil {
		ctx := context.Background()
		redisClient.Set(ctx, key, data, ttl)
	}
}

// GetUserInfoCache retrieves cached user info
func GetUserInfoCache(username string) ([]byte, bool) {
	key := "v1-" + username
	return GetCache(key)
}

// SetUserInfoCache stores user info in cache
func SetUserInfoCache(username string, userInfo interface{}) {
	key := "v1-" + username
	data, err := json.Marshal(userInfo)
	if err != nil {
		return
	}

	ttl := time.Duration(RedisTTL) * time.Millisecond
	SetCache(key, data, ttl)
}
