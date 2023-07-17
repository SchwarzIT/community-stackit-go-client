// Package backup provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/stackit-client-generator version v0.0.3 DO NOT EDIT.
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

	contracts "github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/helpers/runtime"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
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

// InstanceGetBackupResponse defines model for instance.GetBackupResponse.
type InstanceGetBackupResponse struct {
	Item *InstanceBackup `json:"item,omitempty"`
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
	Client contracts.BaseClientInterface
}

// NewRawClient Creates a new Client, with reasonable defaults
func NewRawClient(server string, httpClient contracts.BaseClientInterface) *Client {
	// create a client with sane default values
	client := Client{
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

	// Get request
	GetRaw(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateClone request with any body
	CreateCloneRawWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateCloneRaw(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateRestore request with any body
	CreateRestoreRawWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateRestoreRaw(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListRaw(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) UpdateRawWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) UpdateRaw(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) GetRaw(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) CreateCloneRawWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) CreateCloneRaw(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) CreateRestoreRawWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) CreateRestoreRaw(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

// ClientWithResponses builds on rawClientInterface to offer response payloads
type ClientWithResponses struct {
	rawClientInterface
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient(server string, httpClient contracts.BaseClientInterface) *ClientWithResponses {
	return &ClientWithResponses{NewRawClient(server, httpClient)}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// List request
	List(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*ListResponse, error)

	// Update request with any body
	UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error)

	Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error)

	// Get request
	Get(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*GetResponse, error)

	// CreateClone request with any body
	CreateCloneWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error)

	CreateClone(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error)

	// CreateRestore request with any body
	CreateRestoreWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error)

	CreateRestore(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error)
}

type ListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceListBackupResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
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
	JSON200      *OpsmanagerUpdateScheduleRequest
	JSON400      *InstanceError
	JSON404      *InstanceError
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

type GetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceGetBackupResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
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

type CreateCloneResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON202      *InstanceCreateCloneInstanceResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
	Error        error // Aggregated error
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
	JSON202      *InstanceCreateRestoreInstanceResponse
	JSON400      *InstanceError
	JSON404      *InstanceError
	Error        error // Aggregated error
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

// List request returning *ListResponse
func (c *ClientWithResponses) List(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*ListResponse, error) {
	rsp, err := c.ListRaw(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListResponse(rsp)
}

// UpdateWithBody request with arbitrary body returning *UpdateResponse
func (c *ClientWithResponses) UpdateWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.UpdateRawWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

func (c *ClientWithResponses) Update(ctx context.Context, projectID string, instanceID string, body UpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateResponse, error) {
	rsp, err := c.UpdateRaw(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseUpdateResponse(rsp)
}

// Get request returning *GetResponse
func (c *ClientWithResponses) Get(ctx context.Context, projectID string, instanceID string, backupID string, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.GetRaw(ctx, projectID, instanceID, backupID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetResponse(rsp)
}

// CreateCloneWithBody request with arbitrary body returning *CreateCloneResponse
func (c *ClientWithResponses) CreateCloneWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error) {
	rsp, err := c.CreateCloneRawWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateCloneResponse(rsp)
}

func (c *ClientWithResponses) CreateClone(ctx context.Context, projectID string, instanceID string, body CreateCloneJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCloneResponse, error) {
	rsp, err := c.CreateCloneRaw(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateCloneResponse(rsp)
}

// CreateRestoreWithBody request with arbitrary body returning *CreateRestoreResponse
func (c *ClientWithResponses) CreateRestoreWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error) {
	rsp, err := c.CreateRestoreRawWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateRestoreResponse(rsp)
}

func (c *ClientWithResponses) CreateRestore(ctx context.Context, projectID string, instanceID string, body CreateRestoreJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateRestoreResponse, error) {
	rsp, err := c.CreateRestoreRaw(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateRestoreResponse(rsp)
}

// ParseListResponse parses an HTTP response from a List call
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
	response.Error = validate.DefaultResponseErrorHandler(rsp)

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

	return response, validate.ResponseObject(response)
}

// ParseUpdateResponse parses an HTTP response from a Update call
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
	response.Error = validate.DefaultResponseErrorHandler(rsp)

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
		var dest InstanceGetBackupResponse
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

	return response, validate.ResponseObject(response)
}

// ParseCreateCloneResponse parses an HTTP response from a CreateClone call
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
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest InstanceCreateCloneInstanceResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

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

	return response, validate.ResponseObject(response)
}

// ParseCreateRestoreResponse parses an HTTP response from a CreateRestore call
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
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest InstanceCreateRestoreInstanceResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON202 = &dest

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

	return response, validate.ResponseObject(response)
}
