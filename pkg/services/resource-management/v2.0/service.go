package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"resource_management",
	"https://resource-manager.api.stackit.cloud/v2/",
	"https://resource-manager.api.qa.stackit.cloud/v2/",
	"https://resource-manager.api.dev.stackit.cloud/v2/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
