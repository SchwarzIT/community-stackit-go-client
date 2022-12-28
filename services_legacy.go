package client

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/costs"
	mongodbFlex "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership"
	resourceManagement "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-management"
)

// Service management

// ProductiveServices is the struct representing all productive services
type ProductiveServices struct {
	Costs              *costs.CostsService
	Membership         *membership.MembershipService
	MongoDBFlex        *mongodbFlex.MongoDBService
	ObjectStorage      *objectstorage.ObjectStorageService
	ResourceManagement *resourceManagement.ResourceManagementService
}

// IncubatorServices is the struct representing all services that are under development
type IncubatorServices struct {
}

// init initializes the client and its services and returns the client
func (c *Client) initLegacyServices() *Client {
	// init productive services
	c.Costs = costs.New(c)
	c.Membership = membership.New(c)
	c.MongoDBFlex = mongodbFlex.New(c)
	c.ObjectStorage = objectstorage.New(c)
	c.ResourceManagement = resourceManagement.New(c)

	// init incubator services
	c.Incubator = IncubatorServices{}

	return c
}
