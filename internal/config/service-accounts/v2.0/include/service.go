package serviceaccounts

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"service_accounts",
	"https://api.stackit.cloud/service-account/",
	"https://api-qa.stackit.cloud/service-account/",
	"https://api-dev.stackit.cloud/service-account/",
)

func NewService(c common.Client) *serviceaccounts.ClientWithResponses {
	return serviceaccounts.NewClient(BaseURLs.GetURL(c), c)
}
