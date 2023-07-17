package dataservices

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0"
)

const (
	ElasticSearch = iota
	LogMe
	MariaDB
	MongoDB
	Opensearch
	PostgresDB
	RabbitMQ
	Redis
)

func NewService(c contracts.BaseClientInterface, serviceID int) *dataservices.ClientWithResponses {
	url := GetBaseURLs(serviceID).Get()
	nc, _ := dataservices.NewClient(url, dataservices.WithHTTPClient(c))
	return nc
}

func GetBaseURLs(serviceID int) baseurl.BaseURL {
	switch serviceID {
	case ElasticSearch:
		return setElasticSearchURLs()
	case LogMe:
		return setLogMeURLs()
	case MariaDB:
		return setMariaDBURLs()
	case Opensearch:
		return setOpensearchURLs()
	case PostgresDB:
		return setPostgresDBURLs()
	case RabbitMQ:
		return setRabbitMQURLs()
	case Redis:
		return setRedisURL()
	}
	return baseurl.BaseURL{}
}

func setElasticSearchURLs() baseurl.BaseURL {
	return baseurl.New(
		"elasticsearch",
		"https://elasticsearch.api.eu01.stackit.cloud",
	)
}

func setLogMeURLs() baseurl.BaseURL {
	return baseurl.New(
		"logme",
		"https://logme.api.eu01.stackit.cloud",
	)
}

func setMariaDBURLs() baseurl.BaseURL {
	return baseurl.New(
		"mariadb",
		"https://mariadb.api.eu01.stackit.cloud",
	)
}

func setOpensearchURLs() baseurl.BaseURL {
	return baseurl.New(
		"redis",
		"https://opensearch.api.eu01.stackit.cloud",
	)
}

func setPostgresDBURLs() baseurl.BaseURL {
	return baseurl.New(
		"postgresql",
		"https://postgresql.api.eu01.stackit.cloud",
	)
}

func setRabbitMQURLs() baseurl.BaseURL {
	return baseurl.New(
		"rabbitmq",
		"https://rabbitmq.api.eu01.stackit.cloud",
	)
}

func setRedisURL() baseurl.BaseURL {
	return baseurl.New(
		"redis",
		"https://redis.api.eu01.stackit.cloud",
	)
}
