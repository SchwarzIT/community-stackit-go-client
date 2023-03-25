package clients

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	// Key Access Flow (1)
	ServiceAccountKey = "STACKIT_SERVICE_ACCOUNT_KEY"
	PrivateKey        = "STACKIT_PRIVATE_KEY"

	// Key Access Flow (2) using file paths
	ServiceAccountKeyPath = "STACKIT_SERVICE_ACCOUNT_KEY_PATH"
	PrivateKeyPath        = "STACKIT_PRIVATE_KEY_PATH"
)

var tokenAPI = env.URLs(
	"token",
	"https://api.stackit.cloud/service-account/token",
	"https://api-qa.stackit.cloud/service-account/token",
	"https://api-dev.stackit.cloud/service-account/token",
)

const (
	PrivateKeyBlockType = "PRIVATE KEY"
)

// KeyFlow handles auth with SA key
type KeyFlow struct {
	client     *http.Client
	config     *KeyFlowConfig
	key        *ServiceAccountKeyPrivateResponse
	privateKey *rsa.PrivateKey
	token      *TokenResponseBody
}

// KeyFlowConfig is the flow config
type KeyFlowConfig struct {
	ServiceAccountKeyPath string
	PrivateKeyPath        string
	ServiceAccountKey     []byte
	PrivateKey            []byte
	Environment           env.Environment
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

// GetEnvironment returns the defined API environment
func (c *KeyFlow) GetEnvironment() env.Environment {
	return c.config.Environment
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

// Do performs the reuqest
func (c *KeyFlow) Do(req *http.Request) (*http.Response, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return do(c.client, req, 3, time.Second, time.Minute*2)
}

// GetAccessToken returns short-lived access token
func (c *KeyFlow) GetAccessToken() (string, error) {
	accessTokenIsValid, err := c.validateToken(c.token.AccessToken)
	if err != nil {
		return "", err
	}
	if accessTokenIsValid {
		return c.token.AccessToken, nil
	}

	if err := c.recreateAccessToken(); err != nil {
		return "", err
	}
	return c.token.AccessToken, nil
}

// Flow Configuration

// processConfig processes the given configuration
func (c *KeyFlow) processConfig(cfg ...KeyFlowConfig) {
	defaultCfg := c.getConfigFromEnvironment()
	if len(cfg) > 0 {
		c.config = mergeConfigs(&cfg[0], defaultCfg)
	} else {
		c.config = defaultCfg
	}
}

// getConfigFromEnvironment returns a KeyFlowConfig populated with environment variables.
func (c *KeyFlow) getConfigFromEnvironment() *KeyFlowConfig {
	return &KeyFlowConfig{
		ServiceAccountKeyPath: os.Getenv(ServiceAccountKeyPath),
		PrivateKeyPath:        os.Getenv(PrivateKeyPath),
		ServiceAccountKey:     []byte(os.Getenv(ServiceAccountKey)),
		PrivateKey:            []byte(os.Getenv(PrivateKey)),
		Environment:           env.Parse(os.Getenv(Environment)),
	}
}

// mergeConfigs returns a new KeyFlowConfig that combines the values of cfg and defaultCfg.
func mergeConfigs(cfg, defaultCfg *KeyFlowConfig) *KeyFlowConfig {
	merged := *defaultCfg

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
	if cfg.Environment != "" {
		merged.Environment = cfg.Environment
	}

	return &merged
}

// configureHTTPClient configures the HTTP client
func (c *KeyFlow) configureHTTPClient(ctx context.Context) {
	client := &http.Client{}
	client.Timeout = time.Second * 10
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

	// privateKey, err := c.parsePrivateKeyPEM(c.config.PrivateKey)
	// if err != nil {
	// 	return err
	// }
	// pem, err := c.encodePrivateKeyToPEM(privateKey)
	// if err != nil {
	// 	return err
	// }
	c.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(c.config.PrivateKey)
	if err != nil {
		return err
	}
	return nil
}

// parsePrivateKeyPEM returns a RSA private key parsed from a PEM block in the supplied data.
func (c *KeyFlow) parsePrivateKeyPEM(keyData []byte) (*rsa.PrivateKey, error) {
	var privateKeyPemBlock *pem.Block
	for {
		privateKeyPemBlock, keyData = pem.Decode(keyData)

		// The service account API returns a private key of type "PRIVATE KEY"
		// so here that's the only accepted pem block type
		if privateKeyPemBlock == nil || privateKeyPemBlock.Type != PrivateKeyBlockType {
			break
		}

		// RSA Private Key in unencrypted PKCS#8 format
		key, err := x509.ParsePKCS8PrivateKey(privateKeyPemBlock.Bytes)
		if err == nil {
			return key.(*rsa.PrivateKey), nil
		}
	}

	// we read all the PEM blocks and didn't recognize one
	return nil, fmt.Errorf("data does not contain a valid RSA private key")
}

func (c *KeyFlow) encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) ([]byte, error) {
	// Get ASN.1 DER format
	privDER, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// pem.Block
	privBlock := pem.Block{
		Type:    "PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)
	return privatePEM, nil
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
		return c.createAccessTokenWithRefreshToken()
	}
	return c.createAccessToken()
}

// createAccessToken creates an access token using self signed JWT
func (c *KeyFlow) createAccessToken() error {
	grant := url.PathEscape("urn:ietf:params:oauth:grant-type:jwt-bearer")
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

// validateToken parses and validates a JWT token
func (c *KeyFlow) parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return c.config.PrivateKey, nil
	})
}

// validateToken returns true if tokeb is valid
func (c *KeyFlow) validateToken(token string) (bool, error) {
	if token == "" {
		return false, nil
	}
	if _, err := c.parseToken(token); err != nil {
		return false, err
	}
	return true, nil
}

// requestToken makes a request to the SA token API
func (c *KeyFlow) requestToken(grant, assertion string) (*http.Response, error) {
	payload := strings.NewReader(fmt.Sprintf("grant_type=%s&assertion=%s", grant, assertion))
	req, err := http.NewRequest(http.MethodPost, tokenAPI.GetURL(c.GetEnvironment()), payload)
	if err != nil {
		return nil, err
	}
	return do(&http.Client{}, req, 3, time.Second, time.Minute)
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
