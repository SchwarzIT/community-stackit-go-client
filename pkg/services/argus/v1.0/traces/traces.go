// Package traces provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.23 DO NOT EDIT.
package traces

import (
	"bytes"
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

// Error defines model for Error.
type Error struct {
	Errors  *[]map[string]string `json:"errors,omitempty"`
	Message string               `json:"message"`
}

// Message defines model for Message.
type Message struct {
	Message string `json:"message"`
}

// PermissionDenied defines model for PermissionDenied.
type PermissionDenied struct {
	Detail string `json:"detail"`
}

// TraceConfig defines model for TraceConfig.
type TraceConfig struct {
	Retention string `json:"retention"`
}

// TracesConfigResponse defines model for TracesConfigResponse.
type TracesConfigResponse struct {
	Config  TraceConfig `json:"config"`
	Message string      `json:"message"`
}

// UpdateJSONBody defines parameters for Update.
type UpdateJSONBody struct {
	// Retention How long to keep the traces
	// `Additional Validators:`
	// * Should be a valid time string
	// * Should not be bigger than 30 days
	Retention string `json:"retention"`
}

// UpdateJSONRequestBody defines body for Update for application/json ContentType.
type UpdateJSONRequestBody UpdateJSONBody

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
	// List request
	ListRaw(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Update request with any body
	UpdateRawWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateRaw(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client[K]) ListRaw(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListRequest(ctx, c.Server, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) UpdateRawWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateRequestWithBody(ctx, c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) UpdateRaw(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateRequest(ctx, c.Server, projectID, instanceID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListRequest generates requests for List
func NewListRequest(ctx context.Context, server string, projectID string, instanceID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "instanceID", runtime.ParamLocationPath, instanceID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/traces-configs", pathParam0, pathParam1)
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

// NewUpdateRequest calls the generic Update builder with application/json body
func NewUpdateRequest(ctx context.Context, server string, projectID string, instanceID string, body UpdateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateRequestWithBody(ctx, server, projectID, instanceID, "application/json", bodyReader)
}

// NewUpdateRequestWithBody generates requests for Update with any type of body
func NewUpdateRequestWithBody(ctx context.Context, server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "instanceID", runtime.ParamLocationPath, instanceID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/traces-configs", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// List request
	List(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*ListResponse, error)

	// Update request with any body
	UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error)

	Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error)
}

type ListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *TracesConfigResponse
	JSON403      *PermissionDenied
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r ListResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON202      *Message
	JSON400      *Error
	JSON403      *PermissionDenied
	Error        error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r UpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// List request returning *ListResponse
func (c *ClientWithResponses[K]) List(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*ListResponse, error) {
	rsp, err := c.ListRaw(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListResponse(rsp)
}

// UpdateWithBody request with arbitrary body returning *UpdateResponse
func (c *ClientWithResponses[K]) UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.UpdateRawWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

func (c *ClientWithResponses[K]) Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.UpdateRaw(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

// ParseListResponse parses an HTTP response from a List call
func (c *ClientWithResponses[K]) ParseListResponse(rsp *http.Response) (*ListResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest TracesConfigResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest PermissionDenied
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	}

	return response, validate.ResponseObject(response)
}

// ParseUpdateResponse parses an HTTP response from a Update call
func (c *ClientWithResponses[K]) ParseUpdateResponse(rsp *http.Response) (*UpdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest Message
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest PermissionDenied
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	}

	return response, validate.ResponseObject(response)
}
