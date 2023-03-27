// client file in package common holds the client interface and service struct
// used by each service that the client is connecting with
// services using the Service struct are located under pkg/api
package contracts

import (
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/clients"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

type ClientFlowConfig interface {
	clients.TokenFlowConfig | clients.KeyFlowConfig
}

type ClientInterface[f ClientFlowConfig] interface {
	Do(req *http.Request) (*http.Response, error)
	GetConfig() f
	GetEnvironment() env.Environment
	GetServiceAccountEmail() string
}

type BaseClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
	GetEnvironment() env.Environment
	GetServiceAccountEmail() string
}
