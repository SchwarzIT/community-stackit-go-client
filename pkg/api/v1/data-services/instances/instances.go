// package instances is used to manange DSA instances

package instances

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPathList   = consts.API_PATH_DSA_INSTANCES
	apiPathCreate = consts.API_PATH_DSA_INSTANCES
	apiPathGet    = consts.API_PATH_DSA_INSTANCE
	apiPathUpdate = consts.API_PATH_DSA_INSTANCE
)

// New returns a new handler for the service
func New(c common.Client) *DSAInstancesService {
	return &DSAInstancesService{
		Client: c,
	}
}

// DSAInstancesService is the service that manages DSA instances
type DSAInstancesService common.Service

// ListResponse represents a list of instances returned from the server
type ListResponse struct {
	Instances []Instance `json:"instances,omitempty"`
}

// Instance is a struct representing an instance
type Instance struct {
	InstanceID         string            `json:"instanceId,omitempty"`
	Name               string            `json:"name,omitempty"`
	PlanID             string            `json:"planId,omitempty"`
	DashboardURL       string            `json:"dashboardUrl,omitempty"`
	CFGUID             string            `json:"cfGuid,omitempty"`
	CFSpaceGUID        string            `json:"cfSpaceGuid,omitempty"`
	CFOrganizationGUID string            `json:"cfOrganizationGuid,omitempty"`
	ImageURL           string            `json:"imageUrl,omitempty"`
	Parameters         map[string]string `json:"parameters,omitempty"`
	LastOperation      LastOperation     `json:"lastOperation,omitempty"`
}

// LastOperation is a struct representing instance last operation
type LastOperation struct {
	Type        string `json:"type,omitempty"`
	State       string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetResponse is the server response for Get call
type GetResponse Instance

// CreateRequest holds data for creating instance
type CreateRequest struct {
	PlanID       string            `json:"planId,omitempty"`
	InstanceName string            `json:"instanceName,omitempty"`
	Parameters   map[string]string `json:"parameters,omitempty"`
}

// CreateResponse is the server response for a creation call
type CreateResponse struct {
	InstanceID string `json:"instanceId,omitempty"`
}

// UpdateRequest holds data for updating instance
type UpdateRequest struct {
	PlanID     string            `json:"planId,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// UpdateResponse is the server response for an Update call
type UpdateResponse struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns a list of DSA instances in project
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Instance.list
func (svc *DSAInstancesService) List(ctx context.Context, projectID string) (res ListResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathList, projectID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Get returns the instance information by project and instance IDs
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Instance.get
func (svc *DSAInstancesService) Get(ctx context.Context, projectID, instanceID string) (res GetResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathGet, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Create creates a new DSA instance and returns the server response (CreateResponse) and a wait handler
// which upon call to `Wait()` will wait until the instance is successfully created
// Wait() returns the full instance details (Instance) and error if it occurred
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Instance.provision
func (svc *DSAInstancesService) Create(ctx context.Context, projectID, instanceName, planID string, parameters map[string]string) (res CreateResponse, w *wait.Handler, err error) {

	// build body
	data, _ := svc.buildCreateRequest(instanceName, planID, parameters)

	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPathCreate, projectID), data)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, &res)

	// create Wait service
	w = wait.New(svc.waitForCreation(ctx, projectID, res.InstanceID))
	w.SetTimeout(1 * time.Hour)
	return res, w, err
}

func (svc *DSAInstancesService) buildCreateRequest(instanceName, planID string, parameters map[string]string) ([]byte, error) {
	return json.Marshal(CreateRequest{
		InstanceName: instanceName,
		PlanID:       planID,
		Parameters:   parameters,
	})
}

func (svc *DSAInstancesService) waitForCreation(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.LastOperation.State == consts.DSA_STATE_SUCCEEDED {
			return s, true, nil
		}
		if s.LastOperation.State == consts.DSA_STATE_FAILED {
			return s, false, errors.New("received failed status from DSA instance")
		}
		return s, false, nil
	}
}

// Update updates a DSA instance and returns the server response (UpdateResponse) and a wait handler
// which upon call to `Wait()` will wait until the instance is successfully updated
// Wait() returns the full instance details (Instance) and error if it occurred
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Instance.update
func (svc *DSAInstancesService) Update(ctx context.Context, projectID, instanceID, planID string, parameters map[string]string) (res UpdateResponse, w *wait.Handler, err error) {

	// build body
	data, _ := svc.buildUpdateRequest(planID, parameters)

	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPathUpdate, projectID, instanceID), data)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, &res)

	// create Wait service
	w = wait.New(svc.waitForUpdate(ctx, projectID, instanceID))
	w.SetTimeout(1 * time.Hour)
	return res, w, err
}

func (svc *DSAInstancesService) buildUpdateRequest(planID string, parameters map[string]string) ([]byte, error) {
	return json.Marshal(UpdateRequest{
		PlanID:     planID,
		Parameters: parameters,
	})
}

func (svc *DSAInstancesService) waitForUpdate(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.LastOperation.Type != consts.DSA_OPERATION_TYPE_UPDATE {
			return s, false, nil
		}
		if s.LastOperation.State == consts.DSA_STATE_SUCCEEDED {
			return s, true, nil
		}
		if s.LastOperation.State == consts.DSA_STATE_FAILED {
			return s, false, errors.New("received failed status from DSA instance")
		}
		return s, false, nil
	}
}

// Delete deletes a DSA instance and returns a wait handler and error if occurred
// `Wait()` will wait until the instance is successfully deleted
// Wait() returns nil (empty response from server) and error if it occurred
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Instance.deprovision
func (svc *DSAInstancesService) Delete(ctx context.Context, projectID, instanceID string) (w *wait.Handler, err error) {
	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPathUpdate, projectID, instanceID), nil)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, nil)

	// create Wait service
	w = wait.New(svc.waitForDeletion(ctx, projectID, instanceID))
	w.SetTimeout(1 * time.Hour)
	return w, err
}

func (svc *DSAInstancesService) waitForDeletion(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) ||
				strings.Contains(err.Error(), http.StatusText(http.StatusGone)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		if s.LastOperation.Type != consts.DSA_OPERATION_TYPE_DELETE {
			return s, false, nil
		}
		if s.LastOperation.State == consts.DSA_STATE_SUCCEEDED {
			return s, true, nil
		}
		if s.LastOperation.State == consts.DSA_STATE_FAILED {
			return s, false, errors.New("received failed status for DSA instance deletion")
		}
		return nil, false, nil
	}
}
