package bucket

import (
	"context"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/bucket"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// Wait waits for creation. in case there are no errors, the returned interface is of *GetResponse
func (*CreateResponse) WaitHandler(ctx context.Context, c *bucket.ClientWithResponses, projectID, bucketName string) *wait.Handler {
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

// Wait waits for deletion
func (*DeleteResponse) WaitHandler(ctx context.Context, c *bucket.ClientWithResponses, projectID, bucketName string) *wait.Handler {
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
