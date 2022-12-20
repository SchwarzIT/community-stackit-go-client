package kubernetes

import (
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.4/generated"
)

func NewService(c common.Client) (*kubernetes.ClientWithResponses, error) {
	surl, err := url.JoinPath(consts.DEFAULT_BASE_URL, consts.API_PATH_SKE)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewClientWithResponses(surl, kubernetes.WithHTTPClient(c))
}
