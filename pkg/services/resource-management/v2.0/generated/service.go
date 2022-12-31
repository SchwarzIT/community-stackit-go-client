package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(
		"https://api.stackit.cloud/resource-management/v2/",
		WithHTTPClient(c),
	)
	return nc
}
