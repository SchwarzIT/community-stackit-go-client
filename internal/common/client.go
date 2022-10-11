// client file in package common holds the client interface and service struct
// used by each service that the client is connecting with
// services using the Service struct are located under pkg/api
package common

import (
	"context"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/retry"
)

// Client is the client interface
type Client interface {
	Request(ctx context.Context, method, path string, body []byte) (*http.Request, error)
	Do(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error)
	Retry() *retry.Retry
	OrganizationID() string
}

// Service is the struct every extending service is built on
type Service struct {
	Client Client
}
