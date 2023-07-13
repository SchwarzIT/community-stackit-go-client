package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
)

var BaseURLs = env.URLs(
	"object_storage",
	"https://object-storage-api.api.eu01.stackit.cloud",
	"https://object-storage-api.api.eu01.qa.stackit.cloud",
	"https://object-storage-api.api.eu01.dev.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.GetURL(c.GetEnvironment()),
		WithHTTPClient(c),
	)
	return nc
}
