package clients

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	// Key Flow optional env variable (1)
	ServiceAccountKey = "STACKIT_SERVICE_ACCOUNT_KEY"
	PrivateKey        = "STACKIT_PRIVATE_KEY"

	// Key Flow optional env variable (2) using file paths
	ServiceAccountKeyPath = "STACKIT_SERVICE_ACCOUNT_KEY_PATH"
	PrivateKeyPath        = "STACKIT_PRIVATE_KEY_PATH"
)

v-ar tokenAPI = baseurl.New(
	"token",
	"https://service-account.api.stackit.cloud/token",
)

const (
	PrivateKeyBlockType = "PRIVATE KEY"

	// tokenExpirationLeeway is subtracted from a token's expiration time when
	// deciding whether it is still usable, to avoid it expiring between the
	// validity check here and the API validating it upstream.
	tokenExpirationLeeway = 5 * time.Second
)

// KeyFlow handles auth with SA key
type KeyFlow struct {
	client        *http.Client
	config        *KeyFlowConfig
	doer          func(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error)
	key           *ServiceAccountKeyPrivateResponse
	privateKey    *rsa.PrivateKey
	privateKeyPEM []byte
	token         *TokenResponseBody
}

// KeyFlowConfig is the flow config
type KeyFlowConfig struct {
	ServiceAccountKeyPath string
	PrivateKeyPath        string
	ServiceAccountKey     []byte
	PrivateKey            []byte
	ClientRetry           *RetryConfig
	EnableTraceparent     bool
}

