package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1"
)

var BaseURLs = env.URLs(
	"object_storage",
	"https://object-storage-api.api.eu01.stackit.cloud",
	"https://object-storage-api.api.eu01.qa.stackit.cloud",
	"https://object-storage-api.api.eu01.dev.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *objectstorage.ClientWithResponses {
	nc, _ := objectstorage.NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		objectstorage.WithHTTPClient(c),
	)
	return nc
}
