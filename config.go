// this file holds configuration for both STACKIT and Auth clients
// once the Auth client is retired, related code will be removed

package client

import (
	"errors"
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

// Config is the STACKIT client configuration
type Config struct {
	BaseUrl             *url.URL
	ServiceAccountToken string
	ServiceAccountEmail string
	Environment         string
}

// Validate verifies that the given config is valid
func (c *Config) Validate() error {
	if c.ServiceAccountToken == "" {
		return errors.New("Service Account Access Token cannot be empty")
	}

	if c.ServiceAccountEmail == "" {
		return errors.New("Service Account Email cannot be empty")
	}

	if c.Environment == "" {
		c.Environment = string(common.ENV_PROD)
	}
	return nil
}
