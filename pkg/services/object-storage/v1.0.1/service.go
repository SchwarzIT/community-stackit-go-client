package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"object_storage",
	"https://api.stackit.cloud/object-storage-api/",
	"https://api-qa.stackit.cloud/object-storage-api/",
	"https://api-dev.stackit.cloud/object-storage-api/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
