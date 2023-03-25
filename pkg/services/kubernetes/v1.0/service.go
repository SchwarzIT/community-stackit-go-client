package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"kubernetes",
	"https://ske.api.eu01.stackit.cloud/",
	"https://ske.api.eu01.stg.stackit.cloud/",
	"https://ske.api.eu01.dev.stackit.cloud/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *ClientWithResponses[K] {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
