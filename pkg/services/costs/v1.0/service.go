package costs

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"costs",
	"https://api.stackit.cloud/costs-service/v1/",
	"https://api-qa.stackit.cloud/costs-service/v1/",
	"https://api-dev.stackit.cloud/costs-service/v1/",
)

func NewService(c common.Client) *ClientWithResponses {
	s, _ := NewClient(BaseURLs.GetURL(c), WithHTTPClient(c))
	return s
}
