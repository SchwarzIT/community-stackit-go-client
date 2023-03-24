// Package project provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.17 DO NOT EDIT.
package project

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
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for ProjectState.
const (
	STATE_CREATED     ProjectState = "STATE_CREATED"
	STATE_CREATING    ProjectState = "STATE_CREATING"
	STATE_DELETING    ProjectState = "STATE_DELETING"
	STATE_FAILED      ProjectState = "STATE_FAILED"
	STATE_UNSPECIFIED ProjectState = "STATE_UNSPECIFIED"
)

// Defines values for RuntimeErrorCode.
const (
	SKE_API_SERVER_ERROR         RuntimeErrorCode = "SKE_API_SERVER_ERROR"
	SKE_ARGUS_INSTANCE_NOT_FOUND RuntimeErrorCode = "SKE_ARGUS_INSTANCE_NOT_FOUND"
	SKE_CONFIGURATION_PROBLEM    RuntimeErrorCode = "SKE_CONFIGURATION_PROBLEM"
	SKE_INFRA_ERROR              RuntimeErrorCode = "SKE_INFRA_ERROR"
	SKE_QUOTA_EXCEEDED           RuntimeErrorCode = "SKE_QUOTA_EXCEEDED"
	SKE_RATE_LIMITS              RuntimeErrorCode = "SKE_RATE_LIMITS"
	SKE_REMAINING_RESOURCES      RuntimeErrorCode = "SKE_REMAINING_RESOURCES"
	SKE_TMP_AUTH_ERROR           RuntimeErrorCode = "SKE_TMP_AUTH_ERROR"
	SKE_UNREADY_NODES            RuntimeErrorCode = "SKE_UNREADY_NODES"
	SKE_UNSPECIFIED              RuntimeErrorCode = "SKE_UNSPECIFIED"
)

// Project defines model for Project.
type Project struct {
	ProjectID *string       `json:"projectId,omitempty"`
	State     *ProjectState `json:"state,omitempty"`
}

// ProjectState defines model for ProjectState.
type ProjectState string

// RuntimeError defines model for RuntimeError.
type RuntimeError struct {
	// Code - Code:    "SKE_UNSPECIFIED"
	//   Message: "An error occurred. Please open a support ticket if this error persists."
	// - Code:    "SKE_TMP_AUTH_ERROR"
	//   Message: "Authentication failed. This is a temporary error. Please wait while the system recovers."
	// - Code:    "SKE_QUOTA_EXCEEDED"
	//   Message: "Your project's resource quotas are exhausted. Please make sure your quota is sufficient for the ordered cluster."
	// - Code:    "SKE_ARGUS_INSTANCE_NOT_FOUND"
	//   Message: "The provided Argus instance could not be found."
	// - Code:    "SKE_RATE_LIMITS"
	//   Message: "While provisioning your cluster, request rate limits where incurred. Please wait while the system recovers."
	// - Code:    "SKE_INFRA_ERROR"
	//   Message: "An error occurred with the underlying infrastructure. Please open a support ticket if this error persists."
	// - Code:    "SKE_REMAINING_RESOURCES"
	//   Message: "There are remaining Kubernetes resources in your cluster that prevent deletion. Please make sure to remove them."
	// - Code:    "SKE_CONFIGURATION_PROBLEM"
	//   Message: "A configuration error occurred. Please open a support ticket if this error persists."
	// - Code:    "SKE_UNREADY_NODES"
	//   Message: "Not all worker nodes are ready. Please open a support ticket if this error persists."
	// - Code:    "SKE_API_SERVER_ERROR"
	//   Message: "The Kubernetes API server is not reporting readiness. Please open a support ticket if this error persists."
	Code    *RuntimeErrorCode `json:"code,omitempty"`
	Details *string           `json:"details,omitempty"`
	Message *string           `json:"message,omitempty"`
}

// RuntimeErrorCode - Code:    "SKE_UNSPECIFIED"
//
//		Message: "An error occurred. Please open a support ticket if this error persists."
//	  - Code:    "SKE_TMP_AUTH_ERROR"
//	    Message: "Authentication failed. This is a temporary error. Please wait while the system recovers."
//	  - Code:    "SKE_QUOTA_EXCEEDED"
//	    Message: "Your project's resource quotas are exhausted. Please make sure your quota is sufficient for the ordered cluster."
//	  - Code:    "SKE_ARGUS_INSTANCE_NOT_FOUND"
//	    Message: "The provided Argus instance could not be found."
//	  - Code:    "SKE_RATE_LIMITS"
//	    Message: "While provisioning your cluster, request rate limits where incurred. Please wait while the system recovers."
//	  - Code:    "SKE_INFRA_ERROR"
//	    Message: "An error occurred with the underlying infrastructure. Please open a support ticket if this error persists."
//	  - Code:    "SKE_REMAINING_RESOURCES"
//	    Message: "There are remaining Kubernetes resources in your cluster that prevent deletion. Please make sure to remove them."
//	  - Code:    "SKE_CONFIGURATION_PROBLEM"
//	    Message: "A configuration error occurred. Please open a support ticket if this error persists."
//	  - Code:    "SKE_UNREADY_NODES"
//	    Message: "Not all worker nodes are ready. Please open a support ticket if this error persists."
//	  - Code:    "SKE_API_SERVER_ERROR"
//	    Message: "The Kubernetes API server is not reporting readiness. Please open a support ticket if this error persists."
type RuntimeErrorCode string

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
	// Delete request
	DeleteRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Get request
	GetRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Create request
	CreateRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client[K]) DeleteRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) GetRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client[K]) CreateRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDeleteRequest generates requests for Delete
func NewDeleteRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
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

	operationPath := fmt.Sprintf("/v1/projects/%s", pathParam0)
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

// NewCreateRequest generates requests for Create
func NewCreateRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", queryURL.String(), nil)
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
type ClientWithResponses struct {
	rawClientInterface
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient[K contracts.ClientFlowConfig](server string, httpClient contracts.ClientInterface[K]) *ClientWithResponses {
	return &ClientWithResponses{NewRawClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// Delete request
	Delete(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*DeleteResponse, error)

	// Get request
	Get(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetResponse, error)

	// Create request
	Create(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*CreateResponse, error)
}

type DeleteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON400      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r DeleteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Project
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
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

type CreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Project
	JSON400      *map[string]interface{}
	JSONDefault  *RuntimeError
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r CreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// Delete request returning *DeleteResponse
func (c *ClientWithResponses) Delete(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*DeleteResponse, error) {
	rsp, err := c.DeleteRaw(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseDeleteResponse(rsp)
}

// Get request returning *GetResponse
func (c *ClientWithResponses) Get(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.GetRaw(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetResponse(rsp)
}

// Create request returning *CreateResponse
func (c *ClientWithResponses) Create(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.CreateRaw(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateResponse(rsp)
}

// ParseDeleteResponse parses an HTTP response from a Delete call
func (c *ClientWithResponses) ParseDeleteResponse(rsp *http.Response) (*DeleteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseGetResponse parses an HTTP response from a Get call
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
		var dest Project
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseCreateResponse parses an HTTP response from a Create call
func (c *ClientWithResponses) ParseCreateResponse(rsp *http.Response) (*CreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Project
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSONDefault = &dest

	}

	return response, validate.ResponseObject(response)
}
