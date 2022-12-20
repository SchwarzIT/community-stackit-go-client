// Package instance provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.4.0 DO NOT EDIT.
package instance

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

// InstanceAcl defines model for instance.Acl.
type InstanceAcl struct {
	// Items TODO validating in api (middleware)
	Items *[]string `json:"items,omitempty"`
}

// InstanceCreateInstanceRequest defines model for instance.CreateInstanceRequest.
type InstanceCreateInstanceRequest struct {
	Acl            *InstanceAcl `json:"acl,omitempty"`
	BackupSchedule *string      `json:"backupSchedule,omitempty"`
	FlavorID       *string      `json:"flavorId,omitempty"`

	// Labels Following fields are not certain/clear
	Labels   *map[string]string `json:"labels,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Options  *map[string]string `json:"options,omitempty"`
	Replicas *int               `json:"replicas,omitempty"`
	Storage  *InstanceStorage   `json:"storage,omitempty"`
	Version  *string            `json:"version,omitempty"`
}

// InstanceCreateInstanceResponse defines model for instance.CreateInstanceResponse.
type InstanceCreateInstanceResponse struct {
	ID *string `json:"id,omitempty"`
}

// InstanceError defines model for instance.Error.
type InstanceError struct {
	Code    *int                 `json:"code,omitempty"`
	Fields  *map[string][]string `json:"fields,omitempty"`
	Message *string              `json:"message,omitempty"`
	Type    *string              `json:"type,omitempty"`
}

// InstanceFlavor defines model for instance.Flavor.
type InstanceFlavor struct {
	Cpu         *int    `json:"cpu,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Memory      *int    `json:"memory,omitempty"`
}

// InstanceGetInstanceResponse defines model for instance.GetInstanceResponse.
type InstanceGetInstanceResponse struct {
	Item *InstanceSingleInstance `json:"item,omitempty"`
}

// InstanceListInstance defines model for instance.ListInstance.
type InstanceListInstance struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}

// InstanceListInstanceResponse defines model for instance.ListInstanceResponse.
type InstanceListInstanceResponse struct {
	// Count TODO pagination ?
	Count *int                    `json:"count,omitempty"`
	Items *[]InstanceListInstance `json:"items,omitempty"`
}

// InstanceSingleInstance defines model for instance.SingleInstance.
type InstanceSingleInstance struct {
	Acl            *InstanceAcl       `json:"acl,omitempty"`
	BackupSchedule *string            `json:"backupSchedule,omitempty"`
	Flavor         *InstanceFlavor    `json:"flavor,omitempty"`
	ID             *string            `json:"id,omitempty"`
	Name           *string            `json:"name,omitempty"`
	Options        *map[string]string `json:"options,omitempty"`
	Replicas       *int               `json:"replicas,omitempty"`
	Status         *string            `json:"status,omitempty"`
	Storage        *InstanceStorage   `json:"storage,omitempty"`
	Version        *string            `json:"version,omitempty"`
}

// InstanceStorage defines model for instance.Storage.
type InstanceStorage struct {
	Class *string `json:"class,omitempty"`
	Size  *int    `json:"size,omitempty"`
}

