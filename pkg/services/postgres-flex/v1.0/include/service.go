package postgresflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"postgres_flex",
	"https://postgres-flex-service.api.eu01.stackit.cloud",
	"https://postgres-flex-service.api.eu01.qa.stackit.cloud",
	"https://postgres-flex-service.api.eu01.dev.stackit.cloud",
)

func NewService(c common.Client) *postgresflex.ClientWithResponses {
	nc, _ := postgresflex.NewClientWithResponses(
		BaseURLs.GetURL(c),
		postgresflex.WithHTTPClient(c),
	)
	return nc
}
