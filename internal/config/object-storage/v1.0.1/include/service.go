package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1"
)

var BaseURLs = env.URLs(
	"object_storage",
	"https://api.stackit.cloud/object-storage-api/",
	"https://api-qa.stackit.cloud/object-storage-api/",
	"https://api-dev.stackit.cloud/object-storage-api/",
)

func NewService[K contracts.ClientFlowConfig](c contracts.ClientInterface[K]) *objectstorage.ClientWithResponses[K] {
	nc, _ := objectstorage.NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		objectstorage.WithHTTPClient(c),
	)
	return nc
}
