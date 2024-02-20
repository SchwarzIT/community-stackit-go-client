package iaas

import (
	"context"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"

	openapiTypes "github.com/SchwarzIT/community-stackit-go-client/pkg/helpers/types"
)

func (*V1CreateNetworkResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, networkID openapiTypes.UUID) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.V1GetNetwork(ctx, projectID, networkID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusForbidden)) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if resp.StatusCode() == http.StatusForbidden {
			return nil, false, nil
		}
		if resp.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if resp.JSON200 == nil {
			return nil, false, nil
		}

		return resp, true, nil
	})
}
