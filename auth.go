// the auth client is a separate client consuming STACKIT's auth API
// it was developed for the purpose of migrating Schwarz IT KG users into STACKIT
// but this behavior is deprecated and will soon be removed

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"golang.org/x/oauth2"
)

// AuthClient manages communication with Auth API
type AuthClient struct {
	Client *http.Client
	Config *AuthConfig
}

// NewAuth returns a new Auth API client
func NewAuth(ctx context.Context, cfg *AuthConfig) (*AuthClient, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	c := &AuthClient{Config: cfg}
	return c.init(ctx), nil
}

// init initializes the client
func (c *AuthClient) init(ctx context.Context) *AuthClient {
	c.setHttpClient(ctx)
	return c
}

// setHttpClient creates the client's oauth client
func (c *AuthClient) setHttpClient(ctx context.Context) {
	hcl := oauth2.NewClient(ctx, nil)
	hcl.Timeout = time.Second * 10
	c.Client = hcl
}

// Request creates an API request
func (c *AuthClient) Request(ctx context.Context, method, path string, body string) (*http.Request, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}
	rel := &url.URL{Path: path}
	u := c.Config.BaseUrl.ResolveReference(rel)
	payload := strings.NewReader(body)
	req, err := http.NewRequestWithContext(ctx, method, u.String(), payload)
	if err != nil {
		return nil, err
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Do performs the request and decodes the response if given interface != nil
func (c *AuthClient) Do(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
	}
	return resp, err
}

// AuthAPIClientCredentialFlowRes response structure for client credentials flow
type AuthAPIClientCredentialFlowRes struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	JTI         string `json:"jti"`
}

// GetToken returns a token from Auth API
func (c *AuthClient) GetToken(ctx context.Context) (res AuthAPIClientCredentialFlowRes, err error) {
	params := url.Values{}
	params.Add("client_id", c.Config.ClientID)
	params.Add("client_secret", c.Config.ClientSecret)
	params.Add("grant_type", "client_credentials")
	body := params.Encode()

	req, err := c.Request(ctx, http.MethodPost, "/oauth/token", body)
	if err != nil {
		return
	}

	_, err = c.Do(req, &res)
	return
}

// MockAuthServer mocks an authentication server
// and returns an auth client pointing to it, mux, teardown function and an error indicator
func MockAuthServer() (c *AuthClient, mux *http.ServeMux, teardown func(), err error) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	teardown = server.Close

	u, _ := url.Parse(server.URL)

	c, err = NewAuth(context.Background(), &AuthConfig{
		BaseUrl:      u,
		ClientID:     "id",
		ClientSecret: "secret",
	})

	return
}
