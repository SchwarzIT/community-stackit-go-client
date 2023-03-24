package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"costs",
	"https://api.stackit.cloud/costs-service/v2/",
	"https://api-qa.stackit.cloud/costs-service/v2/",
	"https://api-dev.stackit.cloud/costs-service/v2/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *ClientWithResponses[K] {
	s, _ := NewClient(BaseURLs.GetURL(c.GetEnvironment()), WithHTTPClient(c))
	return s
}
