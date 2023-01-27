package mongodbflex

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	mongodbflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/generated"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/urls"
)

var BaseURLs = urls.Init(
	"mongodb_flex",
	"https://api.stackit.cloud/mongodb/v1/",
	"https://api-qa.stackit.cloud/mongodb/v1/",
	"https://api-dev.stackit.cloud/mongodb/v1/",
)

func NewService(c common.Client) *mongodbflex.ClientWithResponses {
	nc, _ := mongodbflex.NewClientWithResponses(
		BaseURLs.GetURL(c),
		mongodbflex.WithHTTPClient(c),
	)
	return nc
}
