package bucket

import (
	"context"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// WaitForeCreate waits for creation. in case there are no errors, the returned interface is of *GetResponse
func (c *ClientWithResponses[K]) WaitForeCreate(ctx context.Context, projectID, bucketName string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.Get(ctx, projectID, bucketName)
		if err = validate.Response(res, err); err != nil {
			if validate.StatusEquals(res,
				http.StatusBadGateway,
				http.StatusGatewayTimeout,
				http.StatusInternalServerError,
				http.StatusNotFound,
			) {
				return nil, false, nil
			}
			return nil, false, err
		}
		return res, true, nil
	})
}

// WaitForeDelete waits for deletion
func (c *ClientWithResponses[K]) WaitForeDelete(ctx context.Context, projectID, bucketName string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.Get(ctx, projectID, bucketName)
		if err = validate.Response(res, err); err != nil {
			if validate.StatusEquals(res, http.StatusNotFound) {
				return nil, true, nil
			}
			if validate.StatusEquals(res,
				http.StatusBadGateway,
				http.StatusGatewayTimeout,
				http.StatusInternalServerError,
			) {
				return nil, false, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	})
}
