package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"argus",
	"https://argus.api.stackit.cloud",
	"https://argus.api.stg.stackit.cloud",
	"https://argus.api.dev.stackit.cloud",
)

func NewService(c common.Client) *argus.ClientWithResponses {
	nc, _ := argus.NewClient(BaseURLs.GetURL(c), argus.WithHTTPClient(c))
	return nc
}
