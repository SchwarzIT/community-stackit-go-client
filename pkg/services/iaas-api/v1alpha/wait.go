package iaas

import (
	"context"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
	"github.com/pkg/errors"

	openapiTypes "github.com/SchwarzIT/community-stackit-go-client/pkg/helpers/types"
)

// WaitHandler wait for the network to be created and return it.
func (*V1CreateNetworkResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID openapiTypes.UUID, name string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.V1ListNetworksInProject(ctx, projectID)
		if err != nil {
			return nil, false, err
		}

		if resp.JSON200 != nil && len(resp.JSON200.Items) > 0 {
			for _, n := range resp.JSON200.Items {
				if n.Name == name {
					// the network is created successfully
					return n, true, nil
				}
			}

			// the network that is created was not found
			return nil, false, nil
		}

		if resp.JSON400 != nil {
			// if the server returns 400 then we can't retry the same request because the result will be the same
			return nil, false, errors.New(resp.JSON400.Msg)
		}

		if resp.JSON401 != nil {
			// if the server returns 401 then we can't retry the same request because the result will be the same.
			return nil, false, errors.New(resp.JSON401.Msg)
		}

		if resp.JSON403 != nil {
			// if the server returns 403 then we can't retry the same request because the result will be the same
			return nil, false, errors.New(resp.JSON403.Msg)
		}

		// in all other cases we will retry the request until the network is not created or an error occurred.
		return nil, false, nil
	})
}

// WaitHandler wait for the network to be deleted
func (*V1DeleteNetworkResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, networkID openapiTypes.UUID) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.V1GetNetwork(ctx, projectID, networkID)
		if err != nil {
			return nil, false, err
		}

		if resp.JSON404 != nil {
			// the network is deleted successfully
			return resp, true, nil
		}

		if resp.JSON400 != nil {
			// can't retry the same request because the response will be the same
			return nil, false, errors.New(resp.JSON400.Msg)
		}

		if resp.JSON401 != nil {
			// can't retry the same request because the response will be always the same
			return nil, false, errors.New(resp.JSON401.Msg)
		}

		if resp.JSON403 != nil {
			// can't retry the same request because the response will be always the same
			return nil, false, errors.New(resp.JSON403.Msg)
		}

		if resp.StatusCode() == http.StatusConflict {
			// can't delete network. It is still has systems connected to it.
			return nil, false, errors.New(resp.JSON403.Msg)
		}

		// in all other cases we will retry the request until the network is not deleted or an error occurred.
		return nil, false, nil
	})
}
