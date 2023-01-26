package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
)

func NewService(c common.Client) *kubernetes.ClientWithResponses {
	nc, _ := kubernetes.NewClientWithResponses(
		getURL(c),
		kubernetes.WithHTTPClient(c),
	)
	return nc
}

func getURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://ske.api.eu01.dev.stackit.cloud/"
	case common.ENV_QA:
		return "https://ske.api.eu01.stg.stackit.cloud/"
	default:
		return "https://ske.api.eu01.stackit.cloud/"
	}
}
