package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
)

func NewService(c common.Client) *kubernetes.ClientWithResponses {
	nc, _ := kubernetes.NewClientWithResponses(
		"https://ske.api.eu01.stackit.cloud/",
		kubernetes.WithHTTPClient(c),
	)
	return nc
}
