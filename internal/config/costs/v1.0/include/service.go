package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v1.0"
)

var BaseURLs = env.URLs(
	"costs",
	"https://api.stackit.cloud/costs-service/v1/",
	"https://api-qa.stackit.cloud/costs-service/v1/",
	"https://api-dev.stackit.cloud/costs-service/v1/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *costs.ClientWithResponses[K] {
	s, _ := costs.NewClient(BaseURLs.GetURL(c.GetEnvironment()), costs.WithHTTPClient(c))
	return s
}
