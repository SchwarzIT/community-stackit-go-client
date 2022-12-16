// Package postgresflex provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.2.1 DO NOT EDIT.
package postgresflex

import (
	"net/url"
	"strings"

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/backups"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/flavors"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/instance"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/users"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/versions"
)

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// list of connected client services
	Backups  *backups.Client
	Users    *users.Client
	Storage  *storage.Client
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
	Client common.Client
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

func NewClient(server string, opts ...ClientOption) (*Client, error) {
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
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}

	client.Backups = backups.NewClient(server, client.Client)
	client.Users = users.NewClient(server, client.Client)
	client.Storage = storage.NewClient(server, client.Client)
	client.Versions = versions.NewClient(server, client.Client)
	client.Flavors = flavors.NewClient(server, client.Client)
	client.Instance = instance.NewClient(server, client.Client)

	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer common.Client) ClientOption {
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