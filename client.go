package stackit

import (
	"context"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/clients"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services"
)

// NewClientWithKeyAuth creates a new client that authenticates itself with STACKIT APIs using a
// service account key and a private RSA key
// this is the recommended way of authenticating to STACKIT API
func NewClientWithKeyAuth(ctx context.Context, cfg ...clients.KeyFlowConfig) (*services.Services, error) {
	client := &clients.KeyFlow{}
	if err := client.Init(ctx, cfg...); err != nil {
		return nil, err
	}
	return services.Init[clients.KeyFlowConfig](client), nil
}

// MustNewClientWithKeyAuth panics if client initialization failed
func MustNewClientWithKeyAuth(ctx context.Context, cfg ...clients.KeyFlowConfig) *services.Services {
	c, err := NewClientWithKeyAuth(ctx, cfg...)
	if err != nil {
		panic(err)
	}
	return c
}

// NewClientWithTokenAuth creates a new client that authenticates itself with STACKIT APIs using a service account token
// important: this approach is less secure, as the token has a long lifespan
func NewClientWithTokenAuth(ctx context.Context, cfg ...clients.TokenFlowConfig) (*services.Services, error) {
	client := &clients.TokenFlow{}
	if err := client.Init(ctx, cfg...); err != nil {
		return nil, err
	}
	return services.Init[clients.TokenFlowConfig](client), nil
}

// MustNewClientWithTokenAuth panics if client initialization failed
func MustNewClientWithTokenAuth(ctx context.Context, cfg ...clients.TokenFlowConfig) *services.Services {
	c, err := NewClientWithTokenAuth(ctx, cfg...)
	if err != nil {
		panic(err)
	}
	return c
}
