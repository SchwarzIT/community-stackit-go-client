package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"mongodb_flex",
	"https://api.stackit.cloud/mongodb/v1/",
	"https://api-qa.stackit.cloud/mongodb/v1/",
	"https://api-dev.stackit.cloud/mongodb/v1/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
