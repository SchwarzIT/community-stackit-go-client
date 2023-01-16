package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0/generated"
)

const (
	// resource types
	RESOURCE_TYPE_PROJECT = "project"
	RESOURCE_TYPE_ORG     = "organization"
)

func NewService(c common.Client) *membership.ClientWithResponses {
	return membership.NewClientWithResponses(
		"https://api.stackit.cloud/membership/",
		c,
	)
}
