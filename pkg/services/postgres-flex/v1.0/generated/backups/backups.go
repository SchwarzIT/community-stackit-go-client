// Package backups provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.2.2 DO NOT EDIT.
package backups

import (
	"bytes"
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

// InstanceBackup defines model for instance.Backup.
type InstanceBackup struct {
	EndTime   *string            `json:"endTime,omitempty"`
	Error     *string            `json:"error,omitempty"`
	ID        *string            `json:"id,omitempty"`
	Labels    *[]string          `json:"labels,omitempty"`
	Name      *string            `json:"name,omitempty"`
	Options   *map[string]string `json:"options,omitempty"`
	StartTime *string            `json:"startTime,omitempty"`
}

// InstanceError defines model for instance.Error.
type InstanceError struct {
	Code    *int                 `json:"code,omitempty"`
	Fields  *map[string][]string `json:"fields,omitempty"`
	Message *string              `json:"message,omitempty"`
	Type    *string              `json:"type,omitempty"`
}

// InstanceGetBackupResponse defines model for instance.GetBackupResponse.
type InstanceGetBackupResponse struct {
	Item *InstanceBackup `json:"item,omitempty"`
}

// InstanceListBackupResponse defines model for instance.ListBackupResponse.
type InstanceListBackupResponse struct {
	Count *int              `json:"count,omitempty"`
	Items *[]InstanceBackup `json:"items,omitempty"`
}

// InstanceUpdateBackupScheduleRequest defines model for instance.UpdateBackupScheduleRequest.
type InstanceUpdateBackupScheduleRequest struct {
	BackupSchedule *string `json:"backupSchedule,omitempty"`
}

// PutInstanceBackupsJSONRequestBody defines body for PutInstanceBackups for application/json ContentType.
type PutInstanceBackupsJSONRequestBody = InstanceUpdateBackupScheduleRequest

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
	// GetInstanceBackups request
	GetInstanceBackups(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutInstanceBackups request with any body
	PutInstanceBackupsWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutInstanceBackups(ctx context.Context, projectID string, instanceID string, body PutInstanceBackupsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInstanceBackup request
	GetInstanceBackup(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetInstanceBackups(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstanceBackupsRequest(c.Server, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutInstanceBackupsWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutInstanceBackupsRequestWithBody(c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutInstanceBackups(ctx context.Context, projectID string, instanceID string, body PutInstanceBackupsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutInstanceBackupsRequest(c.Server, projectID, instanceID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInstanceBackup(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstanceBackupRequest(c.Server, projectID, instanceID, backupID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetInstanceBackupsRequest generates requests for GetInstanceBackups
func NewGetInstanceBackupsRequest(server string, projectID string, instanceID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/backups", pathParam0, pathParam1)
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

// NewPutInstanceBackupsRequest calls the generic PutInstanceBackups builder with application/json body
func NewPutInstanceBackupsRequest(server string, projectID string, instanceID string, body PutInstanceBackupsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutInstanceBackupsRequestWithBody(server, projectID, instanceID, "application/json", bodyReader)
}

// NewPutInstanceBackupsRequestWithBody generates requests for PutInstanceBackups with any type of body
func NewPutInstanceBackupsRequestWithBody(server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/backups", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetInstanceBackupRequest generates requests for GetInstanceBackup
func NewGetInstanceBackupRequest(server string, projectID string, instanceID string, backupID string) (*http.Request, error) {
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

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "backupID", runtime.ParamLocationPath, backupID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s/backups/%s", pathParam0, pathParam1, pathParam2)
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
	// GetInstanceBackups request
	GetInstanceBackupsWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*GetInstanceBackupsResponse, error)

	// PutInstanceBackups request with any body
	PutInstanceBackupsWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutInstanceBackupsResponse, error)

	PutInstanceBackupsWithResponse(ctx context.Context, projectID string, instanceID string, body PutInstanceBackupsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutInstanceBackupsResponse, error)

	// GetInstanceBackup request
	GetInstanceBackupWithResponse(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*GetInstanceBackupResponse, error)
}

type GetInstanceBackupsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceListBackupResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetInstanceBackupsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstanceBackupsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutInstanceBackupsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r PutInstanceBackupsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutInstanceBackupsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInstanceBackupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceGetBackupResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetInstanceBackupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstanceBackupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetInstanceBackupsWithResponse request returning *GetInstanceBackupsResponse
func (c *ClientWithResponses) GetInstanceBackupsWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*GetInstanceBackupsResponse, error) {
	rsp, err := c.GetInstanceBackups(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInstanceBackupsResponse(rsp)
}

// PutInstanceBackupsWithBodyWithResponse request with arbitrary body returning *PutInstanceBackupsResponse
func (c *ClientWithResponses) PutInstanceBackupsWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutInstanceBackupsResponse, error) {
	rsp, err := c.PutInstanceBackupsWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutInstanceBackupsResponse(rsp)
}

func (c *ClientWithResponses) PutInstanceBackupsWithResponse(ctx context.Context, projectID string, instanceID string, body PutInstanceBackupsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutInstanceBackupsResponse, error) {
	rsp, err := c.PutInstanceBackups(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutInstanceBackupsResponse(rsp)
}

// GetInstanceBackupWithResponse request returning *GetInstanceBackupResponse
func (c *ClientWithResponses) GetInstanceBackupWithResponse(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*GetInstanceBackupResponse, error) {
	rsp, err := c.GetInstanceBackup(ctx, projectID, instanceID, backupID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInstanceBackupResponse(rsp)
}

// ParseGetInstanceBackupsResponse parses an HTTP response from a GetInstanceBackupsWithResponse call
func ParseGetInstanceBackupsResponse(rsp *http.Response) (*GetInstanceBackupsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstanceBackupsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceListBackupResponse
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

	}

	return response, nil
}

// ParsePutInstanceBackupsResponse parses an HTTP response from a PutInstanceBackupsWithResponse call
func ParsePutInstanceBackupsResponse(rsp *http.Response) (*PutInstanceBackupsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutInstanceBackupsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseGetInstanceBackupResponse parses an HTTP response from a GetInstanceBackupWithResponse call
func ParseGetInstanceBackupResponse(rsp *http.Response) (*GetInstanceBackupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstanceBackupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceGetBackupResponse
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

	}

	return response, nil
}
