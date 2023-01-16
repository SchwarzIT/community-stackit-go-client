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
	return NewClientWithResponses(
		"https://api.stackit.cloud/membership/",
		c,
	)
}
