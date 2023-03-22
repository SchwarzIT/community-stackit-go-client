package instances

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

const (
	connection_reset = "read: connection reset" // catch connection reset by peer error
	gateway_timeout  = "Gateway Timeout"
)

// WaitHandler will wait for  creation
// returned interface is nil or *ProjectInstanceUI
func (r CreateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if agg := validate.Response(s, err, "JSON200"); agg != nil {
			if strings.Contains(agg.Error(), connection_reset) ||
				strings.Contains(agg.Error(), gateway_timeout) {
				return nil, false, nil
			}
			if validate.StatusEquals(s,
				http.StatusBadGateway,
				http.StatusGatewayTimeout,
				http.StatusInternalServerError,
				http.StatusForbidden,
			) {
				return nil, false, nil
			}
			return nil, false, agg
		}
		if s.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED {
			return s.JSON200, true, nil
		}
		return s.JSON200, false, nil
	})
}

// WaitHandler will wait for  update
// returned interface is nil or *ProjectInstanceUI
func (r UpdateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	seenUpdating := false
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if agg := validate.Response(s, err, "JSON200"); agg != nil {
			if strings.Contains(agg.Error(), connection_reset) ||
				strings.Contains(agg.Error(), gateway_timeout) {
				return nil, false, nil
			}
			if validate.StatusEquals(s,
				http.StatusBadGateway,
				http.StatusGatewayTimeout,
				http.StatusInternalServerError,
				http.StatusForbidden,
			) {
				return nil, false, nil
			}
			return nil, false, agg
		}

		if s.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_UPDATE_SUCCEEDED ||
			(seenUpdating && s.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED) {
			return s.JSON200, true, nil
		}
		if s.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_UPDATE_FAILED {
			return s.JSON200, true, fmt.Errorf("update failed for  %s", instanceID)
		}
		if s.JSON200.Status != PROJECT_INSTANCE_UI_STATUS_UPDATING {
			// in some cases it takes a long time for the server to change the
			// instance status to UPDATING
			// the following code will wait for the status change for 5 minutes
			// and continue the outer wait on change or fail
			w := wait.New(func() (res interface{}, done bool, err error) {
				si, err := c.Get(ctx, projectID, instanceID)
				if agg := validate.Response(si, err, "JSON200"); agg != nil {
					if strings.Contains(agg.Error(), connection_reset) ||
						strings.Contains(agg.Error(), gateway_timeout) {
						return nil, false, nil
					}
					if validate.StatusEquals(s,
						http.StatusBadGateway,
						http.StatusGatewayTimeout,
						http.StatusInternalServerError,
						http.StatusForbidden,
					) {
						return nil, false, nil
					}
					return nil, false, agg
				}
				if si.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_UPDATING ||
					si.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_UPDATE_SUCCEEDED ||
					si.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_UPDATE_FAILED {
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

// WaitHandler will wait for  deletion
// returned interface is nil
func (r DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if agg := validate.Response(s, err, "JSON200"); agg != nil {
			if strings.Contains(agg.Error(), connection_reset) ||
				strings.Contains(agg.Error(), gateway_timeout) {
				return nil, false, nil
			}
			if validate.StatusEquals(s,
				http.StatusBadGateway,
				http.StatusGatewayTimeout,
				http.StatusInternalServerError,
				http.StatusForbidden,
			) {
				return nil, false, nil
			}
			if validate.StatusEquals(s, http.StatusNotFound) {
				return nil, true, nil
			}
			return nil, false, agg
		}
		if s.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_DELETE_SUCCEEDED {
			return nil, true, nil
		}
		if s.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_DELETE_FAILED {
			return s.JSON200, true, fmt.Errorf("deletion failed for  %s", instanceID)
		}
		if s.JSON200.Status != PROJECT_INSTANCE_UI_STATUS_DELETING {
			// in some cases it takes a long time for the server to change the
			// instance status to status DELETING
			// the following code will wait for the status change for 5 minutes
			// and continue the outer wait on change or fail
			w := wait.New(func() (res interface{}, done bool, err error) {
				si, err := c.Get(ctx, projectID, instanceID)
				if agg := validate.Response(si, err, "JSON200"); agg != nil {
					if strings.Contains(agg.Error(), connection_reset) ||
						strings.Contains(agg.Error(), gateway_timeout) {
						return nil, false, nil
					}
					if validate.StatusEquals(s,
						http.StatusBadGateway,
						http.StatusGatewayTimeout,
						http.StatusInternalServerError,
						http.StatusForbidden,
					) {
						return nil, false, nil
					}
					if validate.StatusEquals(s, http.StatusNotFound) {
						return nil, true, nil
					}
					return nil, false, agg
				}
				if si.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_DELETING ||
					si.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_DELETE_FAILED ||
					si.JSON200.Status == PROJECT_INSTANCE_UI_STATUS_DELETE_SUCCEEDED {
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
