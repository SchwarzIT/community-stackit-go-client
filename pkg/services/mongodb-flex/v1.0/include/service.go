package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/generated"
)

func NewService(c common.Client) *mongodbflex.ClientWithResponses {
	nc, _ := mongodbflex.NewClientWithResponses(
		"https://api.stackit.cloud/mongodb",
		mongodbflex.WithHTTPClient(c),
	)
	return nc
}
