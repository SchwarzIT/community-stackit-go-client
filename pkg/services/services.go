package services

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v2.0"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/services/data-services/v1.0"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1"
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

func Init[K contracts.ClientFlowConfig](c contracts.BaseClientInterface) *Services {
	return &Services{
		Client: c,

		// Services
		Argus:              argus.NewService(c),
		Costs:              costs.NewService(c),
		Kubernetes:         kubernetes.NewService(c),
		Membership:         membership.NewService(c),
		MongoDBFlex:        mongodbflex.NewService(c),
		ObjectStorage:      objectstorage.NewService(c),
		PostgresFlex:       postgresflex.NewService(c),
		ResourceManagement: resourcemanagement.NewService(c),
		ServiceAccounts:    serviceaccounts.NewService(c),

		// DSA
		ElasticSearch: dataservices.NewService(c, dataservices.ElasticSearch),
		LogMe:         dataservices.NewService(c, dataservices.LogMe),
		MariaDB:       dataservices.NewService(c, dataservices.MariaDB),
		PostgresDB:    dataservices.NewService(c, dataservices.PostgresDB),
		RabbitMQ:      dataservices.NewService(c, dataservices.RabbitMQ),
		Redis:         dataservices.NewService(c, dataservices.Redis),
	}
}
