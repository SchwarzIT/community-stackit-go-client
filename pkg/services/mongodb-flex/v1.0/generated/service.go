package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(
		"https://mongodb-flex-service.api.eu01.stackit.cloud",
		WithHTTPClient(c),
	)
	return nc
}
