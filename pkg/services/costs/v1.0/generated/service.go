package costs

import (
	"os"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	s, _ := NewClientWithResponses(getURL(c), WithHTTPClient(c))
	return s
}

func getURL(c common.Client) string {
	url := os.Getenv("STACKIT_COSTS_BASEURL")
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
