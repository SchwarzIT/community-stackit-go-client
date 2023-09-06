package loadbalancer

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	loadbalancer "github.com/SchwarzIT/community-stackit-go-client/pkg/services/load-balancer/1beta.0.0"
)

var BaseURLs = baseurl.New(
	"load_balancer",
	"https://load-balancer.api.eu01.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *loadbalancer.ClientWithResponses {
	nc, _ := loadbalancer.NewClient(
		BaseURLs.Get(),
		loadbalancer.WithHTTPClient(c),
	)
	return nc
}
