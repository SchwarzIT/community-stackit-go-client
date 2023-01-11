package dataservices

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
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

func NewService(c common.Client, serviceID int) *ClientWithResponses {
	url := consts.DEFAULT_BASE_URL

	switch serviceID {
	case ElasticSearch:
		url = "https://elasticsearch.api.eu01.stackit.cloud"
	case LogMe:
		url = "https://logme.api.eu01.stackit.cloud"
	case MariaDB:
		url = "https://mariadb.api.eu01.stackit.cloud"
	case PostgresDB:
		url = "https://postgres.api.eu01.stackit.cloud"
	case RabbitMQ:
		url = "https://rabbitmq.api.eu01.stackit.cloud"
	case Redis:
		url = "https://redis.api.eu01.stackit.cloud"
	}

	nc, _ := NewClientWithResponses(url, WithHTTPClient(c))
	return nc
}
