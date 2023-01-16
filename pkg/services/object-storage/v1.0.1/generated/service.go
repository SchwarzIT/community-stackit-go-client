package objectstorage

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
		return "https://api-dev.stackit.cloud/object-storage-api/"
	case common.ENV_QA:
		return "https://api-qa.stackit.cloud/object-storage-api/"
	default:
		return "https://api.stackit.cloud/object-storage-api/"
	}
}
