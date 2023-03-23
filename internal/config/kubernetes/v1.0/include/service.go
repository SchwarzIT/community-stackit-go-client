package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"kubernetes",
	"https://ske.api.eu01.stackit.cloud/",
	"https://ske.api.eu01.stg.stackit.cloud/",
	"https://ske.api.eu01.dev.stackit.cloud/",
)

func NewService(c common.Client) *kubernetes.ClientWithResponses {
	nc, _ := kubernetes.NewClient(
		BaseURLs.GetURL(c),
		kubernetes.WithHTTPClient(c),
	)
	return nc
}
