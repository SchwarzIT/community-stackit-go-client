package main

import (
	"context"
	"encoding/json"
	"os"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/internal/helpers/types"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// Set these values
const (
	userToken           = ""
	serviceAccountEmail = ""
	projectID           = ""
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewStaticTokenClient(ctx)

	// make sure to create an RSA key-pair
	b, err := os.ReadFile("public_key.pem")
	if err != nil {
		panic(err)
	}
	pk := string(b)

	res, err := c.ServiceAccounts.CreateKeys(
		ctx,
		projectID,
		types.Email(serviceAccountEmail),
		serviceaccounts.CreateKeysJSONRequestBody{
			PublicKey: &pk,
		},
	)
	if err = validate.Response(res, err, "JSON201"); err != nil {
		panic(err)
	}

	b, err = json.MarshalIndent(res.JSON201, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("sa_key.json", b, 0644); err != nil {
		panic(err)
	}
}
