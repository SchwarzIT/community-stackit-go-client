package instances

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/elasticsearch/v1.0/generated/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

const (
	client_timeout_err = "Client.Timeout exceeded while awaiting headers"
)

func (r ProvisionResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil && !strings.Contains(err.Error(), client_timeout_err) {
			return nil, false, err
		}
		if s.HasError != nil && !strings.Contains(err.Error(), client_timeout_err) {
			return nil, false, err
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response, JSON200 is nil")
		}
		if s.JSON200.LastOperation.State == instances.SUCCEEDED {
			return s, true, nil
		}
		if s.JSON200.LastOperation.State == instances.FAILED {
			return s, false, errors.New("received failed status from DSA instance")
		}
		return s, false, nil
	})
}

func (r UpdateResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil && !strings.Contains(err.Error(), client_timeout_err) {
			return nil, false, err
		}
		if s.HasError != nil && !strings.Contains(err.Error(), client_timeout_err) {
			return nil, false, err
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response, JSON200 is nil")
		}
		if s.JSON200.LastOperation.Type != instances.UPDATE {
			return s, false, nil
		}
		if s.JSON200.LastOperation.State == instances.SUCCEEDED {
			return s, true, nil
		}
		if s.JSON200.LastOperation.State == instances.FAILED {
			return s, false, errors.New("received failed status from DSA instance")
		}
		return s, false, nil
	})
}

func (r DeprovisionResponse) WaitHandler(ctx context.Context, c *instances.ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil && !strings.Contains(err.Error(), client_timeout_err) {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) ||
				strings.Contains(err.Error(), http.StatusText(http.StatusGone)) {
				return nil, true, nil
			}
			return s, false, err
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response, JSON200 is nil")
		}
		if s.JSON200.LastOperation.Type != instances.DELETE {
			return s, false, nil
		}
		if s.JSON200.LastOperation.State == instances.SUCCEEDED {
			return s, true, nil
		}
		if s.JSON200.LastOperation.State == instances.FAILED {
			return s, false, errors.New("received failed status for DSA instance deletion")
		}
		return s, false, nil
	})
}
