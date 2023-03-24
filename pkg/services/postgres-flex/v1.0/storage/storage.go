// Package storage provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.23 DO NOT EDIT.
package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	contracts "github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/internal/helpers/runtime"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// InstanceError defines model for instance.Error.
type InstanceError struct {
	Code    *int                 `json:"code,omitempty"`
	Fields  *map[string][]string `json:"fields,omitempty"`
	Message *string              `json:"message,omitempty"`
	Type    *string              `json:"type,omitempty"`
}

// InstanceGetFlavorStorageResponse defines model for instance.GetFlavorStorageResponse.
type InstanceGetFlavorStorageResponse struct {
	StorageClasses *[]string             `json:"storageClasses,omitempty"`
	StorageRange   *InstanceStorageRange `json:"storageRange,omitempty"`
}

// InstanceStorageRange defines model for instance.StorageRange.
type InstanceStorageRange struct {
	Max *int `json:"max,omitempty"`
	Min *int `json:"min,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client[K contracts.ClientFlowConfig] struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client contracts.ClientInterface[K]
}

// NewRawClient Creates a new Client, with reasonable defaults
func NewRawClient[K contracts.ClientFlowConfig](server string, httpClient contracts.ClientInterface[K]) *Client[K] {
	// create a client with sane default values
	client := Client[K]{
		Server: server,
		Client: httpClient,
	}
	return &client
}

// The interface specification for the client above.
type rawClientInterface interface {
	// GetStorageOptions request
	GetStorageOptionsRaw(ctx context.Context, projectID string, flavorID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client[K]) GetStorageOptionsRaw(ctx context.Context, projectID string, flavorID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetStorageOptionsRequest(ctx, c.Server, projectID, flavorID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetStorageOptionsRequest generates requests for GetStorageOptions
func NewGetStorageOptionsRequest(ctx context.Context, server string, projectID string, flavorID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "flavorID", runtime.ParamLocationPath, flavorID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/storages/%s", pathParam0, pathParam1)
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

func (c *Client[K]) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on rawClientInterface to offer response payloads
type ClientWithResponses[K contracts.ClientFlowConfig] struct {
	rawClientInterface
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient[K contracts.ClientFlowConfig](server string, httpClient contracts.ClientInterface[K]) *ClientWithResponses[K] {
	return &ClientWithResponses[K]{NewRawClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface[K contracts.ClientFlowConfig] interface {
	// GetStorageOptions request
	GetStorageOptions(ctx context.Context, projectID string, flavorID string, reqEditors ...RequestEditorFn) (*GetStorageOptionsResponse, error)
}

type GetStorageOptionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceGetFlavorStorageResponse
	JSON400      *InstanceError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetStorageOptionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetStorageOptionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetStorageOptions request returning *GetStorageOptionsResponse
func (c *ClientWithResponses[K]) GetStorageOptions(ctx context.Context, projectID string, flavorID string, reqEditors ...RequestEditorFn) (*GetStorageOptionsResponse, error) {
	rsp, err := c.GetStorageOptionsRaw(ctx, projectID, flavorID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetStorageOptionsResponse(rsp)
}

// ParseGetStorageOptionsResponse parses an HTTP response from a GetStorageOptions call
func (c *ClientWithResponses[K]) ParseGetStorageOptionsResponse(rsp *http.Response) (*GetStorageOptionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetStorageOptionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceGetFlavorStorageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	}

	return response, validate.ResponseObject(response)
}
