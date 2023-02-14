package instances

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0.1/generated/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// WaitHandler will wait for instance creation to complete
// returned interface is of *instance.InstanceSingleInstance
func (r CreateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return createOrUpdateWait(ctx, c, projectID, instanceID, instances.CREATE)
}

// WaitHandler will wait for instance update to complete
// returned interface is of *instance.InstanceSingleInstance
func (r UpdateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	// artifical wait for instance to change from status ready to updating
	time.Sleep(5 * time.Second)
	return createOrUpdateWait(ctx, c, projectID, instanceID, instances.UPDATE)
}

func createOrUpdateWait(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string, opType instances.LastOperationType) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.HasError != nil {
			return nil, false, err
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response")
		}
		state := s.JSON200.LastOperation.State
		if opType != s.JSON200.LastOperation.Type {
			return *s.JSON200, false, nil
		}
		if state == instances.SUCCEEDED {
			return *s.JSON200, true, nil
		}
		if state == instances.FAILED {
			return *s.JSON200, false, errors.New("received status FAILED from server")
		}
		return *s.JSON200, false, nil
	})
}

// WaitHandler will wait for instance deletion
// returned value for deletion wait will always be nil
func (r DeleteResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if res.StatusCode() == http.StatusNotFound {
			return nil, true, nil
		}
		if res.HasError != nil {
			return nil, false, res.HasError
		}
		return nil, false, nil
	})
}
