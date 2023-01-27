package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"resource_management",
	"https://api.stackit.cloud/resource-management/v2/",
	"https://api-qa.stackit.cloud/resource-management/v2/",
	"https://api-dev.stackit.cloud/resource-management/v2/",
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(
		BaseURLs.GetURL(c),
		WithHTTPClient(c),
	)
	return nc
}
