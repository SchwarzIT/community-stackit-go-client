package postgresflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"postgres_flex",
	"https://postgres-flex-service.api.eu01.stackit.cloud",
	"https://postgres-flex-service.api.eu01.qa.stackit.cloud",
	"https://postgres-flex-service.api.eu01.dev.stackit.cloud",
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.GetURL(c),
		WithHTTPClient(c),
	)
	return nc
}
