package projects

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated/projects"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

func (r CreateResponse) WaitHandler(ctx context.Context, c *projects.ClientWithResponses, containerID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		project, err := c.GetWithResponse(ctx, containerID, &projects.GetParams{})
		if err != nil {
			return project, false, err
		}
		if project.StatusCode() == http.StatusForbidden {
			return project, false, nil
		}
		if project.JSON200 == nil {
			return nil, false, errors.New("received an empty response, JSON200 == nil")
		}
		switch project.JSON200.LifecycleState {
		case projects.ACTIVE:
			return project, true, nil
		case projects.CREATING:
			return project, false, nil
		}
		return project, false, fmt.Errorf("received project state '%s'. aborting", project.JSON200.LifecycleState)
	})
}

func (r DeleteResponse) WaitHandler(ctx context.Context, c *projects.ClientWithResponses, containerID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		project, err := c.GetWithResponse(ctx, containerID, &projects.GetParams{})
		if err != nil {
			return project, true, nil
		}
		if project.StatusCode() == http.StatusNotFound {
			return project, true, nil
		}
		return project, false, nil
	})
}
