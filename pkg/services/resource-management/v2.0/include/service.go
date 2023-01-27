package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"resource_management",
	"https://api.stackit.cloud/resource-management/v2/",
	"https://api-qa.stackit.cloud/resource-management/v2/",
	"https://api-dev.stackit.cloud/resource-management/v2/",
)

func NewService(c common.Client) *resourcemanagement.ClientWithResponses {
	nc, _ := resourcemanagement.NewClientWithResponses(
		BaseURLs.GetURL(c),
		resourcemanagement.WithHTTPClient(c),
	)
	return nc
}
