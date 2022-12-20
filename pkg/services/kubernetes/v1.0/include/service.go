package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
)

func NewService(c common.Client) (*kubernetes.ClientWithResponses, error) {
	return kubernetes.NewClientWithResponses(consts.BASE_URL_SKE, kubernetes.WithHTTPClient(c))
}
