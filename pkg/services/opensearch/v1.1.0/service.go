package opensearch

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"elasticsearch",
	"https://elasticsearch.api.eu01.stackit.cloud",
	"https://elasticsearch.api.eu01.qa.stackit.cloud",
	"https://elasticsearch.api.eu01.dev.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	url := BaseURLs.GetURL(c.GetEnvironment())
	nc, _ := NewClient(url, WithHTTPClient(c))
	return nc
}
