package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0"
)

var BaseURLs = baseurl.New(
	"argus",
	"https://argus.api.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *argus.ClientWithResponses {
	nc, _ := argus.NewClient(BaseURLs.Get(), argus.WithHTTPClient(c))
	return nc
}
