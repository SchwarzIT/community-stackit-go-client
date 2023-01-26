package dataservices

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/internal/urls"
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
	switch serviceID {
	case ElasticSearch:
		setElasticSearchURLs()
	case LogMe:
		setLogMeURLs()
	case MariaDB:
		setMariaDBURLs()
	case PostgresDB:
		setPostgresDBURLs()
	case RabbitMQ:
		setRabbitMQURLs()
	case Redis:
		setRedisURL()
	}

	nc, _ := NewClientWithResponses(BaseURLs.GetURL(c), WithHTTPClient(c))
	return nc
}

var BaseURLs *urls.ByEnvs

func setElasticSearchURLs() {
	BaseURLs = urls.Init(
		"elasticsearch",
		"https://elasticsearch.api.eu01.stackit.cloud",
		"https://elasticsearch.api.eu01.qa.stackit.cloud",
		"https://elasticsearch.api.eu01.dev.stackit.cloud",
	)
}

func setLogMeURLs() {
	BaseURLs = urls.Init(
		"logme",
		"https://logme.api.eu01.stackit.cloud",
		"https://logme.api.eu01.qa.stackit.cloud",
		"https://logme.api.eu01.dev.stackit.cloud",
	)
}

func setMariaDBURLs() {
	BaseURLs = urls.Init(
		"mariadb",
		"https://mariadb.api.eu01.stackit.cloud",
		"https://mariadb.api.eu01.qa.stackit.cloud",
		"https://mariadb.api.eu01.dev.stackit.cloud",
	)
}

func setPostgresDBURLs() {
	BaseURLs = urls.Init(
		"postgresql",
		"https://postgresql.api.eu01.stackit.cloud",
		"https://postgresql.api.eu01.qa.stackit.cloud",
		"https://postgresql.api.eu01.dev.stackit.cloud",
	)
}

func setRabbitMQURLs() {
	BaseURLs = urls.Init(
		"rabbitmq",
		"https://rabbitmq.api.eu01.stackit.cloud",
		"https://rabbitmq.api.eu01.qa.stackit.cloud",
		"https://rabbitmq.api.eu01.dev.stackit.cloud",
	)
}

func setRedisURL() {
	BaseURLs = urls.Init(
		"redis",
		"https://redis.api.eu01.stackit.cloud",
		"https://redis.api.eu01.qa.stackit.cloud",
		"https://redis.api.eu01.dev.stackit.cloud",
	)
}
