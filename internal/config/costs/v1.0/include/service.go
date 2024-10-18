package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v1.0"
)

var BaseURLs = baseurl.New(
	"costs",
	"https://metering.api.eu01.stackit.cloud/v1/",
)

func NewService(c contracts.BaseClientInterface) *costs.ClientWithResponses {
	s, _ := costs.NewClient(BaseURLs.Get(), costs.WithHTTPClient(c))
	return s
}
