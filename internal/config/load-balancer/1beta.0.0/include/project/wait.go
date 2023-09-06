package project

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/load-balancer/1beta.0.0/project"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

func (*EnableProjectResponse) WaitHandler(ctx context.Context, c *project.ClientWithResponses, projectID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.GetStatus(ctx, projectID)
		if err = validate.Response(resp, err, "JSON200.Status"); err != nil {
			return nil, false, err
		}
		switch *resp.JSON200.Status {
		case project.STATUS_FAILED:
			fallthrough
		case project.STATUS_DELETING:
			return nil, false, fmt.Errorf("received state: %s for project ID: %s",
				*resp.JSON200.Status,
				projectID,
			)
		case project.STATUS_UNSPECIFIED:
			// in some cases beta APIs do not return a status
			fallthrough
		case project.STATUS_READY:
			return nil, true, nil
		}
		return nil, false, nil
	})
}

func (*DisableProjectResponse) WaitHandler(ctx context.Context, c *project.ClientWithResponses, projectID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		resp, err := c.GetStatus(ctx, projectID)
		if err = validate.Response(resp, err, "JSON200.Status"); err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		if *resp.JSON200.Status == project.STATUS_DISABLED ||
			*resp.JSON200.Status == project.STATUS_UNSPECIFIED {
			return nil, true, nil
		}
		return nil, false, nil
	})
}
