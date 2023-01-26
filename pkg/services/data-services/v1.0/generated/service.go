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

	nc, _ := NewClientWithResponses(baseURLs.GetURL(c), WithHTTPClient(c))
	return nc
}

func (*ClientWithResponses) GetBaseURLs() *urls.ByEnvs {
	return baseURLs
}

var baseURLs *urls.ByEnvs

func setElasticSearchURLs() {
	baseURLs = urls.Init(
		"elasticsearch",
		"https://elasticsearch.api.eu01.stackit.cloud",
		"https://elasticsearch.api.eu01.qa.stackit.cloud",
		"https://elasticsearch.api.eu01.dev.stackit.cloud",
	)
}

func setLogMeURLs() {
	baseURLs = urls.Init(
		"logme",
		"https://logme.api.eu01.stackit.cloud",
		"https://logme.api.eu01.qa.stackit.cloud",
		"https://logme.api.eu01.dev.stackit.cloud",
	)
}

func setMariaDBURLs() {
	baseURLs = urls.Init(
		"mariadb",
		"https://mariadb.api.eu01.stackit.cloud",
		"https://mariadb.api.eu01.qa.stackit.cloud",
		"https://mariadb.api.eu01.dev.stackit.cloud",
	)
}

func setPostgresDBURLs() {
	baseURLs = urls.Init(
		"postgresql",
		"https://postgresql.api.eu01.stackit.cloud",
		"https://postgresql.api.eu01.qa.stackit.cloud",
		"https://postgresql.api.eu01.dev.stackit.cloud",
	)
}

func setRabbitMQURLs() {
	baseURLs = urls.Init(
		"rabbitmq",
		"https://rabbitmq.api.eu01.stackit.cloud",
		"https://rabbitmq.api.eu01.qa.stackit.cloud",
		"https://rabbitmq.api.eu01.dev.stackit.cloud",
	)
}

func setRedisURL() {
	baseURLs = urls.Init(
		"redis",
		"https://redis.api.eu01.stackit.cloud",
		"https://redis.api.eu01.qa.stackit.cloud",
		"https://redis.api.eu01.dev.stackit.cloud",
	)
}
