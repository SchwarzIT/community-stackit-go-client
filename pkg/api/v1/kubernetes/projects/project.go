// package projects enables/disables SKE usage in STACKIT projects
// IMPORTANT: disabling SKE will cause existing clusters to be automatically deleted

package projects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
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

// Create creates an SKE project
// See also https://api.stackit.schwarz/ske-service/openapi.v1.html#operation/SkeService_GetProject
func (svc *KubernetesProjectsService) Create(ctx context.Context, projectID string) (res KubernetesProjectsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// Delete deletes an SKE project
// IMPORTANT: existing clusters to be automatically deleted
// See also https://api.stackit.schwarz/ske-service/openapi.v1.html#operation/SkeService_GetProject
func (svc *KubernetesProjectsService) Delete(ctx context.Context, projectID string) (err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, nil)
	return
}
