package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Abhishek2010dev/Caching-Proxy-With-Redis/models"
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{client}
}

func (c *CacheRepository) StoreCachedEntry(ctx context.Context, key string, cachedEntry *models.CachedEntry, expiration time.Duration) error {
	data, err := json.Marshal(cachedEntry)
	if err != nil {
		return err
	}

	_, err = c.client.Set(ctx, key, data, expiration).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheRepository) GetCachedEntry(ctx context.Context, key string) (*models.CachedEntry, error) {
	val, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var cachedEntry models.CachedEntry
	if err := json.Unmarshal(val, &cachedEntry); err != nil {
		return nil, err
	}

	return &cachedEntry, nil
}
