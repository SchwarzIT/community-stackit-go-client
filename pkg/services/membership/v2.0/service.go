package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

const (
	// resource types
	RESOURCE_TYPE_PROJECT = "project"
	RESOURCE_TYPE_ORG     = "organization"
)

var BaseURLs = urls.Init(
	"membership",
	"https://api.stackit.cloud/membership/",
	"https://api-qa.stackit.cloud/membership/",
	"https://api-dev.stackit.cloud/membership/",
)

func NewService(c common.Client) *ClientWithResponses {
	return NewClient(BaseURLs.GetURL(c), c)
}
