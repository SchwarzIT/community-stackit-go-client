// package permissions is used to retrieve all available permissions
// as well as permissions assigned to a user

package permissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/members"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// constants
const (
	apiPath     = consts.API_PATH_MEMBERSHIP_V2_PERMISSIONS
	apiPathUser = consts.API_PATH_MEMBERSHIP_V2_USER_PERMISSIONS
)

// New returns a new handler for the service
func New(c common.Client) *PermissionsService {
	return &PermissionsService{
		Client: c,
	}
}

// Public types

// PermissionsService is the service that handles
// CRUD functionality for permissions
type PermissionsService common.Service

// UserPermissions is a struct listing user permissions for every resource
type UserPermissions struct {
	Items  []ResourcePermissions `json:"items,omitempty"`
	Limit  int                   `json:"limit,omitempty"`
	Offset int                   `json:"offset,omitempty"`
}

// ResourcePermissions describes the resource and permissions assigned
type ResourcePermissions struct {
	ResourceID   string       `json:"resourceId,omitempty"`
	ResourceType string       `json:"resourceType,omitempty"`
	Permissions  []Permission `json:"permissions,omitempty"`
}

// PermissionList is a list of all permissions
type PermissionList struct {
	Permissions []Permission `json:"permissions,omitempty"`
}

// Permission describes a STACKIT permission
type Permission struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Memberships []members.Member `json:"memberships,omitempty"`
}

// List returns the a list of all available permissions
// Reference: https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/get-permissions
func (svc *PermissionsService) List(ctx context.Context) (p PermissionList, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, apiPath, nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &p)
	return
}

// GetByEmail returns a list of resources and permissions assigned to a user by their email
// Reference: https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/get-user-permissions
func (svc *PermissionsService) GetByEmail(ctx context.Context, email string, limit, offset int) (up UserPermissions, err error) {
	if err = validatePagination(limit, offset); err != nil {
		err = validate.WrapError(err)
		return
	}

	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathUser, email), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &up)
	return
}
