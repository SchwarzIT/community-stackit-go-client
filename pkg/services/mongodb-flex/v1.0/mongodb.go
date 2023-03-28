// Package mongodbflex provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.23 DO NOT EDIT.
package mongodbflex

import (
	"net/url"
	"strings"

	contracts "github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/backup"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/flavors"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/instance"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/user"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/versions"
)

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// list of connected client services
	Backup   *backup.Client
	User     *user.Client
	Versions *versions.Client
	Flavors  *flavors.Client
	Instance *instance.Client

	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client contracts.BaseClientInterface
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

func NewRawClient(server string, opts ...ClientOption) (*Client, error) {
	// create a factory client
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}

	client.Backup = backup.NewRawClient(server, client.Client)
	client.User = user.NewRawClient(server, client.Client)
	client.Versions = versions.NewRawClient(server, client.Client)
	client.Flavors = flavors.NewRawClient(server, client.Client)
	client.Instance = instance.NewRawClient(server, client.Client)

	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer contracts.BaseClientInterface) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponses builds on rawClientInterface to offer response payloads
type ClientWithResponses struct {
	Client *Client

	// list of connected client services
	Backup   *backup.ClientWithResponses
	User     *user.ClientWithResponses
	Versions *versions.ClientWithResponses
	Flavors  *flavors.ClientWithResponses
	Instance *instance.ClientWithResponses
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewRawClient(server, opts...)
	if err != nil {
		return nil, err
	}

	cwr := &ClientWithResponses{Client: client}
	cwr.Backup = backup.NewClient(server, client.Client)
	cwr.User = user.NewClient(server, client.Client)
	cwr.Versions = versions.NewClient(server, client.Client)
	cwr.Flavors = flavors.NewClient(server, client.Client)
	cwr.Instance = instance.NewClient(server, client.Client)

	return cwr, nil
}
