package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

func StartCommand(client *redis.Client) *cobra.Command {
	startCmd := &cobra.Command{
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

		},
	}
	startCmd.Flags().Int16P("port", "p", 0, "Server Port (required)")
	startCmd.Flags().StringP("origin", "o", "", "Origin to sent request (required)")
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
