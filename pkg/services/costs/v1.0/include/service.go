package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v1.0/generated"
)

func NewService(c common.Client) *costs.ClientWithResponses {
	return costs.NewClientWithResponses("https://api.stackit.cloud/costs-service/v1/", c)
}
