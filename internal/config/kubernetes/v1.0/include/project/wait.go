package project

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/project"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
	"github.com/pkg/errors"
)

func (r CreateResponse) WaitHandler(ctx context.Context, c *project.ClientWithResponses, projectID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {

		resp, err := c.Get(ctx, projectID)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed during create request preparation")
		}
		if resp.Error != nil {
			if strings.Contains(resp.Error.Error(), "project has no assigned namespace") {
				return nil, false, nil
			}
			return nil, false, err
		}

		switch *resp.JSON200.State {
		case project.STATE_FAILED:
			fallthrough
		case project.STATE_DELETING:
			return nil, false, fmt.Errorf("received state: %s for project ID: %s",
				*resp.JSON200.State,
				*resp.JSON200.ProjectID,
			)
		case project.STATE_CREATED:
			return nil, true, nil
		}
		return nil, false, nil
	})
}

func (r DeleteResponse) WaitHandler(ctx context.Context, c *project.ClientWithResponses, projectID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.Get(ctx, projectID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, errors.Wrap(err, "failed during delete request preparation")
		}
		if resp.Error != nil {
			if strings.Contains(resp.Error.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	})
}
