package dataservices

import (
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
	Redis
)

func NewService(c common.Client, serviceID int) *dataservices.ClientWithResponses {
	url := consts.DEFAULT_BASE_URL

	switch serviceID {
	case ElasticSearch:
		url = "https://dsa-elasticsearch.api.eu01.stackit.cloud"
	case LogMe:
		url = "https://dsa-logme.api.eu01.stackit.cloud"
	case MariaDB:
		url = "https://dsa-mariadb.api.eu01.stackit.cloud"
	case MongoDB:
		url = "https://mongo-flex-prd.api.eu01.stackit.cloud"
	case PostgresDB:
		url = "https://postgres-flex-service.api.eu01.stackit.cloud"
	case RabbitMQ:
		url = "https://rabbitmq.api.eu01.stackit.cloud"
	case Redis:
		url = "https://redis.api.eu01.stackit.cloud"
	}

	nc, _ := dataservices.NewClientWithResponses(url, dataservices.WithHTTPClient(c))
	return nc
}
