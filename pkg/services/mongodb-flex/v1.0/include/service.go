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
		return "https://api-dev.stackit.cloud/mongodb/v1/"
	case common.ENV_QA:
		return "https://api-qa.stackit.cloud/mongodb/v1/"
	default:
		return "https://api.stackit.cloud/mongodb/v1/"
	}
}
