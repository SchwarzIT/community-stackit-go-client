package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1"
)

var BaseURLs = baseurl.New(
	"object_storage",
	"https://object-storage.api.eu01.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *objectstorage.ClientWithResponses {
	nc, _ := objectstorage.NewClient(
		BaseURLs.Get(),
		objectstorage.WithHTTPClient(c),
	)
	return nc
}
