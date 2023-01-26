// client file in package common holds the client interface and service struct
// used by each service that the client is connecting with
// services using the Service struct are located under pkg/api
package common

import (
	"net/http"
)

type Environment string

const (
	ENV_PROD Environment = "prod"
	ENV_QA   Environment = "qa"
	ENV_DEV  Environment = "dev"
)

const (
	DEFAULT_BASE_URL = "https://api.stackit.cloud/"
)

// Client is the client interface
type Client interface {
	Do(req *http.Request) (*http.Response, error)
	GetEnvironment() Environment
}

// Service is the struct every extending service is built on
type Service struct {
	Client Client
}
