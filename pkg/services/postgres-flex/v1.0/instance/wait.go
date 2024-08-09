package instance

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// Wait will wait for instance create to complete
func (*CreateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return waitForCreateOrUpdate(ctx, c, projectID, instanceID)
}

// Wait will wait for instance update to complete
func (*PutResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return waitForCreateOrUpdate(ctx, c, projectID, instanceID)
}

// Wait will wait for instance update to complete
func (*PatchResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return waitForCreateOrUpdate(ctx, c, projectID, instanceID)
}

// returned interface is of *InstanceSingleInstance
func waitForCreateOrUpdate(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	// artifical wait for instance to change from status ready to updating
	time.Sleep(5 * time.Second)
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if err = validate.Response(s, err, "JSON200.Item"); err != nil {
			return nil, false, err
		}
		if *s.JSON200.Item.Status == STATUS_READY {
			return s.JSON200.Item, true, nil
		}
		if *s.JSON200.Item.Status == STATUS_FAILED {
			return s.JSON200.Item, false, errors.New("received status FAILED from server")
		}
		return s.JSON200.Item, false, nil
	})
}

// Wait will wait for instance deletion
// returned value for deletion wait will always be nil
func (DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.Get(ctx, projectID, instanceID)
		if err = validate.Response(res, err); err != nil {
			// old handling of not found
			if res != nil && res.StatusCode() == http.StatusNotFound {
				return nil, true, nil
			}

			return nil, false, err
		}

		// new soft-deletion
		if res.JSON200 != nil && res.JSON200.Item.Status != nil && strings.ToUpper(*res.JSON200.Item.Status) == "DELETED" {
			return nil, true, nil
		}

		return nil, false, nil
	})
}
