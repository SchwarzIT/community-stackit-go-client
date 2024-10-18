package serviceaccounts

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

var BaseURLs = baseurl.New(
	"service_accounts",
	"https://service-account.api.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	return NewClient(BaseURLs.Get(), c)
}
