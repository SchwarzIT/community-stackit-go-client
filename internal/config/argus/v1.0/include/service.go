package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0"
)

var BaseURLs = env.URLs(
	"argus",
	"https://argus.api.eu01.stackit.cloud",
	"https://argus.api.eu01.qa.stackit.cloud",
	"https://argus.api.eu01.dev.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *argus.ClientWithResponses {
	nc, _ := argus.NewClient(BaseURLs.GetURL(c.GetEnvironment()), argus.WithHTTPClient(c))
	return nc
}
