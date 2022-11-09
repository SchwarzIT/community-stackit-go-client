// package dataservices groups together Data Services Access related functionalities
// such as instances, credentials and offerings

package dataservices

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/credentials"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/options"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// define supported serices
const (
	SERVICE_ELASTICSEARCH = iota
	SERVICE_LOGME
	SERVICE_MARIADB
	SERVICE_POSTGRES
	SERVICE_RABBITMQ
	SERVICE_REDIS
)

// New returns a new handler for the service
// broker is the dsa broker (i.e. ElasticSearch or RabbitMQ broker)
func New(c common.Client, service int, overrideBaseURL string) *DataServicesService {
	// modify client base url
	nc := c.Clone()
	setBaseURL(nc, service, overrideBaseURL)

	// retur an initialized data services service
	return &DataServicesService{
		Credentials: *credentials.New(nc),
		Instances:   *instances.New(nc),
		Options:     *options.New(nc),
	}
}

// DataServicesService is the service that handles
// DSA instances, credentials and offerings
type DataServicesService struct {
	Credentials credentials.DSACredentialsService
	Instances   instances.DSAInstancesService
	Options     options.DSAOptionsService
}

func setBaseURL(c common.Client, service int, overrideBaseURL string) {
	if overrideBaseURL != "" {
		_ = c.SetBaseURL(overrideBaseURL)
		return
	}
	switch service {
	case SERVICE_ELASTICSEARCH:
		_ = c.SetBaseURL(consts.API_BASEURL_DSA_ELASTICSEARCH)
	case SERVICE_LOGME:
		_ = c.SetBaseURL(consts.API_BASEURL_DSA_LOGME)
	case SERVICE_MARIADB:
		_ = c.SetBaseURL(consts.API_BASEURL_DSA_MARIADB)
	case SERVICE_POSTGRES:
		_ = c.SetBaseURL(consts.API_BASEURL_DSA_POSTGRES)
	case SERVICE_RABBITMQ:
		_ = c.SetBaseURL(consts.API_BASEURL_DSA_RABBITMQ)
	case SERVICE_REDIS:
		_ = c.SetBaseURL(consts.API_BASEURL_DSA_REDIS)
	}
}
