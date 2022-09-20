// package projects handles enabling/disabling Object Storage in STACKIT projects

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
	apiPath      = consts.API_PATH_OBJECT_STORAGE_PROJECT
	apiPathForce = consts.API_PATH_OBJECT_STORAGE_PROJECT_FORCE_DELETE
)

// New returns a new handler for the service
func New(c common.Client) *ObjectStorageProjectsService {
	return &ObjectStorageProjectsService{
		Client: c,
	}
}

// ObjectStorageProjectsService is the service that handles
// enabling / disabling Storage for a project
type ObjectStorageProjectsService common.Service

// ObjectStorageProjectResponse is the project ID and state response
type ObjectStorageProjectResponse struct {
	ProjectID string `json:"projectId"`
	Scope     string `json:"scope"`
}

// Get returns 200 if the Storage project is set, and error otherwise
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#tag/project
func (svc *ObjectStorageProjectsService) Get(ctx context.Context, projectID string) (res ObjectStorageProjectResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// Create creates an Storage project
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/create_project_v1_project__projectId__post
func (svc *ObjectStorageProjectsService) Create(ctx context.Context, projectID string) (res ObjectStorageProjectResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// Delete deletes an Storage project
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/delete_project_v1_project__projectId__delete
func (svc *ObjectStorageProjectsService) Delete(ctx context.Context, projectID string) (err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, nil)
	return
}

// ForceDelete force deletes an Storage project
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/delete_project_v1_project__projectId__delete
func (svc *ObjectStorageProjectsService) ForceDelete(ctx context.Context, projectID string) (err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPathForce, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, nil)
	return
}
