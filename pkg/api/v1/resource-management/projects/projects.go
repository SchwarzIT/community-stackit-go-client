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
	apiPath       = consts.API_PATH_RESOURCE_MANAGEMENT_PROJECTS
	apiPathCreate = consts.API_PATH_RESOURCE_MANAGEMENT_ORG_PROJECTS
)

// New returns a new handler for the service
func New(c common.Client) *ProjectService {
	return &ProjectService{
		Client: c,
	}
}

// ProjectService is the service that handles
// CRUD functionality for STACKIT projects
type ProjectService common.Service

// ProjectGetResponse is the generic api response struct
type ProjectGetResponse struct {
	ProjectID      string         `json:"projectId"`
	ContainerID    string         `json:"containerId"`
	LifecycleState string         `json:"lifecycleState"`
	Scope          string         `json:"scope"`
	Name           string         `json:"name"`
	CreateTime     string         `json:"createTime"`
	Labels         ProjectsLabels `json:"labels"`
	Parent         ProjectsParent
}

// ProjectsLabels is the labels response
type ProjectsLabels struct {
	BillingReference string `json:"billingReference"`
}

// ProjectsParent is the parent entity response
type ProjectsParent struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Get returns the project by id
// See also https://api.stackit.schwarz/resource-management/openapi.v1.html#operation/get-projects-projectId
func (svc *ProjectService) Get(ctx context.Context, projectID string) (res ProjectGetResponse, err error) {
	req, err := svc.Client.Request(
		ctx,
		http.MethodGet,
		fmt.Sprintf(apiPath, projectID),
		nil,
	)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}
