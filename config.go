// this file holds configuration for both STACKIT and Auth clients
// once the Auth client is retired, related code will be removed

package client

import (
	"errors"
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// Config is the STACKIT client configuration
type Config struct {
	BaseUrl             *url.URL
	ServiceAccountToken string
	ServiceAccountEmail string
}

// Validate verifies that the given config is valid
func (c *Config) Validate() error {
	if c.BaseUrl == nil {
		c.SetURL("") // set default
	}

	if c.ServiceAccountToken == "" {
		return errors.New("Service Account Access Token cannot be empty")
	}

	if c.ServiceAccountEmail == "" {
		return errors.New("Service Account Email cannot be empty")
	}
	return nil
}

// SetURL sets a given url string as the base url in the config
// if the given value is empty, the default base URL will be used
func (c *Config) SetURL(value string) error {
	if value == "" {
		value = consts.DEFAULT_BASE_URL
	}
	u, err := url.Parse(value)
	if err != nil {
		return err
	}
	c.BaseUrl = u
	return nil
}
