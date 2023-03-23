// this file holds configuration for both STACKIT and Auth clients
// once the Auth client is retired, related code will be removed

package stackit

import (
	"errors"
	"net/url"
	"os"

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
	email := os.Getenv("STACKIT_SERVICE_ACCOUNT_EMAIL")
	token := os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN")
	env := os.Getenv("STACKIT_ENV")

	if c.ServiceAccountToken == "" && token == "" {
		return errors.New("Service Account Access Token cannot be empty")
	}
	if c.ServiceAccountToken == "" {
		c.ServiceAccountToken = token
	}

	if c.ServiceAccountEmail == "" && email == "" {
		return errors.New("Service Account Email cannot be empty")
	}
	if c.ServiceAccountEmail == "" {
		c.ServiceAccountEmail = email
	}

	if c.Environment == "" {
		c.Environment = string(common.ENV_PROD)
	}
	if env != "" {
		c.Environment = env
	}
	return nil
}
