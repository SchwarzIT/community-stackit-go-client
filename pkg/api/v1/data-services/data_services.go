// package dataservices groups together Data Services Access related functionalities
// such as instances, credentials and offerings

package dataservices

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/credentials"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/options"
)

// New returns a new handler for the service
// broker is the dsa broker (i.e. ElasticSearch or RabbitMQ broker)
func New(c common.Client, broker string) *DataServicesService {
	return &DataServicesService{
		Credentials: *credentials.New(c, broker),
		Instances:   *instances.New(c, broker),
		Options:     *options.New(c, broker),
	}
}

// DataServicesService is the service that handles
// DSA instances, credentials and offerings
type DataServicesService struct {
	Credentials credentials.DSACredentialsService
	Instances   instances.DSAInstancesService
	Options     options.DSAOptionsService
}
