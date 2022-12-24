package rabbitmq

import (
	"net/url"
	"path"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	rabbitmq "github.com/SchwarzIT/community-stackit-go-client/pkg/services/rabbitmq/v1.0/generated"
)

func NewService(c common.Client) *rabbitmq.ClientWithResponses {
	u, _ := url.Parse(consts.DEFAULT_BASE_URL)
	u.Path = path.Join(u.Path, consts.API_PATH_POSTGRES_FLEX)
	nc, _ := rabbitmq.NewClientWithResponses(u.String(), rabbitmq.WithHTTPClient(c))
	return nc
}
