package elasticsearch

import (
	"net/url"
	"path"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	elasticsearch "github.com/SchwarzIT/community-stackit-go-client/pkg/services/elasticsearch/v1.0/generated"
)

func NewService(c common.Client) *elasticsearch.ClientWithResponses {
	u, _ := url.Parse(consts.DEFAULT_BASE_URL)
	u.Path = path.Join(u.Path, consts.API_PATH_POSTGRES_FLEX)
	nc, _ := elasticsearch.NewClientWithResponses(u.String(), elasticsearch.WithHTTPClient(c))
	return nc
}
