package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

const (
	// resource types
	RESOURCE_TYPE_PROJECT = "project"
	RESOURCE_TYPE_ORG     = "organization"
)

var BaseURLs = baseurl.New(
	"membership",
	"https://authorization.api.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	return NewClient(BaseURLs.Get(), c)
}
