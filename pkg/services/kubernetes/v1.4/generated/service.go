package kubernetes

import (
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

func NewService(c common.Client) (*ClientWithResponses, error) {
	surl, err := url.JoinPath(consts.DEFAULT_BASE_URL, consts.API_PATH_SKE)
	if err != nil {
		return nil, err
	}
	return NewClientWithResponses(surl, WithHTTPClient(c))
}
