package iaas

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
)

var BaseURLs = baseurl.New(
	"iaas",
	"https://iaas.api.eu01.stackit.cloud",
)

//func NewService(c contracts.BaseClientInterface) *mongodb.ClientWithResponses {
//	nc, _ := mongodb.NewClient(BaseURLs.Get(), mongodb.WithHTTPClient(c))
//	return nc
//}
