// package projects is used for creating & managing projects with STACKIT's Resource Manager v2 API

package projects

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPath        = consts.API_PATH_RESOURCE_MANAGER_V2_PROJECTS
	apiPathProject = consts.API_PATH_RESOURCE_MANAGER_V2_PROJECT
)

// New returns a new handler for the service
func New(c common.Client) *ProjectsService {
	return &ProjectsService{
		Client: c,
	}
}

// ProjectsService is the service that handles
// CRUD functionality for STACKIT projects
type ProjectsService common.Service

// Creation structures

// CreateProjectRequest is the structure representing the request body for creating a project
type CreateProjectRequest struct {
	Name     string            `json:"name"`
	ParentID string            `json:"containerParentId"`
	Members  []ProjectMember   `json:"members"`
	Labels   map[string]string `json:"labels"`
}

// ProjectMember is the structure representing a member of a project
// role can be one of "project.owner" "project.member" "project.admin" "project.auditor"
// the subject is the email address
type ProjectMember struct {
	Role    string `json:"role"`
	Subject string `json:"subject"` // email address
}

// MandatoryLabels represent the project's mandatory labels
type MandatoryLabels struct {
	BillingReference string `json:"billingReference"`
	Scope            string `json:"scope"`
}

func (m *MandatoryLabels) ToMap() map[string]string {
	var mapped map[string]string
	inrec, _ := json.Marshal(m)
	_ = json.Unmarshal(inrec, &mapped)
	return mapped
}

// ProjectResponse is the structure representing the server response for a project
type ProjectResponse struct {
	Name           string            `json:"name"`
	Parent         Parent            `json:"parent"`
	ContainerID    string            `json:"containerId"` // Globally unique, user friendly identifier
	ProjectID      string            `json:"projectId"`   // Legacy identifier
	LifecycleState string            `json:"lifecycleState"`
	Labels         map[string]string `json:"labels"`
	UpdateTime     string            `json:"updateTime"`
	CreationTime   string            `json:"creationTime"`
}

// Parent represents a project's parent details
type Parent struct {
	ID          string `json:"id"`
	ContainerID string `json:"containerId"` // User friendly container ID
	Type        string `json:"type"`        // ORGANIZATION or FOLDER
}

// ProjectsResponse is the List (get all) projects response from the API
type ProjectsResponse struct {
	Iteams []ProjectResponse `json:"items"`
	Offset int               `json:"offset"`
	Limit  int               `json:"limit"`
}

// Implementation

// Create creates a new STACKIT project
// See also https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/post-projects
func (svc *ProjectsService) Create(ctx context.Context, name string, labels map[string]string, members ...ProjectMember) (res ProjectResponse, w *wait.Handler, err error) {
	if err = svc.ValidateCreateData(name, labels, members); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildCreateRequest(name, labels, members)
	req, err := svc.Client.Request(ctx, http.MethodPost, apiPath, body)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	w = wait.New(svc.waitForCreation(ctx, res.ContainerID))
	return
}

func (svc *ProjectsService) buildCreateRequest(name string, labels map[string]string, members []ProjectMember) ([]byte, error) {
	return json.Marshal(CreateProjectRequest{
		Name:     name,
		ParentID: svc.Client.OrganizationID(),
		Members:  members,
		Labels:   labels,
	})
}

func (svc *ProjectsService) waitForCreation(ctx context.Context, containerID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		state, err := svc.GetLifecycleState(ctx, containerID)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusForbidden)) {
				return state, false, nil
			}
			return state, false, err
		}
		switch state {
		case consts.PROJECT_STATUS_ACTIVE:
			return state, true, nil
		case consts.PROJECT_STATUS_CREATING:
			return state, false, nil
		}
		return state, false, fmt.Errorf("received project state '%s'. aborting", state)
	}
}

// Get returns the project by id
// See also https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/get-projects-containerId
func (svc *ProjectsService) Get(ctx context.Context, containerID string) (res ProjectResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathProject, containerID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// List returns a list of projects
// if containerParentID == "" at least one containerID needs to be specified (and vice versa)
// See also https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/get-all-projects
func (svc *ProjectsService) List(ctx context.Context, containerParentID string, filters map[string]string, containerIDs ...string) (res ProjectsResponse, err error) {
	params := url.Values{}
	f := map[string]string{
		"offset":              "0",
		"limit":               "50",
		"creation-time-start": "",
	}
	for k := range f {
		if v, ok := filters[k]; ok {
			f[k] = v
		}
		if f[k] != "" {
			params.Add(k, f[k])
		}
	}

	if err = svc.ValidateList(containerParentID, containerIDs, f["offset"], f["limit"], f["creation-time-start"]); err != nil {
		return
	}

	for _, v := range containerIDs {
		params.Add("containerIds", v)
	}

	if containerParentID != "" {
		params.Add("containerParentId", containerParentID)
	}

	req, err := svc.Client.Request(ctx, http.MethodGet, apiPath+"?"+params.Encode(), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// GetLifecycleState returns the project state
// See also https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/get-projects-containerId
func (svc *ProjectsService) GetLifecycleState(ctx context.Context, containerID string) (string, error) {
	p, err := svc.Get(ctx, containerID)
	return p.LifecycleState, err
}

// Update structures

// UpdateProjectRequest is the update request structure
type UpdateProjectRequest struct {
	Name              string            `json:"name"`
	ContainerParentID string            `json:"containerParentId"`
	Labels            map[string]string `json:"labels"`
}

// Update updates an existing STACKIT project
// See also https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/patch-projects-containerId
func (svc *ProjectsService) Update(ctx context.Context, containerID, name, containerParentID string, labels map[string]string) (res ProjectResponse, err error) {
	if err = svc.ValidateUpdateData(containerID, containerParentID, name, labels); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildUpdateRequest(name, containerParentID, labels)
	req, err := svc.Client.Request(ctx, http.MethodPatch, fmt.Sprintf(apiPathProject, containerID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

func (svc *ProjectsService) buildUpdateRequest(name, containerParentID string, labels map[string]string) ([]byte, error) {
	return json.Marshal(UpdateProjectRequest{
		Name:              name,
		Labels:            labels,
		ContainerParentID: containerParentID,
	})
}

// Delete deletes a project by ID
// See also https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/delete-projects-containerId
func (svc *ProjectsService) Delete(ctx context.Context, containerID string) (w *wait.Handler, err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPathProject, containerID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, nil)
	w = wait.New(svc.waitForDeletion(ctx, containerID))
	return
}

func (svc *ProjectsService) waitForDeletion(ctx context.Context, containerID string) wait.WaitFn {
	return func() (interface{}, bool, error) {
		state, err := svc.GetLifecycleState(ctx, containerID)
		if err != nil {
			return state, true, nil
		}
		return state, false, nil
	}
}
