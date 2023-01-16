// package client groups together all the services that the STACKIT client supports
// Active services are found under `ProductiveServices` whereas new services or services
// that still need to be further developed or tested, can be put under `Incubator`
// All services must be initialized in the `init` method, and the client must be configured
// during initialization

package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
	waitutil "k8s.io/apimachinery/pkg/util/wait"
)

const (
	ClientTimeoutErr         = "Client.Timeout exceeded while awaiting headers"
	ClientContextDeadlineErr = "context deadline exceeded"
	ClientEOFError           = "unexpected EOF"
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

// New returns a new client
func New(ctx context.Context, cfg Config) (*Client, error) {
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

// GetHTTPClient returns the HTTP client
func (c *Client) GetHTTPClient() *http.Client {
	return c.client
}

// GetBaseURL returns the base url string
func (c *Client) GetBaseURL() string {
	return c.config.BaseUrl.String()
}

// SetBaseURL sets the base url
func (c *Client) SetBaseURL(url string) error {
	return c.config.SetURL(url)
}

// GetConfig returns the client config
func (c *Client) GetConfig() Config {
	return c.config
}

func (c *Client) SetToken(token string) {
	c.config.ServiceAccountToken = token
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

// Request creates a new API request
func (c *Client) Request(ctx context.Context, method, path string, body []byte) (*http.Request, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}
	pathSplit := strings.Split(path, "?")
	rel := &url.URL{Path: pathSplit[0]}
	if len(pathSplit) == 2 {
		rel.RawQuery = pathSplit[1]
	}
	u := c.config.BaseUrl.ResolveReference(rel)
	payload := strings.NewReader(string(body))
	req, err := http.NewRequestWithContext(ctx, method, u.String(), payload)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
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
			if resp != nil && resp.StatusCode == http.StatusBadGateway && maxRetries > 0 {
				maxRetries = maxRetries - 1
				return false, nil
			}
			return true, nil
		}),
	); err != nil {
		return resp, err
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

	c, err = New(context.Background(), Config{
		BaseUrl:             u,
		ServiceAccountToken: "token",
		ServiceAccountEmail: "sa-id",
	})

	return
}
