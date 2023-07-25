package secretsmanager

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/baseurl"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/contracts"
	secretsmanager "github.com/SchwarzIT/community-stackit-go-client/pkg/services/secrets-manager/v1.1.0"
)

var BaseURLs = baseurl.New(
	"secrets_manager",
	"https://secrets-manager.api.eu01.stackit.cloud",
)

func NewService(c contracts.BaseClientInterface) *secretsmanager.ClientWithResponses {
	nc, _ := secretsmanager.NewClient(
		BaseURLs.Get(),
		secretsmanager.WithHTTPClient(c),
	)
	return nc
}
