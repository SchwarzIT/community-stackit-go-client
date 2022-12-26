package instances

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// WaitHandler will wait for instance creation
// returned interface is nil or *instances.ProjectInstanceUI
func (r InstanceCreateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.HasError != nil {
			return nil, false, s.HasError
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("received an empty response. JSON200 == nil")
		}
		if s.JSON200.Status == instances.ProjectInstanceUIStatusCREATE_SUCCEEDED {
			return s.JSON200, true, nil
		}
		return s.JSON200, false, nil
	})
}

// WaitHandler will wait for instance update
// returned interface is nil or *instances.ProjectInstanceUI
func (r InstanceUpdateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.HasError != nil {
			return nil, false, s.HasError
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("received an empty response. JSON200 == nil")
		}
		if s.JSON200.Status == instances.ProjectInstanceUIStatusUPDATE_SUCCEEDED {
			return s.JSON200, true, nil
		}
		if s.JSON200.Status == instances.ProjectInstanceUIStatusUPDATE_FAILED {
			return s.JSON200, true, fmt.Errorf("update failed for instance %s", instanceID)
		}
		if s.JSON200.Status == instances.ProjectInstanceUIStatusCREATE_SUCCEEDED {
			// in some cases it takes a long time for the server to change the
			// instance status to UPDATING
			// the following code will wait for the status change for 5 minutes
			// and continue the outer wait on change or fail
			w := wait.New(func() (res interface{}, done bool, err error) {
				si, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
				if err != nil {
					return nil, false, err
				}
				if s.HasError != nil {
					return nil, false, s.HasError
				}
				if s.JSON200 == nil {
					return nil, false, errors.New("received an empty response. JSON200 == nil")
				}
				if si.JSON200.Status == instances.ProjectInstanceUIStatusUPDATING ||
					si.JSON200.Status == instances.ProjectInstanceUIStatusUPDATE_SUCCEEDED ||
					si.JSON200.Status == instances.ProjectInstanceUIStatusUPDATE_FAILED {
					return nil, true, nil
				}
				return nil, false, nil
			})
			_, err := w.SetTimeout(5 * time.Minute).Wait()
			return nil, false, err
		}
		return s.JSON200, false, nil
	})
}

// WaitHandler will wait for instance deletion
// returned interface is nil
func (r InstanceDeleteResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		if s.StatusCode() == http.StatusNotFound {
			return nil, true, nil
		}
		if s.HasError != nil {
			return nil, false, s.HasError
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("received an empty response. JSON200 == nil")
		}
		if s.JSON200.Status == instances.ProjectInstanceUIStatusDELETE_SUCCEEDED {
			return nil, true, nil
		}
		if s.JSON200.Status == instances.ProjectInstanceUIStatusDELETE_FAILED {
			return s.JSON200, true, fmt.Errorf("deletion failed for instance %s", instanceID)
		}
		if s.JSON200.Status != instances.ProjectInstanceUIStatusDELETING {
			// in some cases it takes a long time for the server to change the
			// instance status to status DELETING
			// the following code will wait for the status change for 5 minutes
			// and continue the outer wait on change or fail
			w := wait.New(func() (res interface{}, done bool, err error) {
				si, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
				if err != nil {
					if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
						return nil, true, nil
					}
					return nil, false, err
				}
				if si.StatusCode() == http.StatusNotFound {
					return nil, true, nil
				}
				if si.HasError != nil {
					return nil, false, s.HasError
				}
				if si.JSON200 == nil {
					return nil, false, errors.New("received an empty response. JSON200 == nil")
				}
				if si.JSON200.Status == instances.ProjectInstanceUIStatusDELETING ||
					si.JSON200.Status == instances.ProjectInstanceUIStatusDELETE_FAILED ||
					si.JSON200.Status == instances.ProjectInstanceUIStatusDELETE_SUCCEEDED {
					return nil, true, nil
				}
				return nil, false, nil
			})
			_, err := w.SetTimeout(5 * time.Minute).Wait()
			return nil, false, err
		}
		return nil, false, nil
	})
}
