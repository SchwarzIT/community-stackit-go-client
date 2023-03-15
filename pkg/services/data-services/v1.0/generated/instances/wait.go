package instances

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

const (
	client_timeout_err = "Client.Timeout exceeded while awaiting headers"
)

func (r ProvisionResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if !strings.Contains(err.Error(), client_timeout_err) {
				return s, false, err
			}
			return nil, false, nil
		}
		if s.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if s.Error != nil {
			return nil, false, err
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response, JSON200 is nil")
		}
		if s.JSON200.LastOperation.State == SUCCEEDED {
			return s, true, nil
		}
		if s.JSON200.LastOperation.State == FAILED {
			return s, false, errors.New("received failed status from DSA instance")
		}
		return s, false, nil
	})
}

func (r UpdateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if !strings.Contains(err.Error(), client_timeout_err) {
				return s, false, err
			}
			return nil, false, nil
		}
		if s.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if s.Error != nil {
			return nil, false, err
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response, JSON200 is nil")
		}
		if s.JSON200.LastOperation.Type != UPDATE {
			return s, false, nil
		}
		if s.JSON200.LastOperation.State == SUCCEEDED {
			return s, true, nil
		}
		if s.JSON200.LastOperation.State == FAILED {
			return s, false, errors.New("received failed status from DSA instance")
		}
		return s, false, nil
	})
}

func (r DeprovisionResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if !strings.Contains(err.Error(), client_timeout_err) {
				return s, false, err
			}
			return nil, false, nil
		}
		if s.StatusCode() == http.StatusNotFound || s.StatusCode() == http.StatusGone {
			return nil, true, nil
		}
		if s.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if s.Error != nil {
			return nil, false, s.Error
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response, JSON200 is nil")
		}
		if s.JSON200.LastOperation.Type != DELETE {
			return s, false, nil
		}
		if s.JSON200.LastOperation.State == SUCCEEDED {
			return s, true, nil
		}
		return s, false, nil
	})
}
