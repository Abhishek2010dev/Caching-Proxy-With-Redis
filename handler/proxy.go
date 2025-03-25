package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Abhishek2010dev/Caching-Proxy-With-Redis/repository"
	"github.com/redis/go-redis/v9"
)

type Proxy struct {
	origin     string
	cache      *repository.CacheRepository
	expiration time.Duration
}

func NewProxy(origin string, expiration time.Duration, redisClient *redis.Client) *Proxy {
	return &Proxy{
		origin:     origin,
		expiration: expiration,
		cache:      repository.NewCacheRepository(redisClient),
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	CACHE_KEY := fmt.Sprintf("%s:%s", r.Method, r.URL)
	val, err := p.cache.GetCachedEntry(r.Context(), CACHE_KEY)

	if err != nil {
		if errors.Is(err, redis.Nil) {

		}
	}
}
