// package projects handles creation and management of STACKIT projects
// For this use case, the Service Account used may need special permissions to manage STACKIT projects
// If needed, contact STACKIT support for further assistance on the matter

package projects

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
	"github.com/pkg/errors"
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

// Project struct holds important info about a STACKIT project
type Project struct {
	ID               string
	Name             string
	BillingReference string
	OrganizationID   string
}

// ProjectRole represents a role in the project
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

type projectsCreateReqBody struct {
	Name    string                 `json:"name"`
	Scope   string                 `json:"scope"`
	Members projectsMembersReqBody `json:"members"`
	Labels  projectsLabelsReqBody  `json:"labels"`
}

type projectsUpdateReqBody struct {
	Name   string                `json:"name"`
	Labels projectsLabelsReqBody `json:"labels"`
}

type projectsEntityReqBody struct {
	Role string `json:"role"`
	ID   string `json:"id"`
}

type projectsMembersReqBody struct {
	Users           []projectsEntityReqBody `json:"users"`
	ServiceAccounts []projectsEntityReqBody `json:"serviceAccounts"`
}

type projectsLabelsReqBody struct {
	BillingReference string `json:"billingReference"`
}

// Responses

// ProjectsResBody is the generic api response struct
type ProjectsResBody struct {
	ProjectID      string                `json:"projectId"`
	LifecycleState string                `json:"lifecycleState"`
	Scope          string                `json:"scope"`
	Name           string                `json:"name"`
	CreateTime     string                `json:"createTime"`
	Labels         ProjectsLabelsResBody `json:"labels"`
	Parent         ProjectsParentResBody
}

// ProjectsLabelsResBody is the labels response
type ProjectsLabelsResBody struct {
	BillingReference string `json:"billingReference"`
}

// ProjectsParentResBody is the parent entity response
type ProjectsParentResBody struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Implementation

// Create creates a new STACKIT project
// See also https://api.stackit.schwarz/resource-management/openapi.v1.html#operation/post-organizations-organizationId-projects
func (svc *ProjectService) Create(ctx context.Context, name, billingRef string, roles ...ProjectRole) (Project, error) {
	if err := ValidateProjectCreationRoles(roles); err != nil {
		return Project{}, validate.WrapError(err)
	}
	if err := validate.ProjectName(name); err != nil {
		return Project{}, validate.WrapError(err)
	}

	if err := validate.BillingRef(billingRef); err != nil {
		return Project{}, validate.WrapError(err)
	}

	body, err := svc.buildCreateRequestBody(name, billingRef, roles...)
	if err != nil {
		return Project{}, err
	}

	req, err := svc.Client.Request(
		ctx,
		http.MethodPost,
		fmt.Sprintf(apiPathCreate, svc.Client.OrganizationID()),
		body,
	)
	if err != nil {
		return Project{}, err
	}

	resBody := &ProjectsResBody{}
	if _, err = svc.Client.Do(req, resBody); err != nil {
		return Project{}, errors.Wrap(err, fmt.Sprintf("request was:\n%s", string(body)))
	}

	return Project{
		ID:               resBody.ProjectID,
		Name:             resBody.Name,
		BillingReference: resBody.Labels.BillingReference,
		OrganizationID:   resBody.Parent.ID,
	}, nil
}

// CreateAndWait wraps around `Create` and runs it with retry mechanism (which can be overridden by specifying a retry)
// it returns a wait service - by running Do() the wait service will wait for the project to be in active state
func (svc *ProjectService) CreateAndWait(ctx context.Context, name, billingRef string, roles []ProjectRole) (Project, *wait.Wait, error) {
	p, err := svc.Create(ctx, name, billingRef, roles...)
	if err != nil {
		return p, nil, err
	}

	w := wait.New(func() (interface{}, bool, error) {
		state, err := svc.GetLifecycleState(ctx, p.ID)
		if err != nil {
			return state, false, err
		}
		if state != consts.PROJECT_STATUS_ACTIVE {
			return state, false, nil
		}
		return state, true, nil
	})

	return p, w, err
}

