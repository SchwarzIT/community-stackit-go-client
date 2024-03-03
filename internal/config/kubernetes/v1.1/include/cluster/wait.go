package cluster

import (
	"context"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.1/cluster"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

func (*CreateOrUpdateResponse) WaitHandler(ctx context.Context, c *cluster.ClientWithResponses, projectID, clusterName string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.Get(ctx, projectID, clusterName)
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
		if resp.Error != nil {
			return nil, false, resp.Error
		}

		status := *resp.JSON200.Status.Aggregated
		if status == cluster.STATE_HEALTHY || status == cluster.STATE_HIBERNATED {
			return resp, true, nil
		}
		return resp, false, nil
	})
}

func (*DeleteResponse) WaitHandler(ctx context.Context, c *cluster.ClientWithResponses, projectID, clusterName string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.Get(ctx, projectID, clusterName)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		if resp.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if resp.Error != nil {
			if resp.StatusCode() == http.StatusNotFound {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	})
}
