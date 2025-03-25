package main

import (
	"net/http"
	"time"

	"github.com/Abhishek2010dev/Caching-Proxy-With-Redis/handler"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	proxy := handler.NewProxy("http://dummyjson.com", 5*time.Minute, redisClient)
	http.Handle("/", proxy)
	http.ListenAndServe(":3000", nil)
}
