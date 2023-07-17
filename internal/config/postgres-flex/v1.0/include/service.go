package postgresflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0"
)

var BaseURLs = baseurl.New(
	"postgres_flex",
	"https://postgres-flex-service.api.eu01.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *postgresflex.ClientWithResponses {
	nc, _ := postgresflex.NewClient(
		BaseURLs.Get(),
		postgresflex.WithHTTPClient(c),
	)
	return nc
}
