package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0"
)

const (
	// resource types
	RESOURCE_TYPE_PROJECT = "project"
	RESOURCE_TYPE_ORG     = "organization"
)

var BaseURLs = env.URLs(
	"membership",
	"https://api.stackit.cloud/membership/",
	"https://api-qa.stackit.cloud/membership/",
	"https://api-dev.stackit.cloud/membership/",
)

func NewService(c contracts.BaseClientInterface) *membership.ClientWithResponses {
	return membership.NewClient(BaseURLs.GetURL(c.GetEnvironment()), c)
}
