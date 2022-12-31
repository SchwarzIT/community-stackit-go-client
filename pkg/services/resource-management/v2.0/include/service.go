package resourcemanagement

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated"
)

func NewService(c common.Client) *resourcemanagement.ClientWithResponses {
	nc, _ := resourcemanagement.NewClientWithResponses(
		"https://api.stackit.cloud/resource-management/v2/",
		resourcemanagement.WithHTTPClient(c),
	)
	return nc
}
