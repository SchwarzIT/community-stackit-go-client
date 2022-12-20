package postgresflex

import (
	"net/url"
	"path"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

func NewService(c common.Client) (*ClientWithResponses, error) {
	u, err := url.Parse(consts.DEFAULT_BASE_URL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, consts.API_PATH_POSTGRES_FLEX)
	return NewClientWithResponses(u.String(), WithHTTPClient(c))
}
