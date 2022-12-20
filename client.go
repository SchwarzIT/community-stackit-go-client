// package client groups together all the services that the STACKIT client supports
// Active services are found under `ProductiveServices` whereas new services or services
// that still need to be further developed or tested, can be put under `Incubator`
// All services must be initialized in the `init` method, and the client must be configured
// during initialization

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"golang.org/x/oauth2"
)

// Client service for managing interactions with STACKIT API
type Client struct {
	ctx    context.Context
	client *http.Client
	config Config

	Services services

	// Legacy
	//----------
	// Productive services - services that are ready to be used in production
	ProductiveServices

	// Incubator - services under development or currently being tested
	// not ready for production usage
	Incubator IncubatorServices
}

// New returns a new client
func New(ctx context.Context, cfg Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	c := &Client{
		config: cfg,
		ctx:    ctx,
	}

	c.setHttpClient(c.ctx)
	c.initServices()
	c.initLegacyServices()
	return c, nil
}

// Clone creates a shallow clone of the client
func (c *Client) Clone() common.Client {
	nc := *c
	return &nc
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

// LegacyDo performs the request, including retry if set
// To set retry, use WithRetry() which returns a shalow copy of the client
func (c *Client) LegacyDo(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error) {
	return c.do(req, v, errorHandlers...)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.LegacyDo(req, nil)
}

// Do performs the request and decodes the response if given interface != nil
func (c *Client) do(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// handle errors in the response
	if len(errorHandlers) == 0 {
		errorHandlers = append(errorHandlers, validate.DefaultResponseErrorHandler)
	}
	for _, fn := range errorHandlers {
		if err := fn(resp); err != nil {
			return resp, err
		}
	}

	// parse response JSON
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		defer resp.Body.Close()
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
