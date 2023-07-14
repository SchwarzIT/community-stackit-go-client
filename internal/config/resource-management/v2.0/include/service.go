package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
)

var BaseURLs = env.URLs(
	"resource_management",
	"https://resource-manager.api.eu01.stackit.cloud/v2/",
	"https://resource-manager.api.eu01.qa.stackit.cloud/v2/",
	"https://resource-manager.api.eu01.dev.stackit.cloud/v2/",
)

func NewService(c contracts.BaseClientInterface) *resourcemanagement.ClientWithResponses {
	nc, _ := resourcemanagement.NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		resourcemanagement.WithHTTPClient(c),
	)
	return nc
}
