// package postgres groups together STACKIT Postgres Flex related functionalities
// such as instances and user management

package postgres

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/postgres/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/postgres/options"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/postgres/users"
)

// New returns a new handler for the service
func New(c common.Client) *PostgresService {
	return &PostgresService{
		Instances: instances.New(c),
		Options:   options.New(c),
		Users:     users.New(c),
	}
}

// PostgresService is the service that handles
// Postgres Flex related services, such as instances & users
type PostgresService struct {
	Instances *instances.PostgresInstancesService
	Options   *options.PostgresOptionsService
	Users     *users.PostgresUsersService
}
