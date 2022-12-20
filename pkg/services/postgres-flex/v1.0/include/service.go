package postgresflex

import (
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
)

func NewService(c common.Client) (*postgresflex.ClientWithResponses, error) {
	surl, err := url.JoinPath(consts.DEFAULT_BASE_URL, consts.API_PATH_POSTGRES_FLEX)
	if err != nil {
		return nil, err
	}
	return postgresflex.NewClientWithResponses(surl, postgresflex.WithHTTPClient(c))
}
