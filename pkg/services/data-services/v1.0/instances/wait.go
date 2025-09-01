package instances

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

const (
	client_timeout_err = "Client.Timeout exceeded while awaiting headers"
)

func (ProvisionResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if err = validate.Response(s, err, "JSON200"); err != nil {
			if strings.Contains(err.Error(), client_timeout_err) ||
				validate.StatusEquals(s,
					http.StatusBadGateway,
					http.StatusGatewayTimeout,
					http.StatusInternalServerError,
				) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if s.JSON200.LastOperation.State == SUCCEEDED {
			return s, true, nil
		}
		if s.JSON200.LastOperation.State == FAILED {
			return nil, false, errors.New("received failed status from DSA instance")
		}
		return nil, false, nil
	})
}

func (UpdateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if err = validate.Response(s, err, "JSON200"); err != nil {
			if strings.Contains(err.Error(), client_timeout_err) ||
				validate.StatusEquals(s,
					http.StatusBadGateway,
					http.StatusGatewayTimeout,
					http.StatusInternalServerError,
				) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if s.JSON200.LastOperation.Type == UPDATE {
			return nil, false, nil
		}
		if s.JSON200.LastOperation.State == SUCCEEDED {
			return s, true, nil
		}
		if s.JSON200.LastOperation.State == FAILED {
			return nil, false, errors.New("received failed status from DSA instance")
		}

		return nil, false, fmt.Errorf("received unexpected status from DSA instance: %s", s.JSON200.LastOperation.State)
	})
}

func (DeprovisionResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.Get(ctx, projectID, instanceID)
		if err = validate.Response(s, err, "JSON200"); err != nil {
			if strings.Contains(err.Error(), client_timeout_err) ||
				validate.StatusEquals(s,
					http.StatusBadGateway,
					http.StatusGatewayTimeout,
					http.StatusInternalServerError,
				) {
				return nil, false, nil
			}
			if validate.StatusEquals(s, http.StatusNotFound, http.StatusGone, http.StatusForbidden) {
				return nil, true, nil
			}
			return nil, false, err
		}
		if s.JSON200.LastOperation.Type != DELETE {
			return nil, false, nil
		}
		if s.JSON200.LastOperation.State == SUCCEEDED {
			return s, true, nil
		}
		return nil, false, nil
	})
}
