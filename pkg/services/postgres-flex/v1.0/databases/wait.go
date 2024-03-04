package databases

import (
	"context"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
	"net/http"
)

func (PostInstanceDatabasesResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		s, err := c.GetDatabases(ctx, projectID, instanceID)
		if agg := validate.Response(s, err, "JSON200.Items"); agg != nil {
			return nil, false, agg
		}
		for _, db := range *s.JSON200.Databases {
			if db.ID == nil || *db.ID != instanceID {
				continue
			}
			// database was found
			return nil, false, nil
		}
		return nil, true, nil
	})
}

// WaitHandler will wait for database deletion
// returned value for deletion wait will always be nil
func (DeleteDatabasesDatabaseIDResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID, databaseID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		res, err := c.GetDatabases(ctx, projectID, instanceID)
		if err = validate.Response(res, err); err != nil {
			if res != nil && res.StatusCode() == http.StatusNotFound {
				return nil, true, nil
			}
			return nil, false, err
		}
		for _, db := range *res.JSON200.Databases {
			if db.ID == nil || *db.ID != databaseID {
				continue
			}
			// database was found
			return nil, false, nil
		}
		return nil, true, nil
	})
}
