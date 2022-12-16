// package roles is used for creating and managing custom roles (and permissions assigned to them)
package roles

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/permissions"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath         = consts.API_PATH_MEMBERSHIP_V2_ROLES
	apiPathWithType = consts.API_PATH_MEMBERSHIP_V2_ROLES_WITH_RESOURCE_TYPE
)

// New returns a new handler for the service
func New(c common.Client) *RolesService {
	return &RolesService{
		Client: c,
	}
}

// Public types

// RolesService is the service that handles
// CRUD functionality for membership roles
type RolesService common.Service

// Roles describes the resource and roles assigned to it
type Roles struct {
	ResourceID   string `json:"resourceId,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Roles        []Role `json:"roles,omitempty"`
}

// Role is a STACKIT role (or a custom user role)
type Role struct {
	Name        string                   `json:"name,omitempty"`
	Description string                   `json:"description,omitempty"`
	Permissions []permissions.Permission `json:"permissions,omitempty"`
}

// AddCustom adds new user specified roles to a resource
// https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/patch-roles
func (svc *RolesService) AddCustom(ctx context.Context, resourceID, resourceType string, roles ...Role) (res Roles, err error) {
	body, _ := svc.buildRequest(resourceType, roles)
	req, err := svc.Client.Request(ctx, http.MethodPatch, fmt.Sprintf(apiPath, resourceID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// RemoveCustom removes custom user specified roles from a resource
// Reference: https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/post-roles-remove
func (svc *RolesService) RemoveCustom(ctx context.Context, resourceID, resourceType string, roles ...Role) (r Roles, err error) {
	body, _ := svc.buildRequest(resourceType, roles)
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPath+"/remove", resourceID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &r)
	return
}

func (svc *RolesService) buildRequest(resourceType string, roles []Role) ([]byte, error) {
	return json.Marshal(Roles{
		ResourceType: resourceType,
		Roles:        roles,
	})
}

// GetByResource returns a list of roles and permissions by resources type & ID
// Reference: https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/get-roles
func (svc *RolesService) GetByResource(ctx context.Context, resourceID, resourceType string) (r Roles, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathWithType, resourceType, resourceID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &r)
	return
}
