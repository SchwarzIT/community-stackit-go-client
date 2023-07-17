package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	mongodb "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0"
)

var BaseURLs = baseurl.New(
	"mongodb_flex",
	"https://mongodb-flex-service.api.eu01.stackit.cloud/v1/",
)

func NewService(c contracts.BaseClientInterface) *mongodb.ClientWithResponses {
	nc, _ := mongodb.NewClient(BaseURLs.Get(), mongodb.WithHTTPClient(c))
	return nc
}
