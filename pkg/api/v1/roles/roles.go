// package roles is used for managing roles for users and service accounts

package roles

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
)

// constants
const (
	apiPath = consts.API_PATH_MEMBERSHIP_ROLES
)

// New returns a new handler for the service
func New(c common.Client) *RolesService {
	return &RolesService{
		Client: c,
	}
}

// RolesService is the service that handles
// CRUD functionality for users in roles in a STACKIT project
type RolesService common.Service

// Service Response

// ProjectRoles is the main response struct
// representing a project, its roles and members that belong to the role
type ProjectRoles struct {
	ProjectID string        `json:"projectID"`
	Roles     []ProjectRole `json:"roles"`
}

// ProjectRole represents a role and its members
type ProjectRole struct {
	Name            string              `json:"name"`
	Users           []ProjectRoleMember `json:"users"`
	ServiceAccounts []ProjectRoleMember `json:"service_accounts"`
}

// ProjectRoleMember represents a user or service account
type ProjectRoleMember struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
}

// Requests

// projectRolesAddUserReq represents a request to add users to a project role
type projectRolesAddUserReq struct {
	Role  string   `json:"role"`
	Users []string `json:"users"`
}

type projectRolesDeleteUserReq projectRolesAddUserReq

// projectRoleAddSAReq represents a request to add service account to a project role
type projectRoleAddSAReq struct {
	ServiceAccountID string `json:"serviceAccountId"`
}

// Responses

type projectRolesGetResBody struct {
	Items []struct {
		Role      string `json:"role"`
		ProjectID string `json:"projectId"`
		Members   []struct {
			GlobalID string `json:"globalId"`
			Email    string `json:"email,omitempty"`
		} `json:"members"`
	} `json:"items"`
}

// Implementation

// Get returns the project roles
// Reference: https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/get-projects-projectId-roles
func (svc *RolesService) Get(ctx context.Context, projectID string) (ProjectRoles, error) {
	req, err := svc.Client.Request(
		ctx,
		http.MethodGet,
		fmt.Sprintf(apiPath, projectID),
		nil,
	)
	if err != nil {
		return ProjectRoles{}, err
	}

	resBody := &projectRolesGetResBody{}
	if _, err = svc.Client.Do(req, resBody); err != nil {
		return ProjectRoles{}, err
	}

	return svc.transformResponse(resBody), nil
}

// AddUsers adds users and/or service accounts to a given project role
func (svc *RolesService) AddUsers(ctx context.Context, projectID, role string, users []string, serviceAccounts []string) error {
	if err := validate.Role(role); err != nil {
		return validate.WrapError(err)
	}
	for _, sa := range serviceAccounts {
		if err := svc.addServiceAccounts(ctx, projectID, role, sa); err != nil {
			return err
		}
	}
	return svc.addUsers(ctx, projectID, role, users...)
}

// addUsers adds users to project role
// Reference: https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/patch-projects-projectId-roles
func (svc *RolesService) addUsers(ctx context.Context, projectID, role string, users ...string) error {
	body, err := svc.buildAddUsersRequestBody(role, users...)
	if err != nil {
		return err
	}
	req, err := svc.Client.Request(
		ctx,
		http.MethodPatch,
		fmt.Sprintf(apiPath, projectID),
		body,
	)
	if err != nil {
		return err
	}

	if _, err = svc.Client.Do(req, nil); err != nil {
		return errors.Wrap(err, fmt.Sprintf("request was:\n%s", string(body)))
	}

	return nil
}

