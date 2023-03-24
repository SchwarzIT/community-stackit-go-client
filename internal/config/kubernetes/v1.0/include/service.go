package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0"
)

var BaseURLs = env.URLs(
	"kubernetes",
	"https://ske.api.eu01.stackit.cloud/",
	"https://ske.api.eu01.stg.stackit.cloud/",
	"https://ske.api.eu01.dev.stackit.cloud/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *kubernetes.ClientWithResponses[K] {
	nc, _ := kubernetes.NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		kubernetes.WithHTTPClient(c),
	)
	return nc
}
