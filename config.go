// this file holds configuration for both STACKIT and Auth clients
// once the Auth client is retired, related code will be removed

package client

import (
	"errors"
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// Config is the STACKIT client configuration
type Config struct {
	BaseUrl          *url.URL
	Token            string
	ServiceAccountID string
	OrganizationID   string
}

// Validate verifies that the given config is valid
func (c *Config) Validate() error {
	if c.BaseUrl == nil {
		c.SetURL("") // set default
	}

	if c.Token == "" {
		return errors.New("STACKIT API: access token is empty")
	}

	if c.ServiceAccountID == "" {
		return errors.New("STACKIT API: service account ID cannot be empty")
	}

	if err := validate.OrganizationID(c.OrganizationID); err != nil {
		return err
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

// Auth Config
// Warning: This code is deprecated and will be removed

// AuthConfig holds information for using auth API
type AuthConfig struct {
	BaseUrl      *url.URL
	ClientID     string
	ClientSecret string
}

// Validate verifies that the given config is valid
func (c *AuthConfig) Validate() error {
	if c.BaseUrl == nil {
		c.SetURL("") // set default
	}

	if c.ClientID == "" {
		return errors.New("auth API: client ID is empty")
	}

	if c.ClientSecret == "" {
		return errors.New("auth API: client Secret is empty")
	}

	return nil
}

// SetURL sets a given url string as the base url in the config
// if the given value is empty, the default auth base URL will be used
func (c *AuthConfig) SetURL(value string) error {
	if value == "" {
		value = consts.DEFAULT_AUTH_BASE_URL
	}
	u, err := url.Parse(value)
	if err != nil {
		return err
	}
	c.BaseUrl = u
	return nil
}
