package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Abhishek2010dev/Caching-Proxy-With-Redis/handler"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

func ServeCommand(client *redis.Client) *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the caching proxy server",
		Run: func(cmd *cobra.Command, args []string) {
			clearCache, _ := cmd.Flags().GetBool("clear-cache")
			if clearCache {
				if err := client.FlushAll(cmd.Context()).Err(); err != nil {
					log.Fatalf("Failed to clear cache: %v", err)
				}
				log.Println("Cache cleared successfully.")
				os.Exit(0)
			}

			port, err := cmd.Flags().GetInt("port")
			if err != nil || port == 0 {
				log.Fatal("Error: required flag 'port' not set")
			}

			origin, err := cmd.Flags().GetString("origin")
			if err != nil || origin == "" {
				log.Fatal("Error: required flag 'origin' not set")
			}

			expiryStr, _ := cmd.Flags().GetString("expiry")
			expiry, err := parseExpiry(expiryStr)
			if err != nil {
				log.Fatalf("Invalid expiration format: %v", err)
			}

			log.Printf("Starting server on port %d...", port)
			http.Handle("/", handler.NewProxy(origin, expiry, client))
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
		},
	}

	serveCmd.Flags().IntP("port", "p", 0, "Server port (required)")
	serveCmd.Flags().StringP("origin", "o", "", "Origin URL (required)")
	serveCmd.Flags().StringP("expiry", "e", "5m", "Cache expiry (seconds or Go duration format)")
	serveCmd.Flags().BoolP("clear-cache", "c", false, "Clear the cache and exit")

	return serveCmd
}

func parseExpiry(expiry string) (time.Duration, error) {
	if sec, err := strconv.Atoi(expiry); err == nil {
		return time.Duration(sec) * time.Second, nil
	}
	return time.ParseDuration(expiry)
}
