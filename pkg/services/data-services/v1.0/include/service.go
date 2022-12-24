package dataservices

import (
	"net/url"
	"path"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0/generated"
)

func NewService(c common.Client) *dataservices.ClientWithResponses {
	u, _ := url.Parse(consts.DEFAULT_BASE_URL)
	u.Path = path.Join(u.Path, consts.API_PATH_POSTGRES_FLEX)
	nc, _ := dataservices.NewClientWithResponses(u.String(), dataservices.WithHTTPClient(c))
	return nc
}
