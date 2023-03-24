package stackit

import (
	"context"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/clients"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services"
)

// NewKeyAccessClient creates a new client that authenticates itself with STACKIT APIs using a
// service account key and a private RSA key
// this is the recommended way of authenticating to STACKIT API
func NewKeyAccessClient(ctx context.Context, cfg ...clients.KeyAccessFlowConfig) (*services.Services[clients.KeyAccessFlowConfig], error) {
	client := &clients.KeyAccessFlow{}
	if err := client.Init(ctx, cfg...); err != nil {
		return nil, err
	}
	return services.Init[clients.KeyAccessFlowConfig](client), nil
}

// MustNewStaticTokenClient panics if client initialization failed
func MustNewKeyAccessClient(ctx context.Context, cfg ...clients.KeyAccessFlowConfig) *services.Services[clients.KeyAccessFlowConfig] {
	c, err := NewKeyAccessClient(ctx, cfg...)
	if err != nil {
		panic(err)
	}
	return c
}

// NewStaticTokenClient creates a new client that authenticates itself with STACKIT APIs using a service account token
// important: this approach is less secure, as the token has a long lifespan
func NewStaticTokenClient(ctx context.Context, cfg ...clients.StaticTokenFlowConfig) (*services.Services[clients.StaticTokenFlowConfig], error) {
	client := &clients.StaticTokenFlow{}
	if err := client.Init(ctx, cfg...); err != nil {
		return nil, err
	}
	return services.Init[clients.StaticTokenFlowConfig](client), nil
}

// MustNewStaticTokenClient panics if client initialization failed
func MustNewStaticTokenClient(ctx context.Context, cfg ...clients.StaticTokenFlowConfig) *services.Services[clients.StaticTokenFlowConfig] {
	c, err := NewStaticTokenClient(ctx, cfg...)
	if err != nil {
		panic(err)
	}
	return c
}
