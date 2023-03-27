package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	mongodb "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0"
)

var BaseURLs = env.URLs(
	"mongodb_flex",
	"https://api.stackit.cloud/mongodb/v1/",
	"https://api-qa.stackit.cloud/mongodb/v1/",
	"https://api-dev.stackit.cloud/mongodb/v1/",
)

func NewService(c contracts.BaseClientInterface) *mongodb.ClientWithResponses {
	nc, _ := mongodb.NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		mongodb.WithHTTPClient(c),
	)
	return nc
}
