package postgresflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(
		getURL(c),
		WithHTTPClient(c),
	)
	return nc
}

func getURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://postgres-flex-service.api.eu01.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://postgres-flex-service.api.eu01.qa.stackit.cloud"
	default:
		return "https://postgres-flex-service.api.eu01.stackit.cloud"
	}
}
