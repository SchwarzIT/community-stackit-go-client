package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

var BaseURLs = baseurl.New(
	"resource_management",
	"https://resource-manager.api.stackit.cloud/v2/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.Get(),
		WithHTTPClient(c),
	)
	return nc
}
