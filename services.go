package client

import (
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v1"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0/generated"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0/generated"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/generated"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated"
)

type services struct {
	Argus              *argus.ClientWithResponses
	Kubernetes         *kubernetes.ClientWithResponses
	Membership         *membership.ClientWithResponses
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
	Redis         *dataservices.ClientWithResponses

	// Non generated
	Costs         *costs.CostsService
	ObjectStorage *objectstorage.ObjectStorageService
}

func (c *Client) initServices() {
	c.Argus = argus.NewService(c)
	c.Kubernetes = kubernetes.NewService(c)
	c.Membership = membership.NewService(c)
	c.MongoDBFlex = mongodbflex.NewService(c)
	c.PostgresFlex = postgresflex.NewService(c)
	c.ResourceManagement = resourcemanagement.NewService(c)

	// DSA
	c.ElasticSearch = dataservices.NewService(c, dataservices.ElasticSearch)
	c.LogMe = dataservices.NewService(c, dataservices.LogMe)
	c.MariaDB = dataservices.NewService(c, dataservices.MariaDB)
	c.PostgresDB = dataservices.NewService(c, dataservices.PostgresDB)
	c.RabbitMQ = dataservices.NewService(c, dataservices.RabbitMQ)
	c.Redis = dataservices.NewService(c, dataservices.Redis)

	// Non Generated
	c.Costs = costs.New(c)
	c.ObjectStorage = objectstorage.New(c)
}
