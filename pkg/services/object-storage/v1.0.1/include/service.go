package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/generated"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"object_storage",
	"https://api.stackit.cloud/object-storage-api/",
	"https://api-qa.stackit.cloud/object-storage-api/",
	"https://api-dev.stackit.cloud/object-storage-api/",
)

func NewService(c common.Client) *objectstorage.ClientWithResponses {
	nc, _ := objectstorage.NewClientWithResponses(
		BaseURLs.GetURL(c),
		objectstorage.WithHTTPClient(c),
	)
	return nc
}
