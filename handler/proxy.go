package handler

import (
	"errors"
	"fmt"
	"log"
	"maps"
	"net/http"
	"time"

	"github.com/Abhishek2010dev/Caching-Proxy-With-Redis/models"
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
		// IF cache exits
		if errors.Is(err, redis.Nil) {
			ResponsedWithHeader(w, val, "HIT", CACHE_KEY)
			return
		}
		log.Printf("Failed to get cache entry: %s", err.Error())
		http.Error(w, "Samething went wrong", http.StatusInternalServerError)
		return
	}
}

func ResponsedWithHeader(w http.ResponseWriter, cacheEntry *models.CachedEntry, cacheHeader, KEY string) {
	log.Printf("Cache : %s %s \n", cacheHeader, KEY)
	w.Header().Set("X-Cache", cacheHeader)
	w.WriteHeader(cacheEntry.StatusCode)
	maps.Copy(w.Header(), cacheEntry.Header)
	w.Write(cacheEntry.ResponseBody)
}
