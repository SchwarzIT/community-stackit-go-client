package instance

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

const ClientTimeoutErr = "Client.Timeout exceeded while awaiting headers"

// WaitHandler will wait for instance creation to complete
// returned interface is of *InstanceSingleInstance
func (r CreateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return createOrUpdateWait(ctx, c, projectID, instanceID)
}

// WaitHandler will wait for instance update to complete
// returned interface is of *InstanceSingleInstance
func (r PutResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	// artifical wait for instance to change from status ready to updating
	time.Sleep(5 * time.Second)
	return createOrUpdateWait(ctx, c, projectID, instanceID)
}

func createOrUpdateWait(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (res interface{}, done bool, err error) {
		s, err := c.GetWithResponse(ctx, projectID, instanceID)
		if err != nil {
			if strings.Contains(err.Error(), ClientTimeoutErr) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if s.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if s.HasError != nil {
			return nil, false, err
		}
		if s.JSON200 == nil {
			return nil, false, errors.New("bad response")
		}
		if s.JSON200.Item == nil || *s.JSON200.Item.Status == "" || *s.JSON200.Item.Status == consts.MONGO_DB_STATUS_READY {
			return s.JSON200.Item, true, nil
		}
		if *s.JSON200.Item.Status == consts.MONGO_DB_STATUS_FAILED {
			return s.JSON200.Item, false, errors.New("received status FAILED from server")
		}
		return s.JSON200.Item, false, nil
	})
}

// WaitHandler will wait for instance deletion
// returned value for deletion wait will always be nil
func (r DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.ListWithResponse(ctx, projectID, &ListParams{})
		if err != nil {
			if strings.Contains(err.Error(), ClientTimeoutErr) {
				return nil, false, nil
			}
			return nil, false, err
		}
		if res.StatusCode() == http.StatusInternalServerError {
			return nil, false, nil
		}
		if res.HasError != nil {
			return nil, false, res.HasError
		}
		if res.JSON200 == nil || res.JSON200.Items == nil {
			return nil, false, errors.New("received an empty response for list")
		}
		for _, v := range *res.JSON200.Items {
			if v.ID == nil {
				continue
			}
			if *v.ID == instanceID {
				return nil, false, nil
			}
		}
		return nil, true, nil
	})
}
