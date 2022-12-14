// Package backup provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.1 DO NOT EDIT.
package backup

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

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
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

// InstanceCreateCloneInstanceRequest defines model for instance.CreateCloneInstanceRequest.
type InstanceCreateCloneInstanceRequest struct {
	InstanceID *string `json:"instanceId,omitempty"`
}

// InstanceCreateCloneInstanceResponse defines model for instance.CreateCloneInstanceResponse.
type InstanceCreateCloneInstanceResponse struct {
	InstanceID *string `json:"instanceId,omitempty"`
}

// InstanceCreateRestoreInstanceRequest defines model for instance.CreateRestoreInstanceRequest.
type InstanceCreateRestoreInstanceRequest struct {
	BackupID   *string `json:"backupId,omitempty"`
	InstanceID *string `json:"instanceId,omitempty"`
}

// InstanceCreateRestoreInstanceResponse defines model for instance.CreateRestoreInstanceResponse.
type InstanceCreateRestoreInstanceResponse struct {
	InstanceID *string `json:"instanceId,omitempty"`
}

// InstanceError defines model for instance.Error.
type InstanceError struct {
	Code    *int                 `json:"code,omitempty"`
	Fields  *map[string][]string `json:"fields,omitempty"`
	Message *string              `json:"message,omitempty"`
	Type    *string              `json:"type,omitempty"`
}

// InstanceListBackupResponse defines model for instance.ListBackupResponse.
type InstanceListBackupResponse struct {
	Count *int              `json:"count,omitempty"`
	Items *[]InstanceBackup `json:"items,omitempty"`
}

// OpsmanagerUpdateScheduleRequest defines model for opsmanager.UpdateScheduleRequest.
type OpsmanagerUpdateScheduleRequest struct {
	BackupSchedule                 *string `json:"backupSchedule,omitempty"`
	DailySnapshotRetentionDays     *int    `json:"dailySnapshotRetentionDays,omitempty"`
	MonthlySnapshotRetentionMonths *int    `json:"monthlySnapshotRetentionMonths,omitempty"`
	PointInTimeWindowHours         *int    `json:"pointInTimeWindowHours,omitempty"`
	SnapshotRetentionDays          *int    `json:"snapshotRetentionDays,omitempty"`
	WeeklySnapshotRetentionWeeks   *int    `json:"weeklySnapshotRetentionWeeks,omitempty"`
}

// UpdateJSONRequestBody defines body for Update for application/json ContentType.
type UpdateJSONRequestBody = OpsmanagerUpdateScheduleRequest

// CreateCloneJSONRequestBody defines body for CreateClone for application/json ContentType.
type CreateCloneJSONRequestBody = InstanceCreateCloneInstanceRequest

