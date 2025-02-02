package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

var BaseURLs = baseurl.New(
	"costs",
	"https://metering.api.eu01.stackit.cloud/v2/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	s, _ := NewClient(BaseURLs.Get(), WithHTTPClient(c))
	return s
}
