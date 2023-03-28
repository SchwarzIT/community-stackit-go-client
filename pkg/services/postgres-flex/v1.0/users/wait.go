package users

import (
	"context"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

const ClientTimeoutErr = "Client.Timeout exceeded while awaiting headers"

// Wait will wait for user deletion
// returned value for deletion wait will always be nil
func (*DeleteResponse) WaitHandler(ctx context.Context, c *ClientWithResponses, projectID, instanceID, userID string) *wait.Handler {
	return wait.New(func() (interface{}, bool, error) {
		s, err := c.List(ctx, projectID, instanceID)
		if agg := validate.Response(s, err, "JSON200.Items"); agg != nil {
			if strings.Contains(agg.Error(), ClientTimeoutErr) ||
				validate.StatusEquals(s,
					http.StatusBadGateway,
					http.StatusGatewayTimeout,
					http.StatusInternalServerError,
				) {
				return nil, false, nil
			}
			return nil, false, agg
		}
		for _, v := range *s.JSON200.Items {
			if v.ID == nil || *v.ID != userID {
				continue
			}
			// user was found
			return nil, false, nil
		}
		return nil, true, nil
	})
}
