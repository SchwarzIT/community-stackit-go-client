package client

import (
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v2.0/generated"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0/generated"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0/generated"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/generated"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/generated"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0/generated"
)

type services struct {
	Argus              *argus.ClientWithResponses
	Costs              *costs.ClientWithResponses
	Kubernetes         *kubernetes.ClientWithResponses
	Membership         *membership.ClientWithResponses
	MongoDBFlex        *mongodbflex.ClientWithResponses
	ObjectStorage      *objectstorage.ClientWithResponses
	PostgresFlex       *postgresflex.ClientWithResponses
	ResourceManagement *resourcemanagement.ClientWithResponses
	ServiceAccounts    *serviceaccounts.ClientWithResponses

	// DSA
	ElasticSearch *dataservices.ClientWithResponses
	LogMe         *dataservices.ClientWithResponses
	MariaDB       *dataservices.ClientWithResponses
	MongoDB       *dataservices.ClientWithResponses
	PostgresDB    *dataservices.ClientWithResponses
	RabbitMQ      *dataservices.ClientWithResponses
	Redis         *dataservices.ClientWithResponses
}

func (c *Client) initServices() {
	c.Argus = argus.NewService(c)
	c.Costs = costs.NewService(c)
	c.Kubernetes = kubernetes.NewService(c)
	c.Membership = membership.NewService(c)
	c.MongoDBFlex = mongodbflex.NewService(c)
	c.ObjectStorage = objectstorage.NewService(c)
	c.PostgresFlex = postgresflex.NewService(c)
	c.ResourceManagement = resourcemanagement.NewService(c)
	c.ServiceAccounts = serviceaccounts.NewService(c)

	// DSA
	c.ElasticSearch = dataservices.NewService(c, dataservices.ElasticSearch)
	c.LogMe = dataservices.NewService(c, dataservices.LogMe)
	c.MariaDB = dataservices.NewService(c, dataservices.MariaDB)
	c.PostgresDB = dataservices.NewService(c, dataservices.PostgresDB)
	c.RabbitMQ = dataservices.NewService(c, dataservices.RabbitMQ)
	c.Redis = dataservices.NewService(c, dataservices.Redis)
}
