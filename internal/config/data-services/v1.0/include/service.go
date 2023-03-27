package dataservices

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0"
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

func NewService(c contracts.BaseClientInterface, serviceID int) *dataservices.ClientWithResponses {
	url := GetBaseURLs(serviceID).GetURL(c.GetEnvironment())
	nc, _ := dataservices.NewClient(url, dataservices.WithHTTPClient(c))
	return nc
}

func GetBaseURLs(serviceID int) env.EnvironmentURLs {
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
	return env.EnvironmentURLs{}
}

func setElasticSearchURLs() env.EnvironmentURLs {
	return env.URLs(
		"elasticsearch",
		"https://elasticsearch.api.eu01.stackit.cloud",
		"https://elasticsearch.api.eu01.qa.stackit.cloud",
		"https://elasticsearch.api.eu01.dev.stackit.cloud",
	)
}

func setLogMeURLs() env.EnvironmentURLs {
	return env.URLs(
		"logme",
		"https://logme.api.eu01.stackit.cloud",
		"https://logme.api.eu01.qa.stackit.cloud",
		"https://logme.api.eu01.dev.stackit.cloud",
	)
}

func setMariaDBURLs() env.EnvironmentURLs {
	return env.URLs(
		"mariadb",
		"https://mariadb.api.eu01.stackit.cloud",
		"https://mariadb.api.eu01.qa.stackit.cloud",
		"https://mariadb.api.eu01.dev.stackit.cloud",
	)
}

func setPostgresDBURLs() env.EnvironmentURLs {
	return env.URLs(
		"postgresql",
		"https://postgresql.api.eu01.stackit.cloud",
		"https://postgresql.api.eu01.qa.stackit.cloud",
		"https://postgresql.api.eu01.dev.stackit.cloud",
	)
}

func setRabbitMQURLs() env.EnvironmentURLs {
	return env.URLs(
		"rabbitmq",
		"https://rabbitmq.api.eu01.stackit.cloud",
		"https://rabbitmq.api.eu01.qa.stackit.cloud",
		"https://rabbitmq.api.eu01.dev.stackit.cloud",
	)
}

func setRedisURL() env.EnvironmentURLs {
	return env.URLs(
		"redis",
		"https://redis.api.eu01.stackit.cloud",
		"https://redis.api.eu01.qa.stackit.cloud",
		"https://redis.api.eu01.dev.stackit.cloud",
	)
}
