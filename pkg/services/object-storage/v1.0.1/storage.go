// Package objectstorage provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.17 DO NOT EDIT.
package objectstorage

import (
	"net/url"
	"strings"

	contracts "github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/project"
	accesskey "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/access-key"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/bucket"
	credentialsgroup "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/credentials-group"
)

// Client which conforms to the OpenAPI3 specification for this service.
type Client[K contracts.ClientFlowConfig] struct {
	// list of connected client services
	AccessKey        *accesskey.Client[K]
	Bucket           *bucket.Client[K]
	CredentialsGroup *credentialsgroup.Client[K]
	Project          *project.Client[K]

	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client contracts.ClientInterface[K]
}

// ClientOption allows setting custom parameters during construction
type ClientOption[K contracts.ClientFlowConfig] func(*Client[K]) error

func NewRawClient[K contracts.ClientFlowConfig](server string, opts ...ClientOption[K]) (*Client[K], error) {
	// create a factory client
	client := Client[K]{
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

	client.AccessKey = accesskey.NewRawClient(server, client.Client)
	client.Bucket = bucket.NewRawClient(server, client.Client)
	client.CredentialsGroup = credentialsgroup.NewRawClient(server, client.Client)
	client.Project = project.NewRawClient(server, client.Client)

	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient[K contracts.ClientFlowConfig](doer contracts.ClientInterface[K]) ClientOption[K] {
	return func(c *Client[K]) error {
		c.Client = doer
		return nil
	}
}

// WithBaseURL overrides the baseURL.
func WithBaseURL[K contracts.ClientFlowConfig](baseURL string) ClientOption[K] {
	return func(c *Client[K]) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponses builds on rawClientInterface to offer response payloads
type ClientWithResponses[K contracts.ClientFlowConfig] struct {
	Client *Client[K]

	// list of connected client services
	AccessKey        *accesskey.ClientWithResponses
	Bucket           *bucket.ClientWithResponses
	CredentialsGroup *credentialsgroup.ClientWithResponses
	Project          *project.ClientWithResponses
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient[K contracts.ClientFlowConfig](server string, opts ...ClientOption[K]) (*ClientWithResponses[K], error) {
	client, err := NewRawClient[K](server, opts...)
	if err != nil {
		return nil, err
	}

	cwr := &ClientWithResponses[K]{Client: client}
	cwr.AccessKey = accesskey.NewClient(server, client.Client)
	cwr.Bucket = bucket.NewClient(server, client.Client)
	cwr.CredentialsGroup = credentialsgroup.NewClient(server, client.Client)
	cwr.Project = project.NewClient(server, client.Client)

	return cwr, nil
}
