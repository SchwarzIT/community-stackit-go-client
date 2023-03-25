// Package bucket provides primitives to interact with the openapi HTTP API.
//
// Code generated by dev.azure.com/schwarzit/schwarzit.odj.core/_git/stackit-client-generator.git version v1.0.23 DO NOT EDIT.
package bucket

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
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

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
	// Delete request
	DeleteRaw(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Get request
	GetRaw(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Create request
	CreateRaw(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// List request
	ListRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client[K]) DeleteRaw(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteRequest(ctx, c.Server, projectID, bucketName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) GetRaw(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequest(ctx, c.Server, projectID, bucketName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) CreateRaw(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequest(ctx, c.Server, projectID, bucketName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client[K]) ListRaw(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListRequest(ctx, c.Server, projectID)
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
func NewDeleteRequest(ctx context.Context, server string, projectID string, bucketName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "bucketName", runtime.ParamLocationPath, bucketName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/project/%s/bucket/%s", pathParam0, pathParam1)
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
func NewGetRequest(ctx context.Context, server string, projectID string, bucketName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "bucketName", runtime.ParamLocationPath, bucketName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/project/%s/bucket/%s", pathParam0, pathParam1)
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
func NewCreateRequest(ctx context.Context, server string, projectID string, bucketName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "projectID", runtime.ParamLocationPath, projectID)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "bucketName", runtime.ParamLocationPath, bucketName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/project/%s/bucket/%s", pathParam0, pathParam1)
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

// NewListRequest generates requests for List
func NewListRequest(ctx context.Context, server string, projectID string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/v1/project/%s/buckets", pathParam0)
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
	// Delete request
	Delete(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*DeleteResponse, error)

	// Get request
	Get(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*GetResponse, error)

	// Create request
	Create(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	// List request
	List(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListResponse, error)
}

type DeleteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Bucket Name of the bucket
		Bucket string `json:"bucket"`

		// Project Project ID
		Project string `json:"project"`
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
		Bucket struct {
			Name   string `json:"name"`
			Region string `json:"region"`

			// UrlPathStyle URL in path style
			UrlPathStyle string `json:"urlPathStyle"`

			// UrlVirtualHostedStyle URL in virtual hosted style
			UrlVirtualHostedStyle string `json:"urlVirtualHostedStyle"`
		} `json:"bucket"`

		// Project Project ID
		Project string `json:"project"`
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
	JSON201      *struct {
		// Bucket Name of the bucket
		Bucket string `json:"bucket"`

		// Project Project ID
		Project string `json:"project"`
	}
	JSON404 *struct {
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

type ListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Buckets []struct {
			Name   string `json:"name"`
			Region string `json:"region"`

			// UrlPathStyle URL in path style
			UrlPathStyle string `json:"urlPathStyle"`

			// UrlVirtualHostedStyle URL in virtual hosted style
			UrlVirtualHostedStyle string `json:"urlVirtualHostedStyle"`
		} `json:"buckets"`

		// Project Project ID
		Project string `json:"project"`
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

// Delete request returning *DeleteResponse
func (c *ClientWithResponses[K]) Delete(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*DeleteResponse, error) {
	rsp, err := c.DeleteRaw(ctx, projectID, bucketName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseDeleteResponse(rsp)
}

// Get request returning *GetResponse
func (c *ClientWithResponses[K]) Get(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.GetRaw(ctx, projectID, bucketName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseGetResponse(rsp)
}

// Create request returning *CreateResponse
func (c *ClientWithResponses[K]) Create(ctx context.Context, projectID string, bucketName string, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.CreateRaw(ctx, projectID, bucketName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseCreateResponse(rsp)
}

// List request returning *ListResponse
func (c *ClientWithResponses[K]) List(ctx context.Context, projectID string, reqEditors ...RequestEditorFn) (*ListResponse, error) {
	rsp, err := c.ListRaw(ctx, projectID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return c.ParseListResponse(rsp)
}

// ParseDeleteResponse parses an HTTP response from a Delete call
func (c *ClientWithResponses[K]) ParseDeleteResponse(rsp *http.Response) (*DeleteResponse, error) {
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
			// Bucket Name of the bucket
			Bucket string `json:"bucket"`

			// Project Project ID
			Project string `json:"project"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON200 = &dest

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
func (c *ClientWithResponses[K]) ParseGetResponse(rsp *http.Response) (*GetResponse, error) {
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
			Bucket struct {
				Name   string `json:"name"`
				Region string `json:"region"`

				// UrlPathStyle URL in path style
				UrlPathStyle string `json:"urlPathStyle"`

				// UrlVirtualHostedStyle URL in virtual hosted style
				UrlVirtualHostedStyle string `json:"urlVirtualHostedStyle"`
			} `json:"bucket"`

			// Project Project ID
			Project string `json:"project"`
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
func (c *ClientWithResponses[K]) ParseCreateResponse(rsp *http.Response) (*CreateResponse, error) {
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
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Bucket Name of the bucket
			Bucket string `json:"bucket"`

			// Project Project ID
			Project string `json:"project"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("body was: %s", string(bodyBytes)))
		}
		response.JSON201 = &dest

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
		var dest struct {
			Buckets []struct {
				Name   string `json:"name"`
				Region string `json:"region"`

				// UrlPathStyle URL in path style
				UrlPathStyle string `json:"urlPathStyle"`

				// UrlVirtualHostedStyle URL in virtual hosted style
				UrlVirtualHostedStyle string `json:"urlVirtualHostedStyle"`
			} `json:"buckets"`

			// Project Project ID
			Project string `json:"project"`
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
