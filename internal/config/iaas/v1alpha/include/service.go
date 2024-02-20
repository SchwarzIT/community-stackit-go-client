package iaas

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	iaas "github.com/SchwarzIT/community-stackit-go-client/pkg/services/iaas-api/v1alpha"
)

var BaseURLs = baseurl.New(
	"iaas",
	"https://iaas.api.eu01.stackit.cloud/",
)

func NewService(c contracts.BaseClientInterface) *iaas.ClientWithResponses {
	return iaas.NewClient(BaseURLs.Get(), c)
}
