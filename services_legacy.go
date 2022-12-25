package client

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/costs"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services"
	mongodbFlex "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage"
	postgresFlex "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/postgres-flex"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership"
	resourceManagement "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-management"
)

// Service management

// ProductiveServices is the struct representing all productive services
type ProductiveServices struct {
	Argus              *argus.ArgusService
	Costs              *costs.CostsService
	DataServices       DataServices
	Membership         *membership.MembershipService
	MongoDBFlex        *mongodbFlex.MongoDBService
	ObjectStorage      *objectstorage.ObjectStorageService
	PostgresFlex       *postgresFlex.PostgresService
	ResourceManagement *resourceManagement.ResourceManagementService
}

// IncubatorServices is the struct representing all services that are under development
type IncubatorServices struct {
}

// Extras

type DataServices struct {
	ElasticSearch *dataservices.DataServicesService
	LogMe         *dataservices.DataServicesService
	MariaDB       *dataservices.DataServicesService
	PostgresDB    *dataservices.DataServicesService
	RabbitMQ      *dataservices.DataServicesService
	Redis         *dataservices.DataServicesService
}

// init initializes the client and its services and returns the client
func (c *Client) initLegacyServices() *Client {
	// init productive services
	c.Argus = argus.New(c)
	c.Costs = costs.New(c)
	c.Membership = membership.New(c)
	c.MongoDBFlex = mongodbFlex.New(c)
	c.ObjectStorage = objectstorage.New(c)
	c.ResourceManagement = resourceManagement.New(c)
	c.PostgresFlex = postgresFlex.New(c)

	c.DataServices = DataServices{
		ElasticSearch: dataservices.New(c, dataservices.SERVICE_ELASTICSEARCH, ""),
		LogMe:         dataservices.New(c, dataservices.SERVICE_LOGME, ""),
		MariaDB:       dataservices.New(c, dataservices.SERVICE_MARIADB, ""),
		PostgresDB:    dataservices.New(c, dataservices.SERVICE_POSTGRES, ""),
		RabbitMQ:      dataservices.New(c, dataservices.SERVICE_RABBITMQ, ""),
		Redis:         dataservices.New(c, dataservices.SERVICE_REDIS, ""),
	}

	// init incubator services
	c.Incubator = IncubatorServices{}

	return c
}
