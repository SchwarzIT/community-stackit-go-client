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
	url := GetBaseURLs(serviceID).GetURL(c)
	nc, _ := NewClientWithResponses(url, WithHTTPClient(c))
	return nc
}
func GetBaseURLs(serviceID int) urls.ByEnvs {
	switch serviceID {
	case ElasticSearch:
		return setElasticSearchURLs()
	case LogMe:
		return setLogMeURLs()
	case MariaDB:
		return setMariaDBURLs()
	case PostgresDB:
		return setPostgresDBURLs()
	case RabbitMQ:
		return setRabbitMQURLs()
	case Redis:
		return setRedisURL()
	}
	return urls.ByEnvs{}
}

func setElasticSearchURLs() urls.ByEnvs {
	return urls.Init(
		"elasticsearch",
		"https://elasticsearch.api.eu01.stackit.cloud",
		"https://elasticsearch.api.eu01.qa.stackit.cloud",
		"https://elasticsearch.api.eu01.dev.stackit.cloud",
	)
}

func setLogMeURLs() urls.ByEnvs {
	return urls.Init(
		"logme",
		"https://logme.api.eu01.stackit.cloud",
		"https://logme.api.eu01.qa.stackit.cloud",
		"https://logme.api.eu01.dev.stackit.cloud",
	)
}

func setMariaDBURLs() urls.ByEnvs {
	return urls.Init(
		"mariadb",
		"https://mariadb.api.eu01.stackit.cloud",
		"https://mariadb.api.eu01.qa.stackit.cloud",
		"https://mariadb.api.eu01.dev.stackit.cloud",
	)
}

func setPostgresDBURLs() urls.ByEnvs {
	return urls.Init(
		"postgresql",
		"https://postgresql.api.eu01.stackit.cloud",
		"https://postgresql.api.eu01.qa.stackit.cloud",
		"https://postgresql.api.eu01.dev.stackit.cloud",
	)
}

func setRabbitMQURLs() urls.ByEnvs {
	return urls.Init(
		"rabbitmq",
		"https://rabbitmq.api.eu01.stackit.cloud",
		"https://rabbitmq.api.eu01.qa.stackit.cloud",
		"https://rabbitmq.api.eu01.dev.stackit.cloud",
	)
}

func setRedisURL() urls.ByEnvs {
	return urls.Init(
		"redis",
		"https://redis.api.eu01.stackit.cloud",
		"https://redis.api.eu01.qa.stackit.cloud",
		"https://redis.api.eu01.dev.stackit.cloud",
	)
}
