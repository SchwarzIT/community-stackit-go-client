package resourcemanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

func (*CreateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, containerID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		project, err := c.Get(ctx, containerID, &GetParams{})
		if err != nil {
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "403") {
				return project, false, nil
			}
			return project, false, err
		}
		if project.StatusCode() == http.StatusForbidden {
			return project, false, nil
		}
		if project.JSON200 == nil {
			return nil, false, errors.New("received an empty response, JSON200 == nil")
		}
		switch project.JSON200.LifecycleState {
		case ACTIVE:
			return project, true, nil
		case CREATING:
			return project, false, nil
		}
		return project, false, fmt.Errorf("received project state '%s'. aborting", project.JSON200.LifecycleState)
	})
}

func (*DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, containerID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		project, err := c.Get(ctx, containerID, &GetParams{})
		if err != nil {
			return project, true, nil
		}
		if project.StatusCode() == http.StatusNotFound {
			return project, true, nil
		}
		return project, false, nil
	})
}
