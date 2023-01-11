package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(
		"https://api.stackit.cloud/object-storage-api/",
		WithHTTPClient(c),
	)
	return nc
}
