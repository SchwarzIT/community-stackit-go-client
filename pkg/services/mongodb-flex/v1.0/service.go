package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

var BaseURLs = baseurl.New(
	"mongodb_flex",
	"https://mongodb-flex-service.api.eu01.stackit.cloud/v1/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(BaseURLs.Get(), WithHTTPClient(c))
	return nc
}