// InstanceUpdateInstanceRequest defines model for instance.UpdateInstanceRequest.
type InstanceUpdateInstanceRequest struct {
	Acl            *InstanceAcl `json:"acl,omitempty"`
	BackupSchedule *string      `json:"backupSchedule,omitempty"`
	FlavorID       *string      `json:"flavorId,omitempty"`

	// Labels Following fields are not certain/clear
	Labels   *map[string]string `json:"labels,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Options  *map[string]string `json:"options,omitempty"`
	Replicas *int               `json:"replicas,omitempty"`
	Storage  *InstanceStorage   `json:"storage,omitempty"`
	Version  *string            `json:"version,omitempty"`
}

// InstanceUpdateInstanceResponse defines model for instance.UpdateInstanceResponse.
type InstanceUpdateInstanceResponse struct {
	Item *InstanceSingleInstance `json:"item,omitempty"`
}

// PostInstancesJSONRequestBody defines body for PostInstances for application/json ContentType.
type PostInstancesJSONRequestBody = InstanceCreateInstanceRequest

// PatchInstanceJSONRequestBody defines body for PatchInstance for application/json ContentType.
type PatchInstanceJSONRequestBody = InstanceUpdateInstanceRequest

// PutInstanceJSONRequestBody defines body for PutInstance for application/json ContentType.
type PutInstanceJSONRequestBody = InstanceUpdateInstanceRequest

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
	// GetInstances request
	GetInstances(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostInstances request with any body
	PostInstancesWithBody(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostInstances(ctx context.Context, projectID string, body PostInstancesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteInstance request
	DeleteInstance(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInstance request
	GetInstance(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchInstance request with any body
	PatchInstanceWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchInstance(ctx context.Context, projectID string, instanceID string, body PatchInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutInstance request with any body
	PutInstanceWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutInstance(ctx context.Context, projectID string, instanceID string, body PutInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetInstances(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstancesRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostInstancesWithBody(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostInstancesRequestWithBody(ctx, c.Server, projectID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostInstances(ctx context.Context, projectID string, body PostInstancesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostInstancesRequest(ctx, c.Server, projectID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteInstance(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteInstanceRequest(ctx, c.Server, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInstance(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstanceRequest(ctx, c.Server, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchInstanceWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchInstanceRequestWithBody(ctx, c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchInstance(ctx context.Context, projectID string, instanceID string, body PatchInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchInstanceRequest(ctx, c.Server, projectID, instanceID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutInstanceWithBody(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutInstanceRequestWithBody(ctx, c.Server, projectID, instanceID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutInstance(ctx context.Context, projectID string, instanceID string, body PutInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutInstanceRequest(ctx, c.Server, projectID, instanceID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetInstancesRequest generates requests for GetInstances
func NewGetInstancesRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances", pathParam0)
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

// NewPostInstancesRequest calls the generic PostInstances builder with application/json body
func NewPostInstancesRequest(ctx context.Context, server string, projectID string, body PostInstancesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostInstancesRequestWithBody(ctx, server, projectID, "application/json", bodyReader)
}

// NewPostInstancesRequestWithBody generates requests for PostInstances with any type of body
func NewPostInstancesRequestWithBody(ctx context.Context, server string, projectID string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances", pathParam0)
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

// NewDeleteInstanceRequest generates requests for DeleteInstance
func NewDeleteInstanceRequest(ctx context.Context, server string, projectID string, instanceID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s", pathParam0, pathParam1)
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

// NewGetInstanceRequest generates requests for GetInstance
func NewGetInstanceRequest(ctx context.Context, server string, projectID string, instanceID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s", pathParam0, pathParam1)
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

// NewPatchInstanceRequest calls the generic PatchInstance builder with application/json body
func NewPatchInstanceRequest(ctx context.Context, server string, projectID string, instanceID string, body PatchInstanceJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchInstanceRequestWithBody(ctx, server, projectID, instanceID, "application/json", bodyReader)
}

// NewPatchInstanceRequestWithBody generates requests for PatchInstance with any type of body
func NewPatchInstanceRequestWithBody(ctx context.Context, server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPutInstanceRequest calls the generic PutInstance builder with application/json body
func NewPutInstanceRequest(ctx context.Context, server string, projectID string, instanceID string, body PutInstanceJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutInstanceRequestWithBody(ctx, server, projectID, instanceID, "application/json", bodyReader)
}

// NewPutInstanceRequestWithBody generates requests for PutInstance with any type of body
func NewPutInstanceRequestWithBody(ctx context.Context, server string, projectID string, instanceID string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/instances/%s", pathParam0, pathParam1)
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
	// GetInstances request
	GetInstancesWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetInstancesResponse, error)

	// PostInstances request with any body
	PostInstancesWithBodyWithResponse(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostInstancesResponse, error)

	PostInstancesWithResponse(ctx context.Context, projectID string, body PostInstancesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostInstancesResponse, error)

	// DeleteInstance request
	DeleteInstanceWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*DeleteInstanceResponse, error)

	// GetInstance request
	GetInstanceWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*GetInstanceResponse, error)

	// PatchInstance request with any body
	PatchInstanceWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchInstanceResponse, error)

	PatchInstanceWithResponse(ctx context.Context, projectID string, instanceID string, body PatchInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchInstanceResponse, error)

	// PutInstance request with any body
	PutInstanceWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutInstanceResponse, error)

	PutInstanceWithResponse(ctx context.Context, projectID string, instanceID string, body PutInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*PutInstanceResponse, error)
}

type GetInstancesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceListInstanceResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetInstancesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstancesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostInstancesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceCreateInstanceResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r PostInstancesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostInstancesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteInstanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r DeleteInstanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteInstanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInstanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceGetInstanceResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r GetInstanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchInstanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceUpdateInstanceResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r PatchInstanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchInstanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutInstanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *InstanceUpdateInstanceResponse
	JSON400      *InstanceError
	HasError     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r PutInstanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutInstanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetInstancesWithResponse request returning *GetInstancesResponse
func (c *ClientWithResponses) GetInstancesWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*GetInstancesResponse, error) {
	rsp, err := c.GetInstances(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetInstancesResponse(rsp)
}

// PostInstancesWithBodyWithResponse request with arbitrary body returning *PostInstancesResponse
func (c *ClientWithResponses) PostInstancesWithBodyWithResponse(ctx context.Context, projectID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostInstancesResponse, error) {
	rsp, err := c.PostInstancesWithBody(ctx, projectID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePostInstancesResponse(rsp)
}

func (c *ClientWithResponses) PostInstancesWithResponse(ctx context.Context, projectID string, body PostInstancesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostInstancesResponse, error) {
	rsp, err := c.PostInstances(ctx, projectID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePostInstancesResponse(rsp)
}

// DeleteInstanceWithResponse request returning *DeleteInstanceResponse
func (c *ClientWithResponses) DeleteInstanceWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*DeleteInstanceResponse, error) {
	rsp, err := c.DeleteInstance(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseDeleteInstanceResponse(rsp)
}

// GetInstanceWithResponse request returning *GetInstanceResponse
func (c *ClientWithResponses) GetInstanceWithResponse(ctx context.Context, projectID string, instanceID string, reqEditors ...RequestEditorFn) (*GetInstanceResponse, error) {
	rsp, err := c.GetInstance(ctx, projectID, instanceID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetInstanceResponse(rsp)
}

// PatchInstanceWithBodyWithResponse request with arbitrary body returning *PatchInstanceResponse
func (c *ClientWithResponses) PatchInstanceWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchInstanceResponse, error) {
	rsp, err := c.PatchInstanceWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePatchInstanceResponse(rsp)
}

func (c *ClientWithResponses) PatchInstanceWithResponse(ctx context.Context, projectID string, instanceID string, body PatchInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchInstanceResponse, error) {
	rsp, err := c.PatchInstance(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePatchInstanceResponse(rsp)
}

// PutInstanceWithBodyWithResponse request with arbitrary body returning *PutInstanceResponse
func (c *ClientWithResponses) PutInstanceWithBodyWithResponse(ctx context.Context, projectID string, instanceID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutInstanceResponse, error) {
	rsp, err := c.PutInstanceWithBody(ctx, projectID, instanceID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePutInstanceResponse(rsp)
}

func (c *ClientWithResponses) PutInstanceWithResponse(ctx context.Context, projectID string, instanceID string, body PutInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*PutInstanceResponse, error) {
	rsp, err := c.PutInstance(ctx, projectID, instanceID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParsePutInstanceResponse(rsp)
}

// ParseGetInstancesResponse parses an HTTP response from a GetInstancesWithResponse call
func (c *ClientWithResponses) ParseGetInstancesResponse(rsp *http.Response) (*GetInstancesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstancesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceListInstanceResponse
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

// ParsePostInstancesResponse parses an HTTP response from a PostInstancesWithResponse call
func (c *ClientWithResponses) ParsePostInstancesResponse(rsp *http.Response) (*PostInstancesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostInstancesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceCreateInstanceResponse
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

// ParseDeleteInstanceResponse parses an HTTP response from a DeleteInstanceWithResponse call
func (c *ClientWithResponses) ParseDeleteInstanceResponse(rsp *http.Response) (*DeleteInstanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteInstanceResponse{
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

// ParseGetInstanceResponse parses an HTTP response from a GetInstanceWithResponse call
func (c *ClientWithResponses) ParseGetInstanceResponse(rsp *http.Response) (*GetInstanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceGetInstanceResponse
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

// ParsePatchInstanceResponse parses an HTTP response from a PatchInstanceWithResponse call
func (c *ClientWithResponses) ParsePatchInstanceResponse(rsp *http.Response) (*PatchInstanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PatchInstanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceUpdateInstanceResponse
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

// ParsePutInstanceResponse parses an HTTP response from a PutInstanceWithResponse call
func (c *ClientWithResponses) ParsePutInstanceResponse(rsp *http.Response) (*PutInstanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutInstanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest InstanceUpdateInstanceResponse
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
