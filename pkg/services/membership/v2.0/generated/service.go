package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

const (
	// resource types
	RESOURCE_TYPE_PROJECT = "project"
	RESOURCE_TYPE_ORG     = "organization"
)

func NewService(c common.Client) *ClientWithResponses {
	return NewClientWithResponses(getURL(c), c)
}

func getURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://api-dev.stackit.cloud/membership/"
	case common.ENV_QA:
		return "https://api-qa.stackit.cloud/membership/"
	default:
		return "https://api.stackit.cloud/membership/"
	}
}
