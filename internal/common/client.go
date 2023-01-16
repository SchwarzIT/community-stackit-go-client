// client file in package common holds the client interface and service struct
// used by each service that the client is connecting with
// services using the Service struct are located under pkg/api
package common

import (
	"context"
	"net/http"
)

const (
	DEFAULT_BASE_URL = "https://api.stackit.cloud/"
)

// Client is the client interface
type Client interface {
	Request(ctx context.Context, method, path string, body []byte) (*http.Request, error)
	Do(req *http.Request) (*http.Response, error)
	SetBaseURL(url string) error
	GetBaseURL() string
}

// Service is the struct every extending service is built on
type Service struct {
	Client Client
}
