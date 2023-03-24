package clients

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	"golang.org/x/oauth2"
)

type StaticTokenFlow struct {
	client *http.Client
	config *StaticTokenFlowConfig
}

type StaticTokenFlowConfig struct {
	ServiceAccountEmail string
	ServiceAccountToken string
	Environment         env.Environment
}

// GetEnvironment returns the defined API environment
func (c *StaticTokenFlow) GetEnvironment() env.Environment {
	return c.config.Environment
}

// GetConfig returns the flow configuration
func (c *StaticTokenFlow) GetConfig() StaticTokenFlowConfig {
	if c.config == nil {
		return StaticTokenFlowConfig{}
	}
	return *c.config
}

func (c *StaticTokenFlow) Init(ctx context.Context, cfg ...StaticTokenFlowConfig) error {
	c.processConfig(cfg...)
	c.configureHTTPClient(ctx)
	return c.validate()
}

// processConfig processes the given configuration
func (c *StaticTokenFlow) processConfig(cfg ...StaticTokenFlowConfig) {
	nc := &StaticTokenFlowConfig{
		ServiceAccountEmail: os.Getenv(ServiceAccountEmail),
		ServiceAccountToken: os.Getenv(ServiceAccountToken),
		Environment:         env.Parse(os.Getenv(Environment)),
	}
	for _, c := range cfg {
		if c.ServiceAccountEmail != "" {
			nc.ServiceAccountEmail = c.ServiceAccountEmail
		}
		if c.ServiceAccountToken != "" {
			nc.ServiceAccountToken = c.ServiceAccountToken
		}
		if c.Environment != "" {
			nc.Environment = c.Environment
		}
	}
	c.config = nc
}

// configureHTTPClient configures the HTTP client
func (c *StaticTokenFlow) configureHTTPClient(ctx context.Context) {
	sts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.config.ServiceAccountToken},
	)
	o2nc := oauth2.NewClient(ctx, sts)
	o2nc.Timeout = time.Second * 10
	c.client = o2nc
}

// validate the client is configured well
func (c *StaticTokenFlow) validate() error {
	if c.config.ServiceAccountToken == "" {
		return errors.New("Service Account Access Token cannot be empty")
	}
	if c.config.ServiceAccountEmail == "" {
		return errors.New("Service Account Email cannot be empty")
	}
	return nil
}

func (c *StaticTokenFlow) Do(req *http.Request) (*http.Response, error) {
	return do(c.client, req, 3, time.Second, time.Minute*2)
}