func (svc *ProjectService) buildCreateRequestBody(name, billingRef string, roles ...ProjectRole) ([]byte, error) {
	users := []projectsEntityReqBody{}
	serviceAcounts := []projectsEntityReqBody{}
	isMemberDefined := false
	for _, r := range roles {
		if len(r.Users) > 0 {
			isMemberDefined = true
			users = append(users, projectsEntityReqBody{
				Role: r.Name,
				ID:   r.Users[0].ID,
			})
		}
		if len(r.ServiceAccounts) > 0 {
			isMemberDefined = true
			serviceAcounts = append(serviceAcounts, projectsEntityReqBody{
				Role: r.Name,
				ID:   r.ServiceAccounts[0].ID,
			})
		}
	}
	if !isMemberDefined {
		return []byte{}, errors.New("no user ID or service account ID provided")
	}
	return json.Marshal(projectsCreateReqBody{
		Name:  name,
		Scope: consts.PROJECT_SCOPE_PUBLIC,
		Members: projectsMembersReqBody{
			ServiceAccounts: serviceAcounts,
			Users:           users,
		},
		Labels: projectsLabelsReqBody{
			BillingReference: billingRef,
		},
	})
}

// Get returns the project by id
// See also https://api.stackit.schwarz/resource-management/openapi.v1.html#operation/get-projects-projectId
func (svc *ProjectService) Get(ctx context.Context, projectID string) (Project, error) {
	req, err := svc.Client.Request(
		ctx,
		http.MethodGet,
		fmt.Sprintf(apiPath, projectID),
		nil,
	)
	if err != nil {
		return Project{}, err
	}

	resBody := &ProjectsResBody{}
	if _, err = svc.Client.Do(req, resBody); err != nil {
		return Project{}, err
	}

	return Project{
		ID:               resBody.ProjectID,
		Name:             resBody.Name,
		BillingReference: resBody.Labels.BillingReference,
		OrganizationID:   resBody.Parent.ID,
	}, nil
}

// GetLifecycleState returns the project state
// See also https://api.stackit.schwarz/resource-management/openapi.v1.html#operation/get-projects-projectId
func (svc *ProjectService) GetLifecycleState(ctx context.Context, projectID string) (string, error) {
	req, err := svc.Client.Request(
		ctx,
		http.MethodGet,
		fmt.Sprintf(apiPath, projectID),
		nil,
	)
	if err != nil {
		return "", err
	}

	resBody := &ProjectsResBody{}
	if _, err = svc.Client.Do(req, resBody); err != nil {
		return "", err
	}

	return resBody.LifecycleState, nil
}

// Update updates an existing STACKIT project
// See also https://api.stackit.schwarz/resource-management/openapi.v1.html#operation/patch-project-projectId
func (svc *ProjectService) Update(ctx context.Context, id, name, billingRef string) error {
	if err := validate.ProjectName(name); err != nil {
		return validate.WrapError(err)
	}

	if err := validate.BillingRef(billingRef); err != nil {
		return validate.WrapError(err)
	}

	body, err := svc.buildUpdateRequestBody(name, billingRef)
	if err != nil {
		return err
	}

	req, err := svc.Client.Request(
		ctx,
		http.MethodPatch,
		fmt.Sprintf(apiPath, id),
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

func (svc *ProjectService) buildUpdateRequestBody(name, billingRef string) ([]byte, error) {
	return json.Marshal(projectsUpdateReqBody{
		Name: name,
		Labels: projectsLabelsReqBody{
			BillingReference: billingRef,
		},
	})
}

// UpdateAndWait wraps around `Update` and runs it with retry mechanism (which can be overridden by specifying a retry)
// it returns a wait service - by running Do() the wait service will wait for the project to be in active state
func (svc *ProjectService) UpdateAndWait(ctx context.Context, id, name, billingRef string) (*wait.Wait, error) {
	err := svc.Update(ctx, id, name, billingRef)
	if err != nil {
		return nil, err
	}

	w := wait.New(func() (interface{}, bool, error) {
		state, err := svc.GetLifecycleState(ctx, id)
		if err != nil {
			return state, false, err
		}
		if state != consts.PROJECT_STATUS_ACTIVE {
			return state, false, nil
		}
		return state, true, nil
	})

	return w, err
}

// Delete deletes a project by ID
// See also https://api.stackit.schwarz/resource-management/openapi.v1.html#operation/delete-projects-projectId
func (svc *ProjectService) Delete(ctx context.Context, projectID string) error {
	req, err := svc.Client.Request(
		ctx,
		http.MethodDelete,
		fmt.Sprintf(apiPath, projectID),
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

// DeleteAndWait wraps around `Delete` and runs it with retry mechanism (which can be overridden by specifying a retry)
// it returns a wait service - by running Do() the wait service will wait for the project to be deleted
func (svc *ProjectService) DeleteAndWait(ctx context.Context, projectID string) (*wait.Wait, error) {
	err := svc.Delete(ctx, projectID)
	if err != nil {
		return nil, err
	}

	w := wait.New(func() (interface{}, bool, error) {
		state, err := svc.GetLifecycleState(ctx, projectID)
		if err != nil {
			return state, true, nil
		}
		return state, false, nil
	})

	return w, err
}
