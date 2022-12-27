// Package storage provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.0 DO NOT EDIT.
package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/do87/oapi-codegen/pkg/runtime"
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
	// GetStoragesFlavor request
	GetStoragesFlavor(ctx context.Context, projectID string, flavor string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetStoragesFlavor(ctx context.Context, projectID string, flavor string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetStoragesFlavorRequest(ctx, c.Server, projectID, flavor)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetStoragesFlavorRequest generates requests for GetStoragesFlavor
func NewGetStoragesFlavorRequest(ctx context.Context, server string, projectID string, flavor string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "flavor", runtime.ParamLocationPath, flavor)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/projects/%s/storages/%s", pathParam0, pathParam1)
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
	// GetStoragesFlavor request
	GetStoragesFlavorWithResponse(ctx context.Context, projectID string, flavor string, reqEditors ...RequestEditorFn) (*GetStoragesFlavorResponse, error)
}

type GetStoragesFlavorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceGetFlavorStorageResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetStoragesFlavorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetStoragesFlavorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetStoragesFlavorWithResponse request returning *GetStoragesFlavorResponse
func (c *ClientWithResponses) GetStoragesFlavorWithResponse(ctx context.Context, projectID string, flavor string, reqEditors ...RequestEditorFn) (*GetStoragesFlavorResponse, error) {
	rsp, err := c.GetStoragesFlavor(ctx, projectID, flavor, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetStoragesFlavorResponse(rsp)
}

// ParseGetStoragesFlavorResponse parses an HTTP response from a GetStoragesFlavorWithResponse call
func (c *ClientWithResponses) ParseGetStoragesFlavorResponse(rsp *http.Response) (*GetStoragesFlavorResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetStoragesFlavorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceGetFlavorStorageResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}
