package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
)

var BaseURLs = env.URLs(
	"resource_management",
	"https://api.stackit.cloud/resource-management/v2/",
	"https://api-qa.stackit.cloud/resource-management/v2/",
	"https://api-dev.stackit.cloud/resource-management/v2/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *resourcemanagement.ClientWithResponses[K] {
	nc, _ := resourcemanagement.NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		resourcemanagement.WithHTTPClient(c),
	)
	return nc
}
