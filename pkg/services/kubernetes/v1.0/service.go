package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

var BaseURLs = baseurl.New(
	"kubernetes",
	"https://ske.api.eu01.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(BaseURLs.Get(), WithHTTPClient(c))
	return nc
}
