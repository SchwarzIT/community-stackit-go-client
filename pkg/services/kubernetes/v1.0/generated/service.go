package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"kubernetes",
	"https://ske.api.eu01.stackit.cloud/",
	"https://ske.api.eu01.stg.stackit.cloud/",
	"https://ske.api.eu01.dev.stackit.cloud/",
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClientWithResponses(
		BaseURLs.GetURL(c),
		WithHTTPClient(c),
	)
	return nc
}
