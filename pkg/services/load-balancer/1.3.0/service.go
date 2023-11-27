package loadbalancer

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
)

var BaseURLs = baseurl.New(
	"load_balancer",
	"https://load-balancer.api.eu01.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *ClientWithResponses {
	nc, _ := NewClient(
		BaseURLs.Get(),
		WithHTTPClient(c),
	)
	return nc
}
