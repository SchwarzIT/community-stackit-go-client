package costs

import (
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client, baseURL *url.URL) *ClientWithResponses {
	url := getURLFromEnvironment(c)
	if baseURL != nil {
		url = baseURL.String()
	}
	client := NewClientWithResponses(url, c)
	return client
}

func getURLFromEnvironment(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://api-dev.stackit.cloud/costs-service/v1/"
	case common.ENV_QA:
		return "https://api-qa.stackit.cloud/costs-service/v1/"
	default:
		return "https://api.stackit.cloud/costs-service/v1/"
	}
}
