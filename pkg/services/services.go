package services

import (
	"errors"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/clients"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v2.0"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0"
	loadbalancer "github.com/SchwarzIT/community-stackit-go-client/pkg/services/load-balancer/1beta.0.0"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
	secretsmanager "github.com/SchwarzIT/community-stackit-go-client/pkg/services/secrets-manager/v1.1.0"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0"
)

type Services struct {
	Client             contracts.BaseClientInterface
	Argus              *argus.ClientWithResponses
	Costs              *costs.ClientWithResponses
	Kubernetes         *kubernetes.ClientWithResponses
	LoadBalancer       *loadbalancer.ClientWithResponses
	Membership         *membership.ClientWithResponses
	MongoDBFlex        *mongodbflex.ClientWithResponses
	ObjectStorage      *objectstorage.ClientWithResponses
	PostgresFlex       *postgresflex.ClientWithResponses
	ResourceManagement *resourcemanagement.ClientWithResponses
	SecretsManager     *secretsmanager.ClientWithResponses
	ServiceAccounts    *serviceaccounts.ClientWithResponses

	// DSA
	ElasticSearch *dataservices.ClientWithResponses
	LogMe         *dataservices.ClientWithResponses
	MariaDB       *dataservices.ClientWithResponses
	MongoDB       *dataservices.ClientWithResponses
	Opensearch    *dataservices.ClientWithResponses
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
		PostgresFlex:       postgresflex.NewService(newClient(c)),
		ResourceManagement: resourcemanagement.NewService(newClient(c)),
		ServiceAccounts:    serviceaccounts.NewService(newClient(c)),
		SecretsManager:     secretsmanager.NewService(newClient(c)),

		// DSA
		ElasticSearch: dataservices.NewService(newClient(c), dataservices.ElasticSearch),
		LogMe:         dataservices.NewService(newClient(c), dataservices.LogMe),
		MariaDB:       dataservices.NewService(newClient(c), dataservices.MariaDB),
		Opensearch:    dataservices.NewService(newClient(c), dataservices.Opensearch),
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
