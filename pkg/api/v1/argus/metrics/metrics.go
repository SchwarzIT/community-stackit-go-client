// package metrics is used to get & update metrics configuration, specifically for retention times

package metrics

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
	apiPath = consts.API_PATH_ARGUS_METRICS_RETENTION
)

// New returns a new handler for the service
func New(c common.Client) *MetricsService {
	return &MetricsService{
		Client: c,
	}
}

// MetricsService is the service that handles
// CRUD functionality for Argus metrics storage
type MetricsService common.Service

type GetConfigResponse struct {
	Message                 string `json:"message,omitempty"`
	MetricsRetentionTimeRaw string `json:"metricsRetentionTimeRaw,omitempty"`
	MetricsRetentionTime5m  string `json:"metricsRetentionTime5m,omitempty"`
	MetricsRetentionTime1h  string `json:"metricsRetentionTime1h,omitempty"`
}

type Config struct {
	MetricsRetentionTimeRaw string `json:"metricsRetentionTimeRaw,omitempty"`
	MetricsRetentionTime5m  string `json:"metricsRetentionTime5m,omitempty"`
	MetricsRetentionTime1h  string `json:"metricsRetentionTime1h,omitempty"`
}

type UpdateConfigResponse struct {
	Message string `json:"message,omitempty"`
}

// GetConfig returns argus metrics storage config
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_metrics-storage-retentions_list
func (svc *MetricsService) GetConfig(ctx context.Context, projectID, instanceID string) (res GetConfigResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// UpdateConfig updates argus metrics storage config
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_metrics-storage-retentions_update
func (svc *MetricsService) UpdateConfig(ctx context.Context, projectID, instanceID string, cfg Config) (res UpdateConfigResponse, err error) {
	if err = cfg.Validate(); err != nil {
		err = validate.WrapError(err)
		return
	}
	body, _ := json.Marshal(cfg)
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPath, projectID, instanceID), body)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}
