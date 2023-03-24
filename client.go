package stackit

import (
	"context"

	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/clients"
)

// NewKeyAccessClient creates a new client that authenticates itself with STACKIT APIs using a
// service account key and a private RSA key
// this is the recommended way of authenticating to STACKIT API
func NewKeyAccessClient(ctx context.Context, cfg ...clients.KeyAccessFlowConfig) (contracts.ClientInterface[clients.KeyAccessFlowConfig], error) {
	client := &clients.KeyAccessFlow{}
	err := client.Init(ctx, cfg...)
	return client, err
}

// NewStaticTokenClient creates a new client that authenticates itself with STACKIT APIs using a service account token
// important: this approach is less secure, as the token has a long lifespan
func NewStaticTokenClient(ctx context.Context, cfg ...clients.StaticTokenFlowConfig) (contracts.ClientInterface[clients.StaticTokenFlowConfig], error) {
	client := &clients.StaticTokenFlow{}
	err := client.Init(ctx, cfg...)
	return client, err
}
