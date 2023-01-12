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

const (
	connection_reset = "read: connection reset" // catch connection reset by peer error
)

// WaitHandler will wait for instance creation
// returned interface is nil or *instances.ProjectInstanceUI
func (r InstanceCreateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if strings.Contains(err.Error(), connection_reset) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if s.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if s.StatusCode() == http.StatusBadGateway {
			return nil, false, nil
		}
		if s.HasError != nil {
			return nil, false, s.HasError
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("received an empty response. JSON200 == nil")
		}
		if s.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED {
			return s.JSON200, true, nil
		}
		return s.JSON200, false, nil
	})
}

// WaitHandler will wait for instance update
// returned interface is nil or *instances.ProjectInstanceUI
func (r InstanceUpdateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	seenUpdating := false
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if strings.Contains(err.Error(), connection_reset) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if s.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if s.StatusCode() == http.StatusBadGateway {
			return nil, false, nil
		}
		if s.HasError != nil {
			return nil, false, s.HasError
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("received an empty response. JSON200 == nil")
		}
		if s.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_UPDATE_SUCCEEDED ||
			(seenUpdating && s.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED) {
			return s.JSON200, true, nil
		}
		if s.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_UPDATE_FAILED {
			return s.JSON200, true, fmt.Errorf("update failed for instance %s", instanceID)
		}
		if s.JSON200.Status != instances.PROJECT_INSTANCE_UI_STATUS_UPDATING {
			// in some cases it takes a long time for the server to change the
			// instance status to UPDATING
			// the following code will wait for the status change for 5 minutes
			// and continue the outer wait on change or fail
			w := wait.New(func() (res interface{}, done bool, err error) {
				si, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
				if err != nil {
					return nil, false, err
				}
				if si.HasError != nil {
					return nil, false, s.HasError
				}
				if si.JSON200 == nil {
					return nil, false, errors.New("received an empty response. JSON200 == nil")
				}
				if si.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_UPDATING ||
					si.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_UPDATE_SUCCEEDED ||
					si.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_UPDATE_FAILED {
					return nil, true, nil
				}
				return nil, false, nil
			})
			_, err := w.SetTimeout(5 * time.Minute).WaitWithContext(ctx)
			return nil, false, err
		}
		seenUpdating = true
		return s.JSON200, false, nil
	})
}

// WaitHandler will wait for instance deletion
// returned interface is nil
func (r InstanceDeleteResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.InstanceReadWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if strings.Contains(err.Error(), connection_reset) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if s.StatusCode() == http.StatusNotFound {
			return nil, true, nil
		}
		if s.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if s.StatusCode() == http.StatusBadGateway {
			return nil, false, nil
		}
		if s.HasError != nil {
			return nil, false, s.HasError
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("received an empty response. JSON200 == nil")
		}
		if s.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_DELETE_SUCCEEDED {
			return nil, true, nil
		}
		if s.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_DELETE_FAILED {
			return s.JSON200, true, fmt.Errorf("deletion failed for instance %s", instanceID)
		}
		if s.JSON200.Status != instances.PROJECT_INSTANCE_UI_STATUS_DELETING {
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
				if si.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_DELETING ||
					si.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_DELETE_FAILED ||
					si.JSON200.Status == instances.PROJECT_INSTANCE_UI_STATUS_DELETE_SUCCEEDED {
					return nil, true, nil
				}
				return nil, false, nil
			})
			_, err := w.SetTimeout(5 * time.Minute).WaitWithContext(ctx)
			return nil, false, err
		}
		return nil, false, nil
	})
}
