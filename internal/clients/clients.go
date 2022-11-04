package clients

import (
	"context"
	"errors"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
)

// LocalClient is an implementation of the client
// relying on env variables for initialization
// the env vars are: STACKIT_SERVICE_ACCOUNT_ID, STACKIT_SERVICE_ACCOUNT_TOKEN
func LocalClient() (*client.Client, error) {
	aid := os.Getenv("STACKIT_SERVICE_ACCOUNT_ID")
	if aid == "" {
		return nil, errors.New("STACKIT_SERVICE_ACCOUNT_ID is missing from env variables")
	}

	ato := os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN")
	if ato == "" {
		return nil, errors.New("STACKIT_SERVICE_ACCOUNT_TOKEN is missing from env variables")
	}

	return client.New(context.Background(), client.Config{
		ServiceAccountEmail: aid,
		ServiceAccountToken: ato,
	})
}
