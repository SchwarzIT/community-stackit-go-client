// package instances is used to create and mange Argus instances

package instances

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/plans"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPath               = consts.API_PATH_ARGUS_INSTANCES
	apiPathWithInstanceID = consts.API_PATH_ARGUS_WITH_INSTANCE_ID
)

// InstancesService is the service that handles
// CRUD functionality for Argus instances and also wraps instance credentials service
type InstancesService common.Service

// InstanceList is the structure returned from the list api endpoint
type InstanceList struct {
	Message   string         `json:"message,omitempty"`
	Instances []InstanceItem `json:"instances"`
}

// InstanceItem is an item in the list of instances from the list api endpoint
type InstanceItem struct {
	ID          string `json:"id,omitempty"`
	PlanName    string `json:"planName,omitempty"`
	Instance    string `json:"instance,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
}

// CreateOrUpdateRequest is the structure needed for creating or updating an instance
type CreateOrUpdateRequest struct {
	Name       string            `json:"name"`
	PlanID     string            `json:"planId"`
	Parameters map[string]string `json:"parameter"`
}

// CreateOrUpdateResponse is the api response structure for instance creation/update
type CreateOrUpdateResponse struct {
	Message      string `json:"message,omitempty"`
	InstanceID   string `json:"instanceId,omitempty"`
	DashboardURL string `json:"dashboardUrl,omitempty"`
}

// Instance is the structure returned when reading a single instance
type Instance struct {
	Message      string                  `json:"message,omitempty"`
	Error        string                  `json:"error,omitempty"`
	DashboardURL string                  `json:"dashboardUrl,omitempty"`
	IsUpdatable  bool                    `json:"isUpdatable,omitempty"`
	Name         string                  `json:"name,omitempty"`
	Parameters   map[string]string       `json:"parameters,omitempty"`
	ID           string                  `json:"id,omitempty"`
	ServiceName  string                  `json:"serviceName,omitempty"`
	PlanID       string                  `json:"planId,omitempty"`
	PlanName     string                  `json:"planName,omitempty"`
	PlanSchema   string                  `json:"planSchema,omitempty"`
	Status       string                  `json:"status,omitempty"`
	Instance     InstanceSensitiveFields `json:"instance,omitempty"`
}

// InstanceSensitiveFields provides more elaborated information of the instance, including sensitive data
type InstanceSensitiveFields struct {
	Instance                string        `json:"instance,omitempty"`
	Cluster                 string        `json:"cluster,omitempty"`
	GrafanaURL              string        `json:"grafanaUrl,omitempty"`
	DashboardURL            string        `json:"dashboardUrl,omitempty"`
	GrafanaPlugins          []interface{} `json:"grafanaPlugins,omitempty"`
	Name                    string        `json:"name,omitempty"`
	GrafanaAdminPassword    string        `json:"grafanaAdminPassword,omitempty"`
	GrafanaAdminUser        string        `json:"grafanaAdminUser,omitempty"`
	MetricsRetentionTimeRaw int           `json:"metricsRetentionTimeRaw,omitempty"`
	MetricsRetentionTime5m  int           `json:"MetricsRetentionTime5m,omitempty"`
	MetricsRetentionTime1h  int           `json:"MetricsRetentionTime1h,omitempty"`
	MetricsURL              string        `json:"metricsUrl,omitempty"`
	PushMetricsURL          string        `json:"pushMetricsUrl,omitempty"`
	GrafanaPublicReadAccess bool          `json:"grafanaPublicReadAccess,omitempty"`
	TargetsURL              string        `json:"targetsUrl,omitempty"`
	AlertingURL             string        `json:"alertingUrl,omitempty"`
	Plan                    plans.Plan    `json:"plan,omitempty"`
	LogsURL                 string        `json:"logsUrl,omitempty"`
	LogsPushURL             string        `json:"logsPushUrl,omitempty"`
	JaegerTracesURL         string        `json:"jaegerTracesUrl,omitempty"`
	OtlpTracesURL           string        `json:"otlpTracesUrl,omitempty"`
	ZipkinSpansURL          string        `json:"zipkinSpansUrl,omitempty"`
	JaegerUIURL             string        `json:"jaegerUiUrl,omitempty"`
}

// List returns a list of argus instances in project
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_list
func (svc *InstancesService) List(ctx context.Context, projectID string) (res InstanceList, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Get returns the instance information by project and instance IDs
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_read
func (svc *InstancesService) Get(ctx context.Context, projectID, instanceID string) (res Instance, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathWithInstanceID, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

func (svc *InstancesService) buildRequest(name, planID string, params map[string]string) ([]byte, error) {
	return json.Marshal(CreateOrUpdateRequest{
		Name:       name,
		PlanID:     planID,
		Parameters: params,
	})
}

// Create creates a new Argus instance and returns the server response (CreateOrUpdateResponse) and a wait handler
// which upon call to `Wait()` will wait until the instance is successfully created
// Wait() returns the full instance details (Instance) and error if it occurred
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_create
func (svc *InstancesService) Create(ctx context.Context, projectID, instanceName, planID string, params map[string]string) (res CreateOrUpdateResponse, w *wait.Handler, err error) {
	if err = Validate(projectID, instanceName, planID); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildRequest(instanceName, planID, params)
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPath, projectID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	w = wait.New(svc.waitForCreation(ctx, projectID, res.InstanceID))

	return res, w, err
}

func (svc *InstancesService) waitForCreation(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.Status == consts.ARGUS_INSTANCE_STATUS_CREATE_SUCCEEDED {
			return s, true, nil
		}
		return s, false, nil
	}
}

// Update updates a new Argus instance
// returns API response [CreateOrUpdateResponse], wait handler and error
// The wait handler will wait for the instance status to be set to "UPDATE_SUCCEEDED"
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_update
func (svc *InstancesService) Update(ctx context.Context, projectID, instanceID, instanceName, planID string, params map[string]string) (res CreateOrUpdateResponse, w *wait.Handler, err error) {
	if err = Validate(projectID, instanceName, planID); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildRequest(instanceName, planID, params)
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPathWithInstanceID, projectID, instanceID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	w = wait.New(svc.waitForUpdate(ctx, projectID, instanceID))
	return
}

func (svc *InstancesService) waitForUpdate(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.Status == consts.ARGUS_INSTANCE_STATUS_UPDATE_SUCCEEDED {
			return s, true, nil
		}
		if s.Status == consts.ARGUS_INSTANCE_STATUS_UPDATE_FAILED {
			return s, true, fmt.Errorf("update failed for instance %s", instanceID)
		}
		return s, false, nil
	}
}

// Delete deleted an instance by project and instance IDs
// Delete returns the instance information (Instance), wait handler to wait for the full deletion, and an error
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_delete
func (svc *InstancesService) Delete(ctx context.Context, projectID, instanceID string) (res Instance, w *wait.Handler, err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPathWithInstanceID, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	w = wait.New(svc.waitForDeletion(ctx, projectID, instanceID))
	return
}

func (svc *InstancesService) waitForDeletion(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		if s.Status == consts.ARGUS_INSTANCE_STATUS_DELETE_SUCCEEDED {
			return nil, true, nil
		}
		if s.Status != consts.ARGUS_INSTANCE_STATUS_DELETING {
			return nil, false, fmt.Errorf("expected instance to be in state 'DELETING', but got %s instead", s.Status)
		}
		return nil, false, nil
	}
}
