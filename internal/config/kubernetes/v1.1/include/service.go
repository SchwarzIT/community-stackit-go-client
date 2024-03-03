package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0"
)

var BaseURLs = baseurl.New(
	"kubernetes",
	"https://ske.api.eu01.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) *kubernetes.ClientWithResponses {
	nc, _ := kubernetes.NewClient(BaseURLs.Get(), kubernetes.WithHTTPClient(c))
	return nc
}
