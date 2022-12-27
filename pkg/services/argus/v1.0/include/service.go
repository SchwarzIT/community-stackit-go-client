package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	argus "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated"
)

func NewService(c common.Client) *argus.ClientWithResponses {
	nc, _ := argus.NewClientWithResponses("https://argus.api.eu01.stackit.cloud", argus.WithHTTPClient(c))
	return nc
}
