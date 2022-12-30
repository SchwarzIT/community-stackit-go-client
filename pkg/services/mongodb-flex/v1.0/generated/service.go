package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(
		"https://api.stackit.cloud/mongodb/v1",
		WithHTTPClient(c),
	)
	return nc
}
