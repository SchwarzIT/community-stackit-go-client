package postgresflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
)

func NewService(c common.Client) *postgresflex.ClientWithResponses {
	nc, _ := postgresflex.NewClientWithResponses(
		"https://postgres-flex-service.api.eu01.stackit.cloud",
		postgresflex.WithHTTPClient(c),
	)
	return nc
}
