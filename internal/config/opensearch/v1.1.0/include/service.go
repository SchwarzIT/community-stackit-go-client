package opensearch

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	opensearch "github.com/SchwarzIT/community-stackit-go-client/pkg/services/opensearch/v1.1.0"
)

var BaseURLs = env.URLs(
	"opensearch",
	"https://opensearch.api.eu01.stackit.cloud",
	"https://opensearch.api.eu01.qa.stackit.cloud",
	"https://opensearch.api.eu01.dev.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *opensearch.ClientWithResponses {
	url := BaseURLs.GetURL(c.GetEnvironment())
	nc, _ := opensearch.NewClient(url, opensearch.WithHTTPClient(c))
	return nc
}
