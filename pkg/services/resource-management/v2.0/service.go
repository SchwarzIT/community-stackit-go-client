package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"resource_management",
	"https://api.stackit.cloud/resource-management/v2/",
	"https://api-qa.stackit.cloud/resource-management/v2/",
	"https://api-dev.stackit.cloud/resource-management/v2/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *ClientWithResponses[K] {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
