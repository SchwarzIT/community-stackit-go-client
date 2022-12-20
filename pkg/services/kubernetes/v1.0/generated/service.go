package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(consts.BASE_URL_SKE, WithHTTPClient(c))
	return nc
}
