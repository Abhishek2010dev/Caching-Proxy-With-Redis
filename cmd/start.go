package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Abhishek2010dev/Caching-Proxy-With-Redis/handler"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

func StartCommand(client *redis.Client) *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "This command start server",
		Run: func(cmd *cobra.Command, args []string) {
			port, err := cmd.Flags().GetInt16("port")
			if err != nil {
				fmt.Println("Error retrieving port flag:", err)
				os.Exit(1)
			}

			origin, err := cmd.Flags().GetString("origin")
			if err != nil {
				fmt.Println("Error retrieving origin flag:", err)
				os.Exit(1)
			}

			expiryStr, err := cmd.Flags().GetString("expiry")
			if err != nil {
				fmt.Println("Error retrieving expiry flag:", err)
				os.Exit(1)
			}

			expiry, err := parseExpiry(expiryStr)
			if err != nil {
				fmt.Printf("invalid expiration format: %s \n", err)
				os.Exit(1)
			}

			proxy := handler.NewProxy(origin, expiry, client)
			http.Handle("/", proxy)
			http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
		},
	}
	startCmd.Flags().Int16P("port", "p", 0, "Server Port (required)")
	startCmd.Flags().StringP("origin", "o", "", "Origin to sent request (required)")
	startCmd.Flags().StringP("expiry", "e", "300", "Expiration time in seconds or Go duration format (e.g., '1m30s')")
	startCmd.MarkFlagRequired("port")
	startCmd.MarkFlagRequired("origin")
	return startCmd
}

func parseExpiry(expiry string) (time.Duration, error) {
	if sec, err := strconv.Atoi(expiry); err == nil {
		return time.Duration(sec) * time.Second, nil
	}
	return time.ParseDuration(expiry)
}
