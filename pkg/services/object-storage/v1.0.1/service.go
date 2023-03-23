package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"object_storage",
	"https://api.stackit.cloud/object-storage-api/",
	"https://api-qa.stackit.cloud/object-storage-api/",
	"https://api-dev.stackit.cloud/object-storage-api/",
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.GetURL(c),
		WithHTTPClient(c),
	)
	return nc
}
