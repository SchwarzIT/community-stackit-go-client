// Package flavors provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.0 DO NOT EDIT.
package flavors

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/do87/oapi-codegen/pkg/runtime"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// InfraFlavor defines model for infra.Flavor.
type InfraFlavor struct {
	Categories  *[]string `json:"categories,omitempty"`
	Cpu         *int      `json:"cpu,omitempty"`
	Description *string   `json:"description,omitempty"`
	ID          *string   `json:"id,omitempty"`
	Memory      *int      `json:"memory,omitempty"`
}

// InfraGetFlavorsResponse defines model for infra.GetFlavorsResponse.
type InfraGetFlavorsResponse struct {
	Flavors *[]InfraFlavor `json:"flavors,omitempty"`
}

// InstanceError defines model for instance.Error.
type InstanceError struct {
	Code    *int                 `json:"code,omitempty"`
	Fields  *map[string][]string `json:"fields,omitempty"`
	Message *string              `json:"message,omitempty"`
	Type    *string              `json:"type,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client common.Client
}

// Creates a new Client, with reasonable defaults
func NewClient(server string, httpClient common.Client) *Client {
	// create a client with sane default values
	client := Client{
		Server: server,
		Client: httpClient,
	}
	return &client
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetFlavors request
	GetFlavors(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetFlavors(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetFlavorsRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetFlavorsRequest generates requests for GetFlavors
func NewGetFlavorsRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/projects/%s/flavors", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, httpClient common.Client) *ClientWithResponses {
	return &ClientWithResponses{NewClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetFlavors request
	GetFlavorsWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetFlavorsResponse, error)
}

type GetFlavorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetFlavorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetFlavorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetFlavorsWithResponse request returning *GetFlavorsResponse
func (c *ClientWithResponses) GetFlavorsWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetFlavorsResponse, error) {
	rsp, err := c.GetFlavors(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetFlavorsResponse(rsp)
}

// ParseGetFlavorsResponse parses an HTTP response from a GetFlavorsWithResponse call
func (c *ClientWithResponses) ParseGetFlavorsResponse(rsp *http.Response) (*GetFlavorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetFlavorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
