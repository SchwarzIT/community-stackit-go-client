// Package plans provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.1 DO NOT EDIT.
package plans

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
	openapi_types "github.com/do87/oapi-codegen/pkg/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Offerings defines model for Offerings.
type Offerings struct {
	Description      string        `json:"description"`
	DocumentationURL string        `json:"documentationUrl"`
	ImageURL         string        `json:"imageUrl"`
	Name             string        `json:"name"`
	Plans            []PlanModelUI `json:"plans"`
	Tags             []string      `json:"tags"`
}

// PermissionDenied defines model for PermissionDenied.
type PermissionDenied struct {
	Detail string `json:"detail"`
}

// Plan defines model for Plan.
type Plan struct {
	Message string        `json:"message"`
	Plans   []PlanModelUI `json:"plans"`
}

// PlanModelUI defines model for PlanModelUI.
type PlanModelUI struct {
	AlertMatchers           int                `json:"alertMatchers"`
	AlertReceivers          int                `json:"alertReceivers"`
	AlertRules              int                `json:"alertRules"`
	Amount                  *float32           `json:"amount,omitempty"`
	BucketSize              int                `json:"bucketSize"`
	Description             *string            `json:"description,omitempty"`
	GrafanaGlobalDashboards int                `json:"grafanaGlobalDashboards"`
	GrafanaGlobalOrgs       int                `json:"grafanaGlobalOrgs"`
	GrafanaGlobalSessions   int                `json:"grafanaGlobalSessions"`
	GrafanaGlobalUsers      int                `json:"grafanaGlobalUsers"`
	ID                      openapi_types.UUID `json:"id"`
	IsFree                  *bool              `json:"isFree,omitempty"`
	IsPublic                *bool              `json:"isPublic,omitempty"`
	LogsAlert               int                `json:"logsAlert"`
	LogsStorage             int                `json:"logsStorage"`
	Name                    *string            `json:"name,omitempty"`
	PlanID                  openapi_types.UUID `json:"planId"`
	SamplesPerScrape        int                `json:"samplesPerScrape"`
	Schema                  *string            `json:"schema,omitempty"`
	TargetNumber            int                `json:"targetNumber"`
	TracesStorage           int                `json:"tracesStorage"`
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
	// ListOfferings request
	ListOfferings(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListPlans request
	ListPlans(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListOfferings(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListOfferingsRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListPlans(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListPlansRequest(ctx, c.Server, projectID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListOfferingsRequest generates requests for ListOfferings
func NewListOfferingsRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/offerings", pathParam0)
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

// NewListPlansRequest generates requests for ListPlans
func NewListPlansRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/projects/%s/plans", pathParam0)
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
	// ListOfferings request
	ListOfferingsWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListOfferingsResponse, error)

	// ListPlans request
	ListPlansWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListPlansResponse, error)
}

type ListOfferingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Offerings
	JSON403      *PermissionDenied
	Error     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r ListOfferingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListOfferingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListPlansResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Plan
	JSON403      *PermissionDenied
	Error     error // Aggregated error
}

// Status returns HTTPResponse.Status
func (r ListPlansResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListPlansResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListOfferingsWithResponse request returning *ListOfferingsResponse
func (c *ClientWithResponses) ListOfferingsWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListOfferingsResponse, error) {
	rsp, err := c.ListOfferings(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListOfferingsResponse(rsp)
}

// ListPlansWithResponse request returning *ListPlansResponse
func (c *ClientWithResponses) ListPlansWithResponse(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListPlansResponse, error) {
	rsp, err := c.ListPlans(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListPlansResponse(rsp)
}

// ParseListOfferingsResponse parses an HTTP response from a ListOfferingsWithResponse call
func (c *ClientWithResponses) ParseListOfferingsResponse(rsp *http.Response) (*ListOfferingsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListOfferingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Offerings
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

	return response, nil
}

// ParseListPlansResponse parses an HTTP response from a ListPlansWithResponse call
func (c *ClientWithResponses) ParseListPlansResponse(rsp *http.Response) (*ListPlansResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListPlansResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	response.Error = validate.DefaultResponseErrorHandler(rsp)

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Plan
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

	return response, nil
}
