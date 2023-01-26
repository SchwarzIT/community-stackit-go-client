package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/internal/urls"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v1.0/generated"
)

var BaseURLs = urls.Init(
	"costs",
	"https://api.stackit.cloud/costs-service/v1/",
	"https://api-qa.stackit.cloud/costs-service/v1/",
	"https://api-dev.stackit.cloud/costs-service/v1/",
)

func NewService(c common.Client) *costs.ClientWithResponses {
	s, _ := costs.NewClientWithResponses(BaseURLs.GetURL(c), costs.WithHTTPClient(c))
	return s
}
