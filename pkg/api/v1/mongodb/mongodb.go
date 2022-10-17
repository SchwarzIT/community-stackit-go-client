// package mongodb groups together STACKIT MongoDB Flex related functionalities
// such as instances and user management

package mongodb

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb/options"
)

// New returns a new handler for the service
func New(c common.Client) *MongoDBService {
	return &MongoDBService{
		Options: options.New(c),
	}
}

// MongoDBService is the service that handles
// MongoDB Flex related services, such as instances & users
type MongoDBService struct {
	Options *options.MongoDBOptionsService
}
