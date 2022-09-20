package clients

import (
	"context"
	"errors"
	"os"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
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

	return client.New(context.Background(), &client.Config{
		ServiceAccountID: aid,
		Token:            ato,
		OrganizationID:   consts.SCHWARZ_ORGANIZATION_ID,
	})
}

// LocalAuthClient is an implementation of the auth client
// relying on env variables for initialization
// the env vars are: STACKIT_AUTH_CLIENT_ID, STACKIT_AUTH_CLIENT_SECRET
func LocalAuthClient() (*client.AuthClient, error) {
	aci := os.Getenv("STACKIT_AUTH_CLIENT_ID")
	if aci == "" {
		return nil, errors.New("STACKIT_AUTH_CLIENT_ID is missing from env variables")
	}

	acs := os.Getenv("STACKIT_AUTH_CLIENT_SECRET")
	if acs == "" {
		return nil, errors.New("STACKIT_AUTH_CLIENT_SECRET is missing from env variables")
	}

	return client.NewAuth(context.Background(), &client.AuthConfig{
		ClientID:     aci,
		ClientSecret: acs,
	})
}