// addServiceAccounts adds service account to project role
// Reference: https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/post-organizations-organizationId-projects-projectId-roles-roleName-service-accounts
func (svc *RolesService) addServiceAccounts(ctx context.Context, projectID, role string, serviceAccountID string) error {
	body, err := svc.buildAddSARequestBody(serviceAccountID)
	if err != nil {
		return err
	}
	req, err := svc.Client.Request(
		ctx,
		http.MethodPost,
		fmt.Sprintf(consts.API_PATH_MEMBERSHIP_ORG_PROJECT_ROLE_SERVICE_ACCOUNTS, svc.Client.OrganizationID(), projectID, role),
		body,
	)
	if err != nil {
		return err
	}

	if _, err = svc.Client.Do(req, nil); err != nil {
		return errors.Wrap(err, fmt.Sprintf("request was:\n%s", string(body)))
	}

	return nil
}

func (svc *RolesService) buildAddUsersRequestBody(role string, userIDs ...string) ([]byte, error) {
	return json.Marshal([]projectRolesAddUserReq{{
		Role:  role,
		Users: userIDs,
	}})
}

func (svc *RolesService) buildAddSARequestBody(serviceAccountID string) ([]byte, error) {
	return json.Marshal(projectRoleAddSAReq{
		ServiceAccountID: serviceAccountID,
	})
}

// DeleteUsers removes users from a given role
// Reference: https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/patch-projects-projectId-roles
func (svc *RolesService) DeleteUsers(ctx context.Context, projectID, role string, users []string, serviceAccounts []string) error {
	if err := validate.Role(role); err != nil {
		return validate.WrapError(err)
	}
	for _, sa := range serviceAccounts {
		if err := svc.deleteServiceAccount(ctx, projectID, role, sa); err != nil {
			return err
		}
	}
	return svc.deleteUsers(ctx, projectID, role, users...)
}

// deleteUsers removes users from a given role
func (svc *RolesService) deleteUsers(ctx context.Context, projectID, role string, userIDs ...string) error {
	body, err := svc.buildDeleteUsersRequestBody(role, userIDs...)
	if err != nil {
		return err
	}
	req, err := svc.Client.Request(
		ctx,
		http.MethodPatch,
		fmt.Sprintf(consts.API_PATH_MEMBERSHIP_ROLES_DELETE, projectID),
		body,
	)
	if err != nil {
		return err
	}

	if _, err = svc.Client.Do(req, nil); err != nil {
		return errors.Wrap(err, fmt.Sprintf("request was:\n%s", string(body)))
	}

	return nil
}

// deleteServiceAccount removes a service account from a given role
// Reference: https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/patch-projects-projectId-roles
func (svc *RolesService) deleteServiceAccount(ctx context.Context, projectID, role string, serviceAccountID string) error {
	req, err := svc.Client.Request(
		ctx,
		http.MethodDelete,
		fmt.Sprintf(consts.API_PATH_MEMBERSHIP_ORG_PROJECT_ROLE_SERVICE_ACCOUNT, svc.Client.OrganizationID(), projectID, role, serviceAccountID),
		nil,
	)
	if err != nil {
		return err
	}

	if _, err = svc.Client.Do(req, nil); err != nil {
		return err
	}

	return nil
}

func (svc *RolesService) buildDeleteUsersRequestBody(role string, userIDs ...string) ([]byte, error) {
	return json.Marshal([]projectRolesDeleteUserReq{{
		Role:  role,
		Users: userIDs,
	}})
}

func (svc *RolesService) transformResponse(clientRes *projectRolesGetResBody) ProjectRoles {
	pr := ProjectRoles{Roles: []ProjectRole{}}
	for _, role := range clientRes.Items {
		pr.ProjectID = role.ProjectID
		users := []ProjectRoleMember{}
		sas := []ProjectRoleMember{}
		for _, m := range role.Members {
			if m.Email == "" {
				sas = append(sas, ProjectRoleMember{
					ID: m.GlobalID,
				})
				continue
			}
			users = append(users, ProjectRoleMember{
				Email: m.Email,
				ID:    m.GlobalID,
			})
		}
		pr.Roles = append(pr.Roles, ProjectRole{
			Name:            role.Role,
			Users:           users,
			ServiceAccounts: sas,
		})
	}
	return pr
}
