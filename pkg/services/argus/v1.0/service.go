package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"argus",
	"https://argus.api.stackit.cloud",
	"https://argus.api.stg.stackit.cloud",
	"https://argus.api.dev.stackit.cloud",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *ClientWithResponses[K] {
	nc, _ := NewClient(BaseURLs.GetURL(c.GetEnvironment()), WithHTTPClient(c))
	return nc
}
