package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(getURL(c), WithHTTPClient(c))
	return nc
}

func getURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://argus.api.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://argus.api.stg.stackit.cloud"
	default:
		return "https://argus.api.stackit.cloud"
	}
}
