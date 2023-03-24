package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

type KeyAccessFlow struct {
	client *http.Client
	config *KeyAccessFlowConfig
}

type KeyAccessFlowConfig struct {
	ServiceAccountKeyPath string
	PrivateKeyPath        string
	ServiceAccountKey     []byte
	PrivateKey            []byte
	Environment           env.Environment
	Token                 TokenResponseBody
}

type TokenResponseBody struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// GetEnvironment returns the defined API environment
func (c *KeyAccessFlow) GetEnvironment() env.Environment {
	return c.config.Environment
}

// GetConfig returns the flow configuration
func (c *KeyAccessFlow) GetConfig() KeyAccessFlowConfig {
	if c.config == nil {
		return KeyAccessFlowConfig{}
	}
	return *c.config
}

func (c *KeyAccessFlow) Init(ctx context.Context, cfg ...KeyAccessFlowConfig) error {
	c.processConfig(cfg...)
	c.configureHTTPClient(ctx)
	if err := c.validate(); err != nil {
		return err
	}
	if err := c.loadFiles(); err != nil {
		return err
	}
	return nil
}

// processConfig processes the given configuration
func (c *KeyAccessFlow) processConfig(cfg ...KeyAccessFlowConfig) {
	nc := &KeyAccessFlowConfig{
		ServiceAccountKeyPath: os.Getenv(ServiceAccountKeyPath),
		PrivateKeyPath:        os.Getenv(PrivateKeyPath),
		ServiceAccountKey:     []byte(os.Getenv(ServiceAccountKey)),
		PrivateKey:            []byte(os.Getenv(PrivateKey)),
		Environment:           env.Parse(os.Getenv(Environment)),
	}
	for _, c := range cfg {
		if c.ServiceAccountKeyPath != "" {
			nc.ServiceAccountKeyPath = c.ServiceAccountKeyPath
		}
		if c.PrivateKeyPath != "" {
			nc.PrivateKeyPath = c.PrivateKeyPath
		}
		if len(c.ServiceAccountKey) != 0 {
			nc.ServiceAccountKey = c.ServiceAccountKey
		}
		if len(c.PrivateKey) != 0 {
			nc.PrivateKey = c.PrivateKey
		}
		if c.Environment != "" {
			nc.Environment = c.Environment
		}
	}
	c.config = nc
}

// configureHTTPClient configures the HTTP client
func (c *KeyAccessFlow) configureHTTPClient(ctx context.Context) {

}

// validate the client is configured well
func (c *KeyAccessFlow) validate() error {
	if len(c.config.ServiceAccountKey) == 0 && c.config.ServiceAccountKeyPath == "" {
		return errors.New("Service Account Key or Key path must be specified")
	}
	if len(c.config.PrivateKey) == 0 && c.config.PrivateKeyPath == "" {
		return errors.New("Private Key or Private Key path must be specified")
	}
	return nil
}

// loadFiles checks if files need to be loaded from specified paths
// and sets them accordingly
func (c *KeyAccessFlow) loadFiles() error {
	if len(c.config.ServiceAccountKey) == 0 {
		b, err := os.ReadFile(c.config.ServiceAccountKeyPath)
		if err != nil {
			return err
		}
		c.config.ServiceAccountKey = b
	}
	if len(c.config.ServiceAccountKey) == 0 {
		b, err := os.ReadFile(c.config.PrivateKeyPath)
		if err != nil {
			return err
		}
		c.config.PrivateKey = b
	}
	return nil
}

func (c *KeyAccessFlow) Do(req *http.Request) (*http.Response, error) {
	return nil, nil
}

var tokenAPI = env.URLs(
	"token",
	"https://api.stackit.cloud/service-account/token",
	"https://api-qa.stackit.cloud/service-account/token",
	"https://api-dev.stackit.cloud/service-account/token",
)

func (c *KeyAccessFlow) getTokens() (string, string, error) {
	grant := url.PathEscape("urn:ietf:params:oauth:grant-type:jwt-bearer")
	assertion := ""
	payload := strings.NewReader(fmt.Sprintf("grant_type=%s&assertion=%s", grant, assertion))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, tokenAPI.GetURL(c.GetEnvironment()), payload)
	if err != nil {
		return "", "", err
	}

	res, err := do(client, req, 3, time.Second, time.Second*60)
	if err != nil {
		return "", "", err
	}
	if res == nil || res.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("received HTTP code %d when trying to get tokens", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}
	var v TokenResponseBody
	if err := json.Unmarshal(body, &v); err != nil {
		return "", "", err
	}
	return v.AccessToken, v.RefreshToken, nil
}

func (c *KeyAccessFlow) makeRequest(httpClient *http.Client, accessToken string) (*http.Response, error) {
	// Make request using access token
	// ...

	// Return response or error
	return nil, nil
}

func (c *KeyAccessFlow) refreshAccessToken(httpClient *http.Client, refreshToken string) (string, error) {
	// Make request to refresh access token
	// ...

	// Parse response to get new access token
	newAccessToken := "..."

	return newAccessToken, nil
}
