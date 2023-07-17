package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
)

var BaseURLs = baseurl.New(
	"resource_management",
	"https://resource-manager.api.stackit.cloud/v2/",
)

func NewService(c contracts.BaseClientInterface) *resourcemanagement.ClientWithResponses {
	nc, _ := resourcemanagement.NewClient(
		BaseURLs.Get(),
		resourcemanagement.WithHTTPClient(c),
	)
	return nc
}
