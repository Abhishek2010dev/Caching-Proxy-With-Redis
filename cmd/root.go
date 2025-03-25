package cmd

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

func Init() {

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	roodCmd := &cobra.Command{}
	roodCmd.AddCommand(ServeCommand(redisClient))
	roodCmd.Execute()
}
