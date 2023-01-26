// Package operation provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.1 DO NOT EDIT.
package operation

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

const (
	BearerAuthScopes = "bearerAuth.Scopes"
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
type RuntimeErrorCode string

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
	// TriggerClusterHibernation request
	TriggerClusterHibernation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerClusterMaintenance request
	TriggerClusterMaintenance(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerClusterReconciliation request
	TriggerClusterReconciliation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerClusterCredentialRotation request
	TriggerClusterCredentialRotation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerClusterWakeup request
	TriggerClusterWakeup(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) TriggerClusterHibernation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerClusterHibernationRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerClusterMaintenance(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerClusterMaintenanceRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerClusterReconciliation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerClusterReconciliationRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerClusterCredentialRotation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerClusterCredentialRotationRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TriggerClusterWakeup(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTriggerClusterWakeupRequest(ctx, c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewTriggerClusterHibernationRequest generates requests for TriggerClusterHibernation
func NewTriggerClusterHibernationRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/hibernate", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerClusterMaintenanceRequest generates requests for TriggerClusterMaintenance
func NewTriggerClusterMaintenanceRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/maintenance", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerClusterReconciliationRequest generates requests for TriggerClusterReconciliation
func NewTriggerClusterReconciliationRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/reconcile", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerClusterCredentialRotationRequest generates requests for TriggerClusterCredentialRotation
func NewTriggerClusterCredentialRotationRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/rotate-credentials", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerClusterWakeupRequest generates requests for TriggerClusterWakeup
func NewTriggerClusterWakeupRequest(ctx context.Context, server string, projectID string, clusterName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/wakeup", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), nil)
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
	// TriggerClusterHibernation request
	TriggerClusterHibernationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterHibernationResponse, error)

	// TriggerClusterMaintenance request
	TriggerClusterMaintenanceWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterMaintenanceResponse, error)

	// TriggerClusterReconciliation request
	TriggerClusterReconciliationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterReconciliationResponse, error)

	// TriggerClusterCredentialRotation request
	TriggerClusterCredentialRotationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterCredentialRotationResponse, error)

	// TriggerClusterWakeup request
	TriggerClusterWakeupWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterWakeupResponse, error)
}

type TriggerClusterHibernationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerClusterHibernationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerClusterHibernationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerClusterMaintenanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerClusterMaintenanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerClusterMaintenanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerClusterReconciliationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerClusterReconciliationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerClusterReconciliationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerClusterCredentialRotationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerClusterCredentialRotationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerClusterCredentialRotationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TriggerClusterWakeupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON202      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r TriggerClusterWakeupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TriggerClusterWakeupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// TriggerClusterHibernationWithResponse request returning *TriggerClusterHibernationResponse
func (c *ClientWithResponses) TriggerClusterHibernationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterHibernationResponse, error) {
	rsp, err := c.TriggerClusterHibernation(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerClusterHibernationResponse(rsp)
}

// TriggerClusterMaintenanceWithResponse request returning *TriggerClusterMaintenanceResponse
func (c *ClientWithResponses) TriggerClusterMaintenanceWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterMaintenanceResponse, error) {
	rsp, err := c.TriggerClusterMaintenance(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerClusterMaintenanceResponse(rsp)
}

// TriggerClusterReconciliationWithResponse request returning *TriggerClusterReconciliationResponse
func (c *ClientWithResponses) TriggerClusterReconciliationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterReconciliationResponse, error) {
	rsp, err := c.TriggerClusterReconciliation(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerClusterReconciliationResponse(rsp)
}

// TriggerClusterCredentialRotationWithResponse request returning *TriggerClusterCredentialRotationResponse
func (c *ClientWithResponses) TriggerClusterCredentialRotationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterCredentialRotationResponse, error) {
	rsp, err := c.TriggerClusterCredentialRotation(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerClusterCredentialRotationResponse(rsp)
}

// TriggerClusterWakeupWithResponse request returning *TriggerClusterWakeupResponse
func (c *ClientWithResponses) TriggerClusterWakeupWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterWakeupResponse, error) {
	rsp, err := c.TriggerClusterWakeup(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseTriggerClusterWakeupResponse(rsp)
}

// ParseTriggerClusterHibernationResponse parses an HTTP response from a TriggerClusterHibernationWithResponse call
func (c *ClientWithResponses) ParseTriggerClusterHibernationResponse(rsp *http.Response) (*TriggerClusterHibernationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerClusterHibernationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

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

	return response, nil
}

// ParseTriggerClusterMaintenanceResponse parses an HTTP response from a TriggerClusterMaintenanceWithResponse call
func (c *ClientWithResponses) ParseTriggerClusterMaintenanceResponse(rsp *http.Response) (*TriggerClusterMaintenanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerClusterMaintenanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

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

	return response, nil
}

// ParseTriggerClusterReconciliationResponse parses an HTTP response from a TriggerClusterReconciliationWithResponse call
func (c *ClientWithResponses) ParseTriggerClusterReconciliationResponse(rsp *http.Response) (*TriggerClusterReconciliationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerClusterReconciliationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

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

	return response, nil
}

// ParseTriggerClusterCredentialRotationResponse parses an HTTP response from a TriggerClusterCredentialRotationWithResponse call
func (c *ClientWithResponses) ParseTriggerClusterCredentialRotationResponse(rsp *http.Response) (*TriggerClusterCredentialRotationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerClusterCredentialRotationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

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

	return response, nil
}

// ParseTriggerClusterWakeupResponse parses an HTTP response from a TriggerClusterWakeupWithResponse call
func (c *ClientWithResponses) ParseTriggerClusterWakeupResponse(rsp *http.Response) (*TriggerClusterWakeupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerClusterWakeupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

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

	return response, nil
}
