// package projects is used for retrieving organization information using STACKIT's Resource Manager v2 API

package organizations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_RESOURCE_MANAGEMENT_V2_ORG
)

// New returns a new handler for the service
func New(c common.Client) *OrganizationsService {
	return &OrganizationsService{
		Client: c,
	}
}

// OrganizationsService is the service that handles
// CRUD functionality for STACKIT organizations
type OrganizationsService common.Service

// structures

// OrganizationResponse is the structure representing the response for organization from the API
type OrganizationResponse struct {
	Name           string            `json:"name"`
	ContainerID    string            `json:"containerId"`
	OrganizationID string            `json:"organizationId"`
	LifecycleState string            `json:"lifecycleState"`
	CreationTime   string            `json:"creationTime"`
	UpdateTime     string            `json:"updateTime"`
	Labels         map[string]string `json:"labels"`
}

// Implementation

// Get returns the organization by container id
// See also https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/get-organizations-containerId
func (svc *OrganizationsService) Get(ctx context.Context, containerID string) (res OrganizationResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, containerID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}
