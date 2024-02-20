package iaas

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

var BaseURLs = baseurl.New(
	"kubernetes",
	"https://ske.api.eu01.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	return NewClient(BaseURLs.Get(), c)
}
