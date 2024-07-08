package serviceenablement

import (
	"context"
	"fmt"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
	"github.com/pkg/errors"
)

func (*GetServiceResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, serviceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {

		resp, err := c.GetService(ctx, projectID, serviceID)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed during create request preparation")
		}
		if err != nil {
			return nil, false, err
		}

		switch *resp.JSON200.State {
		case DISABLING:
			return nil, false, fmt.Errorf("received state: %s for project ID: %s and service ID: %s",
				*resp.JSON200.State,
				projectID,
				*resp.JSON200.ServiceID,
			)
		case ENABLED:
			return nil, true, nil
		}

		return nil, false, nil
	})
}
