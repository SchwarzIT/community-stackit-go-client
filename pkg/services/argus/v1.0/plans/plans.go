// Package plans provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.17 DO NOT EDIT.
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

	contracts "github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/internal/helpers/runtime"
	openapiTypes "github.com/SchwarzIT/community-stackit-go-client/internal/helpers/types"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
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
	AlertMatchers           int               `json:"alertMatchers"`
	AlertReceivers          int               `json:"alertReceivers"`
	AlertRules              int               `json:"alertRules"`
	Amount                  *float32          `json:"amount,omitempty"`
	BucketSize              int               `json:"bucketSize"`
	Description             *string           `json:"description,omitempty"`
	GrafanaGlobalDashboards int               `json:"grafanaGlobalDashboards"`
	GrafanaGlobalOrgs       int               `json:"grafanaGlobalOrgs"`
	GrafanaGlobalSessions   int               `json:"grafanaGlobalSessions"`
	GrafanaGlobalUsers      int               `json:"grafanaGlobalUsers"`
	ID                      openapiTypes.UUID `json:"id"`
	IsFree                  *bool             `json:"isFree,omitempty"`
	IsPublic                *bool             `json:"isPublic,omitempty"`
	LogsAlert               int               `json:"logsAlert"`
	LogsStorage             int               `json:"logsStorage"`
	Name                    *string           `json:"name,omitempty"`
	PlanID                  openapiTypes.UUID `json:"planId"`
	SamplesPerScrape        int               `json:"samplesPerScrape"`
	Schema                  *string           `json:"schema,omitempty"`
	TargetNumber            int               `json:"targetNumber"`
	TracesStorage           int               `json:"tracesStorage"`
}

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
	// ListOfferings request
	ListOfferingsRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListPlans request
	ListPlansRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client[K]) ListOfferingsRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client[K]) ListPlansRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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
	// ListOfferings request
	ListOfferings(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListOfferingsResponse, error)

	// ListPlans request
	ListPlans(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListPlansResponse, error)
}

type ListOfferingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Offerings
	JSON403      *PermissionDenied
	Error        error // Aggregated error
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
	Error        error // Aggregated error
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

// ListOfferings request returning *ListOfferingsResponse
func (c *ClientWithResponses) ListOfferings(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListOfferingsResponse, error) {
	rsp, err := c.ListOfferingsRaw(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListOfferingsResponse(rsp)
}

// ListPlans request returning *ListPlansResponse
func (c *ClientWithResponses) ListPlans(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListPlansResponse, error) {
	rsp, err := c.ListPlansRaw(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListPlansResponse(rsp)
}

// ParseListOfferingsResponse parses an HTTP response from a ListOfferings call
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

	return response, validate.ResponseObject(response)
}

// ParseListPlansResponse parses an HTTP response from a ListPlans call
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

	return response, validate.ResponseObject(response)
}
