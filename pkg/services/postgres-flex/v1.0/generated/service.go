package postgresflex

import (
	"net/url"
	"path"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

func NewService(c common.Client) *ClientWithResponses {
	u, _ := url.Parse(consts.DEFAULT_BASE_URL)
	u.Path = path.Join(u.Path, consts.API_PATH_POSTGRES_FLEX)
	nc, _ := NewClientWithResponses(u.String(), WithHTTPClient(c))
	return nc
}
