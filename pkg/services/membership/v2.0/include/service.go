package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0/generated"
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

func NewService(c common.Client) *membership.ClientWithResponses {
	return membership.NewClientWithResponses(BaseURLs.GetURL(c), c)
}
