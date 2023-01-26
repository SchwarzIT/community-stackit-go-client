package costs

import (
	"os"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v1.0/generated"
)

func NewService(c common.Client) *costs.ClientWithResponses {
	s, _ := costs.NewClientWithResponses(getURL(c), costs.WithHTTPClient(c))
	return s
}

func getURL(c common.Client) string {
	url := os.Getenv("STACKIT_COST_BASEURL")
	if url != "" {
		return url
	}

	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://api-dev.stackit.cloud/costs-service/v1/"
	case common.ENV_QA:
		return "https://api-qa.stackit.cloud/costs-service/v1/"
	default:
		return "https://api.stackit.cloud/costs-service/v1/"
	}
}
