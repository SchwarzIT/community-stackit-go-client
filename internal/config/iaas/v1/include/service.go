package iaas

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/iaas-api/v1"
)

var BaseURLs = baseurl.New(
	"iaas",
	"https://iaas.api.eu01.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) (*iaas.ClientWithResponses, error) {
	return iaas.NewClient(BaseURLs.Get(), c)
}
