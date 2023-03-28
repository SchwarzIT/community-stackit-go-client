package instance

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

const ClientTimeoutErr = "Client.Timeout exceeded while awaiting headers"

// WaitHandler will wait for instance creation to complete
// returned interface is of *InstanceSingleInstance
func (r CreateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return createOrUpdateWait(ctx, c, projectID, instanceID)
}

// WaitHandler will wait for instance update to complete
// returned interface is nil
func (r PutResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	// artificial wait for instance to change from status ready to updating
	time.Sleep(5 * time.Second)
	return createOrUpdateWait(ctx, c, projectID, instanceID)
}

// WaitHandler will wait for instance update to complete
// returned interface is nil
func (r PatchResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	// artificial wait for instance to change from status ready to updating
	time.Sleep(5 * time.Second)
	return createOrUpdateWait(ctx, c, projectID, instanceID)
}

// Wait will wait for instance update to complete
// returned interface is nil
func createOrUpdateWait(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	// artificial wait for instance to change status
	time.Sleep(5 * time.Second)

	outerfound := false
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.List(ctx, projectID, &ListParams{})
		if err = validate.Response(s, err, "JSON200.Items"); err != nil {
			if strings.Contains(err.Error(), ClientTimeoutErr) ||
				validate.StatusEquals(s,
					http.StatusBadGateway,
					http.StatusGatewayTimeout,
					http.StatusInternalServerError,
				) {
				return nil, false, nil
			}
			return nil, false, err
		}

		innerfound := false
		for _, item := range *s.JSON200.Items {
			if item.ID == nil || item.Status == nil || *item.ID != instanceID {
				continue
			}
			outerfound = true
			innerfound = true
			if *item.Status == STATUS_READY {
				return nil, true, nil
			}
		}
		if !innerfound && outerfound {
			return nil, false, fmt.Errorf("instance %s is not in the project's instance list", instanceID)
		}
		return nil, false, nil
	})
}

// WaitHandler will wait for instance deletion
// returned value for deletion wait will always be nil
func (DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		s, err := c.List(ctx, projectID, &ListParams{})
		if err = validate.Response(s, err, "JSON200.Items"); err != nil {
			if strings.Contains(err.Error(), ClientTimeoutErr) ||
				validate.StatusEquals(s,
					http.StatusBadGateway,
					http.StatusGatewayTimeout,
					http.StatusInternalServerError,
				) {
				return nil, false, nil
			}
			return nil, false, err
		}
		for _, v := range *s.JSON200.Items {
			if v.ID == nil || *v.ID != instanceID {
				continue
			}
			// instance was found
			return nil, false, nil
		}
		return nil, true, nil
	})
}
