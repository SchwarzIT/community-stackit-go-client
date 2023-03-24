package postgresflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"postgres_flex",
	"https://postgres-flex-service.api.eu01.stackit.cloud",
	"https://postgres-flex-service.api.eu01.qa.stackit.cloud",
	"https://postgres-flex-service.api.eu01.dev.stackit.cloud",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *ClientWithResponses[K] {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
