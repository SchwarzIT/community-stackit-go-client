package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"mongodb_flex",
	"https://mongodb-flex-service.api.eu01.stackit.cloud",
	"https://mongodb-flex-service.api.eu01.qa.stackit.cloud",
	"https://mongodb-flex-service.api.eu01.dev.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
