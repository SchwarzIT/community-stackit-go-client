package serviceaccounts

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"service_accounts",
	"https://api.stackit.cloud/service-account/",
	"https://api-qa.stackit.cloud/service-account/",
	"https://api-dev.stackit.cloud/service-account/",
)

func NewService(c common.Client) *ClientWithResponses {
	return NewClientWithResponses(BaseURLs.GetURL(c), c)
}