// CreateRestoreJSONRequestBody defines body for CreateRestore for application/json ContentType.
type CreateRestoreJSONRequestBody = InstanceCreateRestoreInstanceRequest

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
	// List request
	List(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Update request with any body
	UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Get request
	Get(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateClone request with any body
	CreateCloneWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateClone(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateRestore request with any body
	CreateRestoreWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateRestore(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) List(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) Get(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequest(ctx, c.Server, projectID, instanceID, backupID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateCloneWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateCloneRequestWithBody(ctx, c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateClone(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateCloneRequest(ctx, c.Server, projectID, instanceID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateRestoreWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRestoreRequestWithBody(ctx, c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateRestore(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRestoreRequest(ctx, c.Server, projectID, instanceID, body)
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

	operationPath := fmt.Sprintf("/projects/%s/instances/%s/backups", pathParam0, pathParam1)
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

	operationPath := fmt.Sprintf("/projects/%s/instances/%s/backups", pathParam0, pathParam1)
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

// NewGetRequest generates requests for Get
func NewGetRequest(ctx context.Context, server string, projectID string, instanceID string, backupID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/projects/%s/instances/%s/backups/%s", pathParam0, pathParam1, pathParam2)
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

// NewCreateCloneRequest calls the generic CreateClone builder with application/json body
func NewCreateCloneRequest(ctx context.Context, server string, projectID string, instanceID string, body CreateCloneJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateCloneRequestWithBody(ctx, server, projectID, instanceID, "application/json", bodyReader)
}

// NewCreateCloneRequestWithBody generates requests for CreateClone with any type of body
func NewCreateCloneRequestWithBody(ctx context.Context, server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/projects/%s/instances/%s/clone", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewCreateRestoreRequest calls the generic CreateRestore builder with application/json body
func NewCreateRestoreRequest(ctx context.Context, server string, projectID string, instanceID string, body CreateRestoreJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateRestoreRequestWithBody(ctx, server, projectID, instanceID, "application/json", bodyReader)
}

// NewCreateRestoreRequestWithBody generates requests for CreateRestore with any type of body
func NewCreateRestoreRequestWithBody(ctx context.Context, server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/projects/%s/instances/%s/restore", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// List request
	ListWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*ListResponse, error)

	// Update request with any body
	UpdateWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error)

	UpdateWithResponse(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error)

	// Get request
	GetWithResponse(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*GetResponse, error)

	// CreateClone request with any body
	CreateCloneWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error)

	CreateCloneWithResponse(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error)

	// CreateRestore request with any body
	CreateRestoreWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error)

	CreateRestoreWithResponse(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error)
}

type ListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceListBackupResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
	HasError     error // Aggregated error
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
	JSON200      *OpsmanagerUpdateScheduleRequest
	JSON400      *InstanceError
	JSON404      *InstanceError
	HasError     error // Aggregated error
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

type GetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceBackup
	JSON400      *InstanceError
	JSON404      *InstanceError
	HasError     error // Aggregated error
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

type CreateCloneResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceCreateCloneInstanceResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r CreateCloneResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateCloneResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateRestoreResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceCreateRestoreInstanceResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r CreateRestoreResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateRestoreResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListWithResponse request returning *ListResponse
func (c *ClientWithResponses) ListWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*ListResponse, error) {
	rsp, err := c.List(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListResponse(rsp)
}

// UpdateWithBodyWithResponse request with arbitrary body returning *UpdateResponse
func (c *ClientWithResponses) UpdateWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.UpdateWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

func (c *ClientWithResponses) UpdateWithResponse(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.Update(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

// GetWithResponse request returning *GetResponse
func (c *ClientWithResponses) GetWithResponse(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.Get(ctx, projectID, instanceID, backupID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetResponse(rsp)
}

// CreateCloneWithBodyWithResponse request with arbitrary body returning *CreateCloneResponse
func (c *ClientWithResponses) CreateCloneWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error) {
	rsp, err := c.CreateCloneWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateCloneResponse(rsp)
}

func (c *ClientWithResponses) CreateCloneWithResponse(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error) {
	rsp, err := c.CreateClone(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateCloneResponse(rsp)
}

// CreateRestoreWithBodyWithResponse request with arbitrary body returning *CreateRestoreResponse
func (c *ClientWithResponses) CreateRestoreWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error) {
	rsp, err := c.CreateRestoreWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateRestoreResponse(rsp)
}

func (c *ClientWithResponses) CreateRestoreWithResponse(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error) {
	rsp, err := c.CreateRestore(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateRestoreResponse(rsp)
}

// ParseListResponse parses an HTTP response from a ListWithResponse call
func (c *ClientWithResponses) ParseListResponse(rsp *http.Response) (*ListResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceListBackupResponse
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateResponse parses an HTTP response from a UpdateWithResponse call
func (c *ClientWithResponses) ParseUpdateResponse(rsp *http.Response) (*UpdateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpsmanagerUpdateScheduleRequest
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
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
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceBackup
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseCreateCloneResponse parses an HTTP response from a CreateCloneWithResponse call
func (c *ClientWithResponses) ParseCreateCloneResponse(rsp *http.Response) (*CreateCloneResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateCloneResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceCreateCloneInstanceResponse
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseCreateRestoreResponse parses an HTTP response from a CreateRestoreWithResponse call
func (c *ClientWithResponses) ParseCreateRestoreResponse(rsp *http.Response) (*CreateRestoreResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateRestoreResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.HasError = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceCreateRestoreInstanceResponse
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

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest InstanceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	}

	return response, nil
}
