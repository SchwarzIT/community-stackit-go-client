package serviceaccounts

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0"
)

var BaseURLs = baseurl.New(
	"service_accounts",
	"https://service-account.api.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) *serviceaccounts.ClientWithResponses {
	return serviceaccounts.NewClient(BaseURLs.Get(), c)
}
