package postgresflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"postgres_flex",
	"https://postgres-flex-service.api.eu01.stackit.cloud/v1/",
	"https://postgres-flex-service.api.eu01.qa.stackit.cloud/v1/",
	"https://postgres-flex-service.api.eu01.dev.stackit.cloud/v1/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *ClientWithResponses[K] {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
