// package mongodb groups together STACKIT MongoDB Flex related functionalities
// such as instances and user management

package mongodb

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex/options"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex/users"
)

// New returns a new handler for the service
func New(c common.Client) *MongoDBService {
	return &MongoDBService{
		Instances: instances.New(c),
		Options:   options.New(c),
		Users:     users.New(c),
	}
}

// MongoDBService is the service that handles
// MongoDB Flex related services, such as instances & users
type MongoDBService struct {
	Instances *instances.MongoDBInstancesService
	Options   *options.MongoDBOptionsService
	Users     *users.MongoDBUsersService
}
