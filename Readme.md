# 🚀 Cache Proxy With Redis

## 📌 Overview
Cache Proxy With Redis is a simple HTTP proxy server built in Go that caches responses using Redis. If a requested resource is not found in the cache, the proxy fetches it from the origin server and stores the response for future requests.

## 🛠 Features
- 🌍 Forwards HTTP requests to the origin server.
- 🏪 Caches responses in Redis for fast retrieval.
- ⚡ Reduces load on the origin server by serving cached responses.
- 📜 Supports custom expiration time for cache entries.
- 📎 Includes cache hit/miss headers for better debugging.
- 🛠 CLI support using Cobra for flexible server management.
- 🧹 Option to clear the cache using a CLI flag.

## 🔧 Installation & Setup
### Prerequisites
- Go installed on your system.
- Redis server running.

### Clone the Repository
```sh
 git clone https://github.com/Abhishek2010dev/Caching-Proxy-With-Redis.git
 cd Caching-Proxy-With-Redis
```

### Install Dependencies
```sh
go mod tidy
```

### Run the Proxy Server
```sh
go run main.go serve --port 8080 --origin "http://example.com" --expiry "10m"
```

### Clear Cache
```sh
go run main.go serve --clear-cache
```

## 🏗 Usage
- Start the proxy server with the required flags.
- Send HTTP requests through the proxy.
- Cached responses are served instantly if available.
- Use the `--clear-cache` flag to flush the Redis cache.

## 📝 License
This project is licensed under the MIT License.

---
Made with ❤️ using Go & Redis!

