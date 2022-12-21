package project

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
	"github.com/pkg/errors"
)

func (r CreateProjectResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {

		resp, err := c.GetProjectWithResponse(ctx, projectID)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed during create request preparation")
		}
		if resp.HasError != nil {
			if strings.Contains(resp.HasError.Error(), "project has no assigned namespace") {
				return nil, false, nil
			}
			return nil, false, err
		}

		switch *resp.JSON200.State {
		case STATE_FAILED:
			fallthrough
		case STATE_DELETING:
			return nil, false, fmt.Errorf("received state: %s for project ID: %s",
				*resp.JSON200.State,
				*resp.JSON200.ProjectID,
			)
		case STATE_CREATED:
			return nil, true, nil
		}
		return nil, false, nil
	})
}

func (r DeleteProjectResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.GetProjectWithResponse(ctx, projectID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, errors.Wrap(err, "failed during delete request preparation")
		}
		if resp.HasError != nil {
			if strings.Contains(resp.HasError.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	})
}
