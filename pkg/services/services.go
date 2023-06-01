package services

import (
	"errors"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/clients"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v2.0"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1"
	opensearch "github.com/SchwarzIT/community-stackit-go-client/pkg/services/opensearch/v1.1.0"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0"
)

type Services struct {
	Client             contracts.BaseClientInterface
	Argus              *argus.ClientWithResponses
	Costs              *costs.ClientWithResponses
	Kubernetes         *kubernetes.ClientWithResponses
	Membership         *membership.ClientWithResponses
	MongoDBFlex        *mongodbflex.ClientWithResponses
	ObjectStorage      *objectstorage.ClientWithResponses
	Opensearch         *opensearch.ClientWithResponses
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

func Init(c contracts.BaseClientInterface) (*Services, error) {
	nc := newClient(c)
	if nc == nil {
		return nil, errors.New("client cloning failed")
	}

	return &Services{
		Client: c,

		// Services
		Argus:              argus.NewService(nc),
		Costs:              costs.NewService(newClient(c)),
		Kubernetes:         kubernetes.NewService(newClient(c)),
		Membership:         membership.NewService(newClient(c)),
		MongoDBFlex:        mongodbflex.NewService(newClient(c)),
		ObjectStorage:      objectstorage.NewService(newClient(c)),
		Opensearch:         opensearch.NewService(newClient(c)),
		PostgresFlex:       postgresflex.NewService(newClient(c)),
		ResourceManagement: resourcemanagement.NewService(newClient(c)),
		ServiceAccounts:    serviceaccounts.NewService(newClient(c)),

		// DSA
		ElasticSearch: dataservices.NewService(newClient(c), dataservices.ElasticSearch),
		LogMe:         dataservices.NewService(newClient(c), dataservices.LogMe),
		MariaDB:       dataservices.NewService(newClient(c), dataservices.MariaDB),
		PostgresDB:    dataservices.NewService(newClient(c), dataservices.PostgresDB),
		RabbitMQ:      dataservices.NewService(newClient(c), dataservices.RabbitMQ),
		Redis:         dataservices.NewService(newClient(c), dataservices.Redis),
	}, nil
}

func newClient(c contracts.BaseClientInterface) contracts.BaseClientInterface {
	nc := c.Clone()
	if v, ok := nc.(*clients.KeyFlow); ok {
		return contracts.BaseClientInterface(v)
	}
	if v, ok := nc.(*clients.TokenFlow); ok {
		return contracts.BaseClientInterface(v)
	}
	return nil
}
