package main

import (
	"context"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/retry"
)

func main() {
	c, err := client.New(context.Background(), client.Config{
		ServiceAccountEmail: os.Getenv("STACKIT_SERVICE_ACCOUNT_ID"),
		ServiceAccountToken: os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN"),
	})
	if err != nil {
		panic(err)
	}

	c.SetRetry(retry.New())
}
