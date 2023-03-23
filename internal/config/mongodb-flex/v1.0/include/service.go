package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	mongodb "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"mongodb_flex",
	"https://api.stackit.cloud/mongodb/v1/",
	"https://api-qa.stackit.cloud/mongodb/v1/",
	"https://api-dev.stackit.cloud/mongodb/v1/",
)

func NewService(c common.Client) *mongodb.ClientWithResponses {
	nc, _ := mongodb.NewClient(
		BaseURLs.GetURL(c),
		mongodb.WithHTTPClient(c),
	)
	return nc
}
