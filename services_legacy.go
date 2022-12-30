package client

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/costs"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership"
	resourceManagement "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-management"
)

// Service management

// ProductiveServices is the struct representing all productive services
type ProductiveServices struct {
	Costs              *costs.CostsService
	Membership         *membership.MembershipService
	ObjectStorage      *objectstorage.ObjectStorageService
	ResourceManagement *resourceManagement.ResourceManagementService
}

// init initializes the client and its services and returns the client
func (c *Client) initLegacyServices() *Client {
	// init productive services
	c.Costs = costs.New(c)
	c.Membership = membership.New(c)
	c.ObjectStorage = objectstorage.New(c)
	c.ResourceManagement = resourceManagement.New(c)

	return c
}
