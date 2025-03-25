package handler

import (
	"errors"
	"fmt"
	"io"
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
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Printf("Failed to get cache entry: %s", err.Error())
		http.Error(w, "Error Getting cache entry", http.StatusInternalServerError)
		return
	}

	// If cache exits.
	if val != nil {
		ResponsedWithHeader(w, val, "HIT", CACHE_KEY)
		return
	}

	// If cache does not exits
	fmt.Printf("Cache Not Present for key : %s \n", CACHE_KEY)
	originUrl := p.origin + r.URL.String()
	resp, err := http.Get(originUrl)
	if err != nil {
		log.Printf("Failed to forward request: %s", err.Error())
		http.Error(w, "Error Forwarding Request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to forward request body: %s", err.Error())
		http.Error(w, "Error Forwarding Request Body", http.StatusInternalServerError)
		return
	}

	entry := models.CachedEntry{
		StatusCode:   resp.StatusCode,
		Header:       resp.Header,
		ResponseBody: body,
		Created:      time.Now(),
	}
	err = p.cache.StoreCachedEntry(r.Context(), CACHE_KEY, &entry, p.expiration)
	if err != nil {
		log.Printf("Failed to set cache entry: %s", err.Error())
		http.Error(w, "Error Setting cache entry", http.StatusInternalServerError)
		return

	}

	ResponsedWithHeader(w, &entry, "MISS", CACHE_KEY)
}

func ResponsedWithHeader(w http.ResponseWriter, cacheEntry *models.CachedEntry, cacheHeader, KEY string) {
	log.Printf("Cache : %s %s \n", cacheHeader, KEY)
	w.Header().Set("X-Cache", cacheHeader)
	maps.Copy(w.Header(), cacheEntry.Header)
	w.WriteHeader(cacheEntry.StatusCode)
	w.Write(cacheEntry.ResponseBody)
}
