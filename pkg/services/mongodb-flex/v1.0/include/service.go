package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/generated"
)

func NewService(c common.Client) *mongodbflex.ClientWithResponses {
	nc, _ := mongodbflex.NewClientWithResponses(
		getURL(c),
		mongodbflex.WithHTTPClient(c),
	)
	return nc
}

func getURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://mongo-flex-dev.api.eu01.stackit.cloud/"
	case common.ENV_QA:
		return "https://mongo-flex-qa.api.eu01.stackit.cloud/"
	default:
		return "https://mongo-flex-prd.api.eu01.stackit.cloud/"
	}
}
