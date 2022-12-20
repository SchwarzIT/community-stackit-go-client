package postgresflex

import (
	"net/url"
	"path"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
)

func NewService(c common.Client) (*postgresflex.ClientWithResponses, error) {
	u, err := url.Parse(consts.DEFAULT_BASE_URL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, consts.API_PATH_POSTGRES_FLEX)
	return postgresflex.NewClientWithResponses(u.String(), postgresflex.WithHTTPClient(c))
}
