// package client groups together all the services that the STACKIT client supports
// Active services are found under `ProductiveServices` whereas new services or services
// that still need to be further developed or tested, can be put under `Incubator`
// All services must be initialized in the `init` method, and the client must be configured
// during initialization

package stackit

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	waitutil "k8s.io/apimachinery/pkg/util/wait"
)

const (
	ClientTimeoutErr           = "Client.Timeout exceeded while awaiting headers"
	ClientContextDeadlineErr   = "context deadline exceeded"
	ClientConnectionRefusedErr = "connection refused"
	ClientEOFError             = "unexpected EOF"
	ClientGWTimeoutFError      = "Gateway Timeout"
)

// Client service for managing interactions with STACKIT API
type Client struct {
	ctx    context.Context
	client *http.Client
	config Config

	// when an internal server error is encountered
	// the call will be retried
	RetryTimout time.Duration // timeout for retrying a call
	RetryWait   time.Duration // how long to wait before trying again

	services
}

// NewClientWithConfig returns a new client
func NewClientWithConfig(ctx context.Context, cfg Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	c := &Client{
		config:      cfg,
		ctx:         ctx,
		RetryTimout: 2 * time.Minute,
		RetryWait:   30 * time.Second,
	}

	c.setHttpClient(c.ctx)
	c.initServices()
	return c, nil
}

// NewClient returns a new client and
// panics if there's an error
// to avoid panics, use NewClientWithConfig instead
func NewClient(ctx context.Context) *Client {
	cfg := Config{}
	err := cfg.Validate()
	if err != nil {
		panic(err)
	}

	c := &Client{
		config:      cfg,
		ctx:         ctx,
		RetryTimout: 2 * time.Minute,
		RetryWait:   30 * time.Second,
	}

	c.setHttpClient(c.ctx)
	c.initServices()
	return c
}

// GetHTTPClient returns the HTTP client
func (c *Client) GetHTTPClient() *http.Client {
	return c.client
}

// GetConfig returns the client config
func (c *Client) GetConfig() Config {
	return c.config
}

// GetEnvironment returns the client environment
func (c *Client) GetEnvironment() common.Environment {
	switch strings.ToLower(c.config.Environment) {
	case "dev":
		return common.ENV_DEV
	case "qa":
		return common.ENV_QA
	default:
		return common.ENV_PROD
	}
}

// setHttpClient creates the client's oauth client
func (c *Client) setHttpClient(ctx context.Context) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.config.ServiceAccountToken},
	)
	hcl := oauth2.NewClient(ctx, ts)
	hcl.Timeout = time.Second * 10
	c.client = hcl
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.do(req)
}

// Do performs the request and decodes the response if given interface != nil
func (c *Client) do(req *http.Request) (resp *http.Response, err error) {
	maxRetries := 3
	if err := waitutil.PollImmediate(c.RetryWait, c.RetryTimout, waitutil.ConditionFunc(
		func() (bool, error) {
			resp, err = c.client.Do(req)
			if err != nil {
				if maxRetries > 0 {
					if strings.Contains(err.Error(), ClientTimeoutErr) ||
						(strings.Contains(err.Error(), ClientContextDeadlineErr) && req.Context().Err() == nil) ||
						strings.Contains(err.Error(), ClientConnectionRefusedErr) ||
						strings.Contains(err.Error(), ClientGWTimeoutFError) ||
						(req.Method == http.MethodGet && strings.Contains(err.Error(), ClientEOFError)) {

						// reduce retries counter and retry
						maxRetries = maxRetries - 1
						return false, nil
					}
				}
				return false, err
			}
			if resp != nil && resp.StatusCode == http.StatusInternalServerError {
				return false, nil
			}
			if resp != nil && (resp.StatusCode == http.StatusBadGateway || resp.StatusCode == http.StatusGatewayTimeout) && maxRetries > 0 {
				maxRetries = maxRetries - 1
				return false, nil
			}
			return true, nil
		}),
	); err != nil {
		return resp, errors.Wrap(err, fmt.Sprintf("url: %s\nmethod: %s", req.URL.String(), req.Method))
	}

	return resp, err
}

// MockServer mocks STACKIT api server
// and returns a client pointing to it, mux, teardown function and an error indicator
func MockServer() (c *Client, mux *http.ServeMux, teardown func(), err error) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	teardown = server.Close

	u, _ := url.Parse(server.URL)

	c, err = NewClientWithConfig(context.Background(), Config{
		BaseUrl:             u,
		ServiceAccountToken: "token",
		ServiceAccountEmail: "sa-id",
	})

	return
}
