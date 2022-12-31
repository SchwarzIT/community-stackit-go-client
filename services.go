package client

import (
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0/generated"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/generated"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated"
)

type services struct {
	Argus              *argus.ClientWithResponses
	Kubernetes         *kubernetes.ClientWithResponses
	MongoDBFlex        *mongodbflex.ClientWithResponses
	PostgresFlex       *postgresflex.ClientWithResponses
	ResourceManagement *resourcemanagement.ClientWithResponses

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
	c.Services.Argus = argus.NewService(c)
	c.Services.Kubernetes = kubernetes.NewService(c)
	c.Services.MongoDBFlex = mongodbflex.NewService(c)
	c.Services.PostgresFlex = postgresflex.NewService(c)
	c.Services.ResourceManagement = resourcemanagement.NewService(c)

	// DSA
	c.Services.ElasticSearch = dataservices.NewService(c, dataservices.ElasticSearch)
	c.Services.LogMe = dataservices.NewService(c, dataservices.LogMe)
	c.Services.MariaDB = dataservices.NewService(c, dataservices.MariaDB)
	c.Services.MongoDB = dataservices.NewService(c, dataservices.MongoDB)
	c.Services.PostgresDB = dataservices.NewService(c, dataservices.PostgresDB)
	c.Services.RabbitMQ = dataservices.NewService(c, dataservices.RabbitMQ)
	c.Services.Reddis = dataservices.NewService(c, dataservices.Reddis)
}
