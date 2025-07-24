package scf

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	scf "github.com/SchwarzIT/community-stackit-go-client/pkg/services/scf/v1.0"
)

var BaseURLs = baseurl.New(
	"scf",
	"https://scf.api.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *scf.ClientWithResponses {
	nc, _ := scf.NewClient(BaseURLs.Get(), scf.WithHTTPClient(c))
	return nc
}
