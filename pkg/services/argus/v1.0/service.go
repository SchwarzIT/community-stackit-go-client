package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"argus",
	"https://argus.api.stackit.cloud",
	"https://argus.api.stg.stackit.cloud",
	"https://argus.api.dev.stackit.cloud",
)

func NewService(c common.Client) *ClientWithResponses {
	nc, _ := NewClient(BaseURLs.GetURL(c), WithHTTPClient(c))
	return nc
}
