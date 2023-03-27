package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0"
)

var BaseURLs = env.URLs(
	"argus",
	"https://argus.api.stackit.cloud",
	"https://argus.api.stg.stackit.cloud",
	"https://argus.api.dev.stackit.cloud",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *argus.ClientWithResponses[K] {
	nc, _ := argus.NewClient(BaseURLs.GetURL(c.GetEnvironment()), argus.WithHTTPClient(c))
	return nc
}
