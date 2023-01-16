package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	return NewClientWithResponses("https://api.stackit.cloud/costs-service/v1/", c)
}
