package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated"
)

func NewService(c common.Client) *membership.ClientWithResponses {
	nc, _ := membership.NewClientWithResponses(
		"https://api.stackit.cloud/membership/",
		membership.WithHTTPClient(c),
	)
	return nc
}
