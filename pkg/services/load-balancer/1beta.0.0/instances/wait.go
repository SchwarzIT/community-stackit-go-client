package instances

import (
	"context"
	"errors"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// Wait will wait for instance create to complete
func (*CreateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, name string) *wait.Handler {
	maxFailCount := 5
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, name)
		if err = validate.Response(s, err, "JSON200.Status"); err != nil {
			return nil, false, err
		}
		status := *s.JSON200.Status
		if status == STATUS_READY {
			return s, true, nil
		}
		if status == STATUS_ERROR {
			if maxFailCount == 0 {
				return s, false, errors.New("received status FAILED from server")
			}
			maxFailCount--
		}
		if status == STATUS_TERMINATING {
			return s, false, errors.New("received status TERMINATING from server")
		}
		return s, false, nil
	})
}

// Wait will wait for instance deletion
// returned value for deletion wait will always be nil
func (DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, name string) *wait.Handler {
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
