// package traces is used to get & configure trances for an Argus instance

package traces

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// constants
const (
	apiPath = consts.API_PATH_ARGUS_TRACES
)

// New returns a new handler for the service
func New(c common.Client) *TracesService {
	return &TracesService{
		Client: c,
	}
}

// TracesService is the service that handles
// CRUD functionality for Argus traces
type TracesService common.Service

type Config struct {
	Retention string `json:"retention,omitempty"`
}

type GetConfigResponse struct {
	Message string `json:"message,omitempty"`
	Config  Config `json:"config,omitempty"`
}

type UpdateConfigResponse struct {
	Message string `json:"message,omitempty"`
}

// GetConfig returns argus grafana config
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#tag/grafana-configs
func (svc *TracesService) GetConfig(ctx context.Context, projectID, instanceID string) (res GetConfigResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// UpdateConfig updates argus grafana config
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_traces-configs_update
func (svc *TracesService) UpdateConfig(ctx context.Context, projectID, instanceID string, config Config) (res UpdateConfigResponse, err error) {
	if _, err = validate.Duration(config.Retention); err != nil {
		return
	}
	body, _ := json.Marshal(config)
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPath, projectID, instanceID), body)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}
