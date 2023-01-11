package bucket

import (
	"context"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// WaitHandler for creation. in case there are no errors, the returned interface is of *GetResponse
func (svc *CreateResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID string, bucketName string) wait.WaitFn {
	return func() (interface{}, bool, error) {
		res, err := c.GetWithResponse(ctx, projectID, bucketName)
		if err != nil {
			return nil, false, err
		}
		if res.StatusCode() == http.StatusInternalServerError {
			return res, false, nil
		}
		if res.StatusCode() == http.StatusNotFound {
			return res, false, nil
		}
		return res, true, nil
	}
}

// WaitHandler for deletion
func (svc *DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID string, bucketName string) wait.WaitFn {
	return func() (interface{}, bool, error) {
		res, err := c.GetWithResponse(ctx, projectID, bucketName)
		if err != nil {
			return nil, false, err
		}
		if res.StatusCode() == http.StatusNotFound {
			return nil, true, nil
		}
		return nil, false, nil
	}
}
