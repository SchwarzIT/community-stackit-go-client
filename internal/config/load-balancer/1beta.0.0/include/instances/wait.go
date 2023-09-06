package instances

import (
	"context"
	"errors"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/load-balancer/1beta.0.0/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// Wait will wait for instance create to complete
func (*CreateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, name string) *wait.Handler {
	return waitForCreate(ctx, c, projectID, name)
}

func waitForCreate(ctx context.Context, c *instances.ClientWithResponses, projectID, name string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, name)
		if err = validate.Response(s, err, "JSON200"); err != nil {
			return nil, false, err
		}
		if *s.JSON200.Status == instances.STATUS_READY {
			return s.JSON200, true, nil
		}
		if *s.JSON200.Status == instances.STATUS_ERROR {
			return s.JSON200, false, errors.New("received status FAILED from server")
		}
		return s.JSON200, false, nil
	})
}

// Wait will wait for instance deletion
// returned value for deletion wait will always be nil
func (DeleteResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, name string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.Get(ctx, projectID, name)
		if err = validate.Response(res, err); err != nil {
			if res != nil && res.StatusCode() == http.StatusNotFound {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	})
}
