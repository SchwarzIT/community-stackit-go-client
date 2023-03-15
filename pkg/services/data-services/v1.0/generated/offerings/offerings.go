// Package offerings provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.1 DO NOT EDIT.
package offerings

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/do87/oapi-codegen/pkg/runtime"
)

// InstanceSchema defines model for InstanceSchema.
type InstanceSchema struct {
	Create Schema `json:"create"`
	Update Schema `json:"update"`
}

// Offering defines model for Offering.
type Offering struct {
	Description      string          `json:"description"`
	DocumentationUrl string          `json:"documentationUrl"`
	ImageUrl         string          `json:"imageUrl"`
	Latest           bool            `json:"latest"`
	Name             string          `json:"name"`
	Plans            []Plan          `json:"plans"`
	QuotaCount       int             `json:"quotaCount"`
	Schema           *InstanceSchema `json:"schema,omitempty"`
	Version          string          `json:"version"`
}

// Offerings defines model for Offerings.
type Offerings struct {
	Offerings []Offering `json:"offerings"`
}

// Plan defines model for Plan.
type Plan struct {
	Description string `json:"description"`
	Free        bool   `json:"free"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

// Schema defines model for Schema.
type Schema struct {
	Parameters map[string]interface{} `json:"parameters"`
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
	// Get request
	Get(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) Get(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetRequest generates requests for Get
func NewGetRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/offerings", pathParam0)
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
	// Get request
	GetWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetResponse, error)
}

type GetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Offerings
	Error     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetWithResponse request returning *GetResponse
func (c *ClientWithResponses) GetWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.Get(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetResponse(rsp)
}

// ParseGetResponse parses an HTTP response from a GetWithResponse call
func (c *ClientWithResponses) ParseGetResponse(rsp *http.Response) (*GetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Offerings
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	}

	return response, nil
}
