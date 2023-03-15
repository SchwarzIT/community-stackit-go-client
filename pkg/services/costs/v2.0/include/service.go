package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v2.0/generated"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"costs",
	"https://api.stackit.cloud/costs-service/v2/",
	"https://api-qa.stackit.cloud/costs-service/v2/",
	"https://api-dev.stackit.cloud/costs-service/v2/",
)

func NewService(c common.Client) *costs.ClientWithResponses {
	s, _ := costs.NewClientWithResponses(BaseURLs.GetURL(c), costs.WithHTTPClient(c))
	return s
}
