package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated"
)

func NewService(c common.Client) *resourcemanagement.ClientWithResponses {
	nc, _ := resourcemanagement.NewClientWithResponses(
		getURL(c),
		resourcemanagement.WithHTTPClient(c),
	)
	return nc
}

func getURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://api-dev.stackit.cloud/resource-management/v2/"
	case common.ENV_QA:
		return "https://api-qa.stackit.cloud/resource-management/v2/"
	default:
		return "https://api.stackit.cloud/resource-management/v2/"
	}
}
