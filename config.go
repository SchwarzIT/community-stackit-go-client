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
	e := os.Getenv(email)
	t := os.Getenv(token)
	env := os.Getenv(apienv)

	if c.ServiceAccountToken == "" && t == "" {
		return errors.New("Service Account Access Token cannot be empty")
	}
	if c.ServiceAccountToken == "" {
		c.ServiceAccountToken = t
	}

	if c.ServiceAccountEmail == "" && e == "" {
		return errors.New("Service Account Email cannot be empty")
	}
	if c.ServiceAccountEmail == "" {
		c.ServiceAccountEmail = e
	}

	if c.Environment == "" {
		c.Environment = string(common.ENV_PROD)
	}
	if env != "" {
		c.Environment = env
	}
	return nil
}