// TokenResponseBody is the API response
// when requesting a new token
type TokenResponseBody struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// ServiceAccountKeyPrivateResponse is the API response
// when creating a new SA key
type ServiceAccountKeyPrivateResponse struct {
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"createdAt"`
	Credentials struct {
		Aud        string    `json:"aud"`
		Iss        string    `json:"iss"`
		Kid        string    `json:"kid"`
		PrivateKey *string   `json:"privateKey,omitempty"`
		Sub        uuid.UUID `json:"sub"`
	} `json:"credentials"`
	ID           uuid.UUID  `json:"id"`
	KeyAlgorithm string     `json:"keyAlgorithm"`
	KeyOrigin    string     `json:"keyOrigin"`
	KeyType      string     `json:"keyType"`
	PublicKey    string     `json:"publicKey"`
	ValidUntil   *time.Time `json:"validUntil,omitempty"`
}

// GetConfig returns the flow configuration
func (c *KeyFlow) GetConfig() KeyFlowConfig {
	if c.config == nil {
		return KeyFlowConfig{}
	}
	return *c.config
}

// GetServiceAccountEmail returns the service account email
func (c *KeyFlow) GetServiceAccountEmail() string {
	if c.key == nil {
		return ""
	}
	return c.key.Credentials.Iss
}

// Init intializes the flow
func (c *KeyFlow) Init(ctx context.Context, cfg ...KeyFlowConfig) error {
	c.client = &http.Client{}
	c.token = new(TokenResponseBody)
	c.doer = do
	c.processConfig(cfg...)
	c.configureHTTPClient(ctx)
	if err := c.validateConfig(); err != nil {
		return err
	}
	if err := c.loadFiles(); err != nil {
		return err
	}
	return nil
}

// Clone creates a clone of the client
func (c *KeyFlow) Clone() interface{} {
	sc := *c
	nc := &sc
	cl := *nc.client
	cf := *nc.config
	ke := *nc.key
	to := *nc.token
	nc.client = &cl
	nc.config = &cf
	nc.key = &ke
	nc.token = &to
	return c
}

// Do performs the reuqest
func (c *KeyFlow) Do(req *http.Request) (*http.Response, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return c.doer(c.client, req, c.config.ClientRetry)
}

// GetAccessToken returns short-lived access token
func (c *KeyFlow) GetAccessToken() (string, error) {
	accessTokenIsValid, err := c.validateToken(c.token.AccessToken)
	if err != nil {
		return "", errors.Wrap(err, "failed to validate existing keyflow token")
	}
	if accessTokenIsValid {
		return c.token.AccessToken, nil
	}
	if err := c.recreateAccessToken(); err != nil {
		return "", errors.Wrap(err, "failed to recreate keyflow token")
	}
	return c.token.AccessToken, nil
}

// Flow Configuration

// processConfig processes the given configuration
func (c *KeyFlow) processConfig(cfg ...KeyFlowConfig) {
	if c.config == nil {
		c.config = &KeyFlowConfig{}
	}
	c.config = c.mergeConfigs(c.config, c.getConfigFromEnvironment())
	if c.config.ClientRetry == nil {
		c.config.ClientRetry = NewRetryConfig()
	}
	for _, m := range cfg {
		c.config = c.mergeConfigs(&m, c.config)
	}
	c.config.ClientRetry.Traceparent = &c.config.EnableTraceparent
}

// getConfigFromEnvironment returns a KeyFlowConfig populated with environment variables.
func (c *KeyFlow) getConfigFromEnvironment() *KeyFlowConfig {
	return &KeyFlowConfig{
		ServiceAccountKeyPath: os.Getenv(ServiceAccountKeyPath),
		PrivateKeyPath:        os.Getenv(PrivateKeyPath),
		ServiceAccountKey:     []byte(os.Getenv(ServiceAccountKey)),
		PrivateKey:            []byte(os.Getenv(PrivateKey)),
	}
}

// mergeConfigs returns a new KeyFlowConfig that combines the values of cfg and currentCfg.
func (c *KeyFlow) mergeConfigs(cfg, currentCfg *KeyFlowConfig) *KeyFlowConfig {
	merged := *currentCfg
	if cfg.ServiceAccountKeyPath != "" {
		merged.ServiceAccountKeyPath = cfg.ServiceAccountKeyPath
	}
	if cfg.PrivateKeyPath != "" {
		merged.PrivateKeyPath = cfg.PrivateKeyPath
	}
	if len(cfg.ServiceAccountKey) != 0 {
		merged.ServiceAccountKey = cfg.ServiceAccountKey
	}
	if len(cfg.PrivateKey) != 0 {
		merged.PrivateKey = cfg.PrivateKey
	}

	merged.EnableTraceparent = cfg.EnableTraceparent || merged.EnableTraceparent
	return &merged
}

// configureHTTPClient configures the HTTP client
func (c *KeyFlow) configureHTTPClient(ctx context.Context) {
	client := &http.Client{}
	client.Timeout = DefaultClientTimeout
	c.client = client
}

// validate the client is configured well
func (c *KeyFlow) validateConfig() error {
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
func (c *KeyFlow) loadFiles() error {
	if len(c.config.ServiceAccountKey) == 0 {
		b, err := os.ReadFile(c.config.ServiceAccountKeyPath)
		if err != nil {
			return err
		}
		c.config.ServiceAccountKey = b
	}
	if len(c.config.PrivateKey) == 0 {
		b, err := os.ReadFile(c.config.PrivateKeyPath)
		if err != nil {
			return err
		}
		c.config.PrivateKey = b
	}
	c.key = new(ServiceAccountKeyPrivateResponse)
	err := json.Unmarshal(c.config.ServiceAccountKey, c.key)
	if err != nil {
		return err
	}

	c.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(c.config.PrivateKey)
	if err != nil {
		return err
	}

	// Encode the private key in PEM format
	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.privateKey),
	}
	c.privateKeyPEM = pem.EncodeToMemory(privKeyPEM)

	return nil
}

// Flow auth functions

// recreateAccessToken is used to create a new access token
// when the existing one isn't valid anymore
func (c *KeyFlow) recreateAccessToken() error {
	refreshTokenIsValid, err := c.validateToken(c.token.RefreshToken)
	if err != nil {
		return err
	}
	if refreshTokenIsValid {
		if err := c.createAccessTokenWithRefreshToken(); err == nil {
			return nil
		}
	}
	return c.createAccessToken()
}

// createAccessToken creates an access token using self signed JWT
func (c *KeyFlow) createAccessToken() error {
	grant := "urn:ietf:params:oauth:grant-type:jwt-bearer"
	assertion, err := c.generateSelfSignedJWT()
	if err != nil {
		return err
	}
	res, err := c.requestToken(grant, assertion)
	if err != nil {
		return err
	}
	return c.parseTokenResponse(res)
}

// createAccessTokenWithRefreshToken creates an access token using
// an existing pre-validated refresh token
func (c *KeyFlow) createAccessTokenWithRefreshToken() error {
	res, err := c.requestToken("refresh_token", c.token.RefreshToken)
	if err != nil {
		return err
	}
	return c.parseTokenResponse(res)
}

// generateSelfSignedJWT generates JWT token
func (c *KeyFlow) generateSelfSignedJWT() (string, error) {
	claims := jwt.MapClaims{
		"iss": c.key.Credentials.Iss,
		"sub": c.key.Credentials.Sub,
		"jti": uuid.New(),
		"aud": c.key.Credentials.Aud,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	token.Header["kid"] = c.key.Credentials.Kid
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// requestToken makes a request to the SA token API
func (c *KeyFlow) requestToken(grant, assertion string) (*http.Response, error) {
	body := url.Values{}
	body.Set("grant_type", grant)
	body.Set("assertion", assertion)
	payload := strings.NewReader(body.Encode())
	req, err := http.NewRequest(http.MethodPost, tokenAPI.Get(), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c.doer(&http.Client{}, req, c.config.ClientRetry)
}

// parseTokenResponse parses the response from the server
func (c *KeyFlow) parseTokenResponse(res *http.Response) error {
	if res == nil {
		return errors.New("received bad response from API")
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received: %+v", res)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	c.token = new(TokenResponseBody)
	return json.Unmarshal(body, c.token)
}

// validateToken returns true if token is valid
func (c *KeyFlow) validateToken(token string) (bool, error) {
	if token == "" {
		return false, nil
	}
	parsedToken, err := c.parseToken(token)
	if err != nil {
		var validationError *jwt.ValidationError
		if errors.As(err, &validationError) {
			c.token = new(TokenResponseBody)
			return false, nil
		}
		return false, err
	}

	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)
	if !ok || claims.ExpiresAt == nil {
		c.token = new(TokenResponseBody)
		return false, nil
	}

	// Pretend to be tokenExpirationLeeway into the future, so a token that is
	// about to expire is recreated instead of being sent and rejected upstream.
	return time.Now().Add(tokenExpirationLeeway).Before(claims.ExpiresAt.Time), nil
}

// parseToken parses a JWT token without verifying its signature.
//
// The signature is intentionally not verified: the token was just obtained by
// this client from the token endpoint over TLS, so it is not user input being
// authenticated here - it is only inspected for its expiration time. The API
// the token is sent to does verify it. The JWKS endpoint this used to fetch
// (https://service-account.api.stackit.cloud/.well-known/jwks.json) has been
// retired by STACKIT and now returns 404, which made every request fail with
// "failed to validate existing keyflow token". The official SDK
// (github.com/stackitcloud/stackit-sdk-go/core) does the same thing.
func (c *KeyFlow) parseToken(token string) (*jwt.Token, error) {
	parsedToken, _, err := jwt.NewParser().ParseUnverified(token, &jwt.RegisteredClaims{})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}
