// package projects enables/disables SKE usage in STACKIT projects
// IMPORTANT: disabling SKE will cause existing clusters to be automatically deleted

package projects

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPath = consts.API_PATH_SKE_PROJECTS
)

// New returns a new handler for the service
func New(c common.Client) *KubernetesProjectsService {
	return &KubernetesProjectsService{
		Client: c,
	}
}

// KubernetesProjectsService is the service that handles
// enabling / disabling SKE for a project
type KubernetesProjectsService common.Service

// KubernetesProjectsResponse is the project ID and state response
type KubernetesProjectsResponse struct {
	ProjectID string `json:"projectId"`
	State     string `json:"state"`
}

// Get returns the SKE project status
// See also https://api.stackit.schwarz/ske-service/openapi.v1.html#operation/SkeService_GetProject
func (svc *KubernetesProjectsService) Get(ctx context.Context, projectID string) (res KubernetesProjectsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// Create creates a SKE project
// See also https://api.stackit.schwarz/ske-service/openapi.v1.html#operation/SkeService_GetProject
func (svc *KubernetesProjectsService) Create(ctx context.Context, projectID string) (res KubernetesProjectsResponse, w *wait.Handler, err error) {
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	w = wait.New(svc.waitForCreation(ctx, projectID))
	return
}

// waitForCreation returns a wait function to determine if kubernetes project has been created
func (svc *KubernetesProjectsService) waitForCreation(ctx context.Context, projectID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		res, err = svc.Get(ctx, projectID)
		if err != nil {
			if strings.Contains(err.Error(), "project has no assigned namespace") {
				return nil, false, nil
			}
			return nil, false, err
		}
		project := res.(KubernetesProjectsResponse)
		switch project.State {
		case consts.SKE_PROJECT_STATUS_FAILED:
			fallthrough
		case consts.SKE_PROJECT_STATUS_DELETING:
			err = fmt.Errorf("received state: %s for project ID: %s",
				project.State,
				project.ProjectID,
			)
			return
		case consts.SKE_PROJECT_STATUS_CREATED:
			return nil, true, nil
		}
		return nil, false, nil
	}
}

// Delete deletes a SKE project
// IMPORTANT: existing clusters to be automatically deleted
// See also https://api.stackit.schwarz/ske-service/openapi.v1.html#operation/SkeService_GetProject
func (svc *KubernetesProjectsService) Delete(ctx context.Context, projectID string) (w *wait.Handler, err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, nil)
	w = wait.New(svc.waitForDeletion(ctx, projectID))
	return
}

// waitForDeletion returns a wait function to determine if kubernetes project has been deleted
func (svc *KubernetesProjectsService) waitForDeletion(ctx context.Context, projectID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		res, err = svc.Get(ctx, projectID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	}
}
