package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/generated"
)

func NewService(c common.Client) *objectstorage.ClientWithResponses {
	nc, _ := objectstorage.NewClientWithResponses(
		"https://api.stackit.cloud/object-storage-api/",
		objectstorage.WithHTTPClient(c),
	)
	return nc
}
