package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	mongodb "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0"
)

var BaseURLs = env.URLs(
	"mongodb_flex",
	"https://mongodb-flex-service.api.eu01.stackit.cloud",
	"https://mongodb-flex-service.api.eu01.qa.stackit.cloud",
	"https://mongodb-flex-service.api.eu01.dev.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *mongodb.ClientWithResponses {
	nc, _ := mongodb.NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		mongodb.WithHTTPClient(c),
	)
	return nc
}
