// Package credentials provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.0.2 DO NOT EDIT.
package credentials

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/do87/oapi-codegen/pkg/runtime"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for RuntimeErrorCode.
const (
	SKE_API_SERVER_ERROR      RuntimeErrorCode = "SKE_API_SERVER_ERROR"
	SKE_CONFIGURATION_PROBLEM RuntimeErrorCode = "SKE_CONFIGURATION_PROBLEM"
	SKE_INFRA_ERROR           RuntimeErrorCode = "SKE_INFRA_ERROR"
	SKE_QUOTA_EXCEEDED        RuntimeErrorCode = "SKE_QUOTA_EXCEEDED"
	SKE_RATE_LIMITS           RuntimeErrorCode = "SKE_RATE_LIMITS"
	SKE_REMAINING_RESOURCES   RuntimeErrorCode = "SKE_REMAINING_RESOURCES"
	SKE_TMP_AUTH_ERROR        RuntimeErrorCode = "SKE_TMP_AUTH_ERROR"
	SKE_UNREADY_NODES         RuntimeErrorCode = "SKE_UNREADY_NODES"
	SKE_UNSPECIFIED           RuntimeErrorCode = "SKE_UNSPECIFIED"
)

// Credentials defines model for Credentials.
type Credentials struct {
	CertificateAuthorityData *string `json:"certificateAuthorityData,omitempty"`

	// Kubeconfig This string contains the kubeconfig as yaml. If you want to directly get the yaml without any
	//  characters you can use the following command: curl -s 'api.stackit.cloud/ske/v1/projects/{projectId}/clusters/{clusterName}/credentials' |jq -r .kubeconfig
	Kubeconfig *string `json:"kubeconfig,omitempty"`
	Server     *string `json:"server,omitempty"`
	Token      *string `json:"token,omitempty"`
}

// RuntimeError defines model for RuntimeError.
type RuntimeError struct {
	// Code - Code:    "SKE_UNSPECIFIED"
	//   Message: "An error occurred. Please open a support ticket if this error persists."
	// - Code:    "SKE_TMP_AUTH_ERROR"
	//   Message: "Authentication failed. This is a temporary error. Please wait while the system recovers."
	// - Code:    "SKE_QUOTA_EXCEEDED"
	//   Message: "Your project's resource quotas are exhausted. Please make sure your quota is sufficient for the ordered cluster."
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

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer
}

// Creates a new Client, with reasonable defaults
func NewClient(server string, httpClient HttpRequestDoer) *Client {
	// create a client with sane default values
	client := Client{
		Server: server,
		Client: httpClient,
	}
	return &client
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetClusterCredentials request
	GetClusterCredentials(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TriggerClusterCredentialRotation request
	TriggerClusterCredentialRotation(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetClusterCredentials(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClusterCredentialsRequest(c.Server, projectID, clusterName)
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
	req, err := NewTriggerClusterCredentialRotationRequest(c.Server, projectID, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetClusterCredentialsRequest generates requests for GetClusterCredentials
func NewGetClusterCredentialsRequest(server string, projectID string, clusterName string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/clusters/%s/credentials", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTriggerClusterCredentialRotationRequest generates requests for TriggerClusterCredentialRotation
func NewTriggerClusterCredentialRotationRequest(server string, projectID string, clusterName string) (*http.Request, error) {
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

	req, err := http.NewRequest("POST", queryURL.String(), nil)
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
func NewClientWithResponses(server string, httpClient HttpRequestDoer) *ClientWithResponses {
	return &ClientWithResponses{NewClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetClusterCredentials request
	GetClusterCredentialsWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*GetClusterCredentialsResponse, error)

	// TriggerClusterCredentialRotation request
	TriggerClusterCredentialRotationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterCredentialRotationResponse, error)
}

type GetClusterCredentialsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Credentials
	JSON400      *map[string]interface{}
	JSON404      *map[string]interface{}
	JSONDefault  *RuntimeError
}

// Status returns HTTPResponse.Status
func (r GetClusterCredentialsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetClusterCredentialsResponse) StatusCode() int {
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

// GetClusterCredentialsWithResponse request returning *GetClusterCredentialsResponse
func (c *ClientWithResponses) GetClusterCredentialsWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*GetClusterCredentialsResponse, error) {
	rsp, err := c.GetClusterCredentials(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClusterCredentialsResponse(rsp)
}

// TriggerClusterCredentialRotationWithResponse request returning *TriggerClusterCredentialRotationResponse
func (c *ClientWithResponses) TriggerClusterCredentialRotationWithResponse(ctx context.Context, projectID string, clusterName string, reqEditors ...RequestEditorFn) (*TriggerClusterCredentialRotationResponse, error) {
	rsp, err := c.TriggerClusterCredentialRotation(ctx, projectID, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTriggerClusterCredentialRotationResponse(rsp)
}

// ParseGetClusterCredentialsResponse parses an HTTP response from a GetClusterCredentialsWithResponse call
func ParseGetClusterCredentialsResponse(rsp *http.Response) (*GetClusterCredentialsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetClusterCredentialsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Credentials
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseTriggerClusterCredentialRotationResponse parses an HTTP response from a TriggerClusterCredentialRotationWithResponse call
func ParseTriggerClusterCredentialRotationResponse(rsp *http.Response) (*TriggerClusterCredentialRotationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TriggerClusterCredentialRotationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RuntimeError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
