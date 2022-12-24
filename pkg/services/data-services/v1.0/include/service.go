package dataservices

import (
	"net/url"
	"path"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0/generated"
)

const (
	ElasticSearch = iota
	LogMe
	MariaDB
	MongoDB
	PostgresDB
	RabbitMQ
	Reddis
)

func NewService(c common.Client, serviceID int) *dataservices.ClientWithResponses {
	u, _ := url.Parse(consts.DEFAULT_BASE_URL)

	switch serviceID {
	case ElasticSearch:
		u.Path = path.Join(u.Path, "https://dsa-elasticsearch.api.eu01.stackit.cloud")
	case LogMe:
		u.Path = path.Join(u.Path, "https://dsa-logme.api.eu01.stackit.cloud")
	case MariaDB:
		u.Path = path.Join(u.Path, "https://dsa-mariadb.api.eu01.stackit.cloud")
	case MongoDB:
		u.Path = path.Join(u.Path, "https://mongo-flex-prd.api.eu01.stackit.cloud")
	case PostgresDB:
		u.Path = path.Join(u.Path, "https://postgres-flex-service.api.eu01.stackit.cloud")
	case RabbitMQ:
		u.Path = path.Join(u.Path, "https://rabbitmq.api.eu01.stackit.cloud")
	case Reddis:
		u.Path = path.Join(u.Path, "https://redis.api.eu01.stackit.cloud")
	}

	nc, _ := dataservices.NewClientWithResponses(u.String(), dataservices.WithHTTPClient(c))
	return nc
}
