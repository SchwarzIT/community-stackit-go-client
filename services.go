package client

import (
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0/generated"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/generated"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
)

type services struct {
	Kubernetes   *kubernetes.ClientWithResponses
	PostgresFlex *postgresflex.ClientWithResponses
	MongoDBFlex  *mongodbflex.ClientWithResponses

	// DSA
	ElasticSearch *dataservices.ClientWithResponses
	LogMe         *dataservices.ClientWithResponses
	MariaDB       *dataservices.ClientWithResponses
	MongoDB       *dataservices.ClientWithResponses
	PostgresDB    *dataservices.ClientWithResponses
	RabbitMQ      *dataservices.ClientWithResponses
	Reddis        *dataservices.ClientWithResponses
}

func (c *Client) initServices() {
	c.Services.Kubernetes = kubernetes.NewService(c)
	c.Services.PostgresFlex = postgresflex.NewService(c)
	c.Services.MongoDBFlex = mongodbflex.NewService(c)

	// DSA
	c.Services.ElasticSearch = dataservices.NewService(c, dataservices.ElasticSearch)
	c.Services.LogMe = dataservices.NewService(c, dataservices.LogMe)
	c.Services.MariaDB = dataservices.NewService(c, dataservices.MariaDB)
	c.Services.MongoDB = dataservices.NewService(c, dataservices.MongoDB)
	c.Services.PostgresDB = dataservices.NewService(c, dataservices.PostgresDB)
	c.Services.RabbitMQ = dataservices.NewService(c, dataservices.RabbitMQ)
	c.Services.Reddis = dataservices.NewService(c, dataservices.Reddis)
}
