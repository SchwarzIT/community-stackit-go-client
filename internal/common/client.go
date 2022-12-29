// client file in package common holds the client interface and service struct
// used by each service that the client is connecting with
// services using the Service struct are located under pkg/api
package common

import (
	"context"
	"net/http"
)

// Client is the client interface
type Client interface {
	Request(ctx context.Context, method, path string, body []byte) (*http.Request, error)
	LegacyDo(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
	SetBaseURL(url string) error
	GetBaseURL() string
}

// Service is the struct every extending service is built on
type Service struct {
	Client Client
}
