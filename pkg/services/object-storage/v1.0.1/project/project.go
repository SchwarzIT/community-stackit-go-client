// Package project provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.5 DO NOT EDIT.
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

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/internal/helpers/runtime"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

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

// NewRawClient Creates a new Client, with reasonable defaults
func NewRawClient(server string, httpClient common.Client) *Client {
	// create a client with sane default values
	client := Client{
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

func (c *Client) DeleteRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) GetRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

func (c *Client) CreateRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
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

	operationPath := fmt.Sprintf("/v1/project/%s", pathParam0)
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

	operationPath := fmt.Sprintf("/v1/project/%s", pathParam0)
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

	operationPath := fmt.Sprintf("/v1/project/%s", pathParam0)
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

// ClientWithResponses builds on rawClientInterface to offer response payloads
type ClientWithResponses struct {
	rawClientInterface
}

// NewClient creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClient(server string, httpClient common.Client) *ClientWithResponses {
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
	JSON200      *struct {
		// Project Project ID
		Project string `json:"project"`

		// Scope Project Scope
		Scope interface{} `json:"scope"`
	}
	JSON400 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON403 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON404 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON422 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON500 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	Error error // Aggregated error
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
	JSON200      *struct {
		// Project Project ID
		Project string `json:"project"`

		// Scope Project Scope
		Scope interface{} `json:"scope"`
	}
	JSON403 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON404 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON422 *struct {
		Details *[]struct {
			Loc  []string `json:"loc"`
			Msg  string   `json:"msg"`
			Type string   `json:"type"`
		} `json:"detail,omitempty"`
	}
	JSON500 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	Error error // Aggregated error
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
	JSON200      *struct {
		// Project Project ID
		Project string `json:"project"`

		// Scope Project Scope
		Scope interface{} `json:"scope"`
	}
	JSON201 *struct {
		// Project Project ID
		Project string `json:"project"`

		// Scope Project Scope
		Scope interface{} `json:"scope"`
	}
	JSON403 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON409 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	JSON422 *struct {
		Details *[]struct {
			Loc  []string `json:"loc"`
			Msg  string   `json:"msg"`
			Type string   `json:"type"`
		} `json:"detail,omitempty"`
	}
	JSON500 *struct {
		Details []struct {
			Key string `json:"key"`
			Msg string `json:"msg"`
		} `json:"detail"`
	}
	Error error // Aggregated error
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
		var dest struct {
			// Project Project ID
			Project string `json:"project"`

			// Scope Project Scope
			Scope interface{} `json:"scope"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON500 = &dest

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
		var dest struct {
			// Project Project ID
			Project string `json:"project"`

			// Scope Project Scope
			Scope interface{} `json:"scope"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest struct {
			Details *[]struct {
				Loc  []string `json:"loc"`
				Msg  string   `json:"msg"`
				Type string   `json:"type"`
			} `json:"detail,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON500 = &dest

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
		var dest struct {
			// Project Project ID
			Project string `json:"project"`

			// Scope Project Scope
			Scope interface{} `json:"scope"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Project Project ID
			Project string `json:"project"`

			// Scope Project Scope
			Scope interface{} `json:"scope"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest struct {
			Details *[]struct {
				Loc  []string `json:"loc"`
				Msg  string   `json:"msg"`
				Type string   `json:"type"`
			} `json:"detail,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Details []struct {
				Key string `json:"key"`
				Msg string `json:"msg"`
			} `json:"detail"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON500 = &dest

	}

	return response, validate.ResponseObject(response)
}
