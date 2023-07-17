package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0"
)

const (
	// resource types
	RESOURCE_TYPE_PROJECT = "project"
	RESOURCE_TYPE_ORG     = "organization"
)

var BaseURLs = baseurl.New(
	"membership",
	"https://api.stackit.cloud/membership/",
)

func NewService(c contracts.BaseClientInterface) *membership.ClientWithResponses {
	return membership.NewClient(BaseURLs.Get(), c)
}
