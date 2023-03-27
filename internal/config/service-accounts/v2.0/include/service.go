package serviceaccounts

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0"
)

var BaseURLs = env.URLs(
	"service_accounts",
	"https://api.stackit.cloud/service-account/",
	"https://api-qa.stackit.cloud/service-account/",
	"https://api-dev.stackit.cloud/service-account/",
)

func NewService(c contracts.BaseClientInterface) *serviceaccounts.ClientWithResponses {
	return serviceaccounts.NewClient(BaseURLs.GetURL(c.GetEnvironment()), c)
}
