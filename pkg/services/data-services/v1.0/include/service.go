package dataservices

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
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
	url := common.DEFAULT_BASE_URL

	switch serviceID {
	case ElasticSearch:
		url = getElasticSearchURL(c)
	case LogMe:
		url = getLogMeURL(c)
	case MariaDB:
		url = getMariaDBURL(c)
	case PostgresDB:
		url = getPostgresDBURL(c)
	case RabbitMQ:
		url = getRabbitMQURL(c)
	case Redis:
		url = getRedisURL(c)
	}

	nc, _ := dataservices.NewClientWithResponses(url, dataservices.WithHTTPClient(c))
	return nc
}

func getElasticSearchURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://elasticsearch.api.eu01.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://elasticsearch.api.eu01.qa.stackit.cloud"
	default:
		return "https://elasticsearch.api.eu01.stackit.cloud"
	}
}

func getLogMeURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://logme.api.eu01.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://logme.api.eu01.qa.stackit.cloud"
	default:
		return "https://logme.api.eu01.stackit.cloud"
	}
}

func getMariaDBURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://mariadb.api.eu01.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://mariadb.api.eu01.qa.stackit.cloud"
	default:
		return "https://mariadb.api.eu01.stackit.cloud"
	}
}

func getPostgresDBURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://postgresql.api.eu01.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://postgresql.api.eu01.qa.stackit.cloud"
	default:
		return "https://postgresql.api.eu01.stackit.cloud"
	}
}

func getRabbitMQURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://rabbitmq.api.eu01.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://rabbitmq.api.eu01.qa.stackit.cloud"
	default:
		return "https://rabbitmq.api.eu01.stackit.cloud"
	}
}

func getRedisURL(c common.Client) string {
	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return "https://redis.api.eu01.dev.stackit.cloud"
	case common.ENV_QA:
		return "https://redis.api.eu01.qa.stackit.cloud"
	default:
		return "https://redis.api.eu01.stackit.cloud"
	}
}
