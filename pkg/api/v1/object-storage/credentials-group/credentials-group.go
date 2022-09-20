// package credentialsGroups handles creation and management of STACKIT Object Storage credentialsGroups

package credentialsgroup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// constants
const (
	listPath   = consts.API_PATH_OBJECT_STORAGE_CREDENTIALS_LIST
	createPath = consts.API_PATH_OBJECT_STORAGE_CREDENTIALS_CREATE
	deletePath = consts.API_PATH_OBJECT_STORAGE_CREDENTIALS_DELETE
)

// New returns a new handler for the service
func New(c common.Client) *ObjectStorageCredentialsGroupService {
	return &ObjectStorageCredentialsGroupService{
		Client: c,
	}
}

// ObjectStorageCredentialsGroupService is the service that handles
// CRUD functionality for SKE credentialsGroups
type ObjectStorageCredentialsGroupService common.Service

// CredentialsGroupResponse is a struct representation of stackit's object storage api response for a credentialsGroup
type CredentialsGroupResponse struct {
	Project           string             `json:"project"`
	CredentialsGroups []CredentialsGroup `json:"credentialsGroups"`
}

type CredentialsGroupDeleteResponse struct {
	Project            string `json:"project"`
	CredentialsGroupId string `json:"credentialsGroupId"`
}

// CredentialsGroup holds all the credentialsGroup information
type CredentialsGroup struct {
	CredentialsGroupId string `json:"credentialsGroupId"`
	Urn                string `json:"urn"`
	DisplayName        string `json:"displayName"`
}

// List returns the a credentialsGroup by project ID and credentialsGroup name
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/get-credentials-groups-v1-project-projectId-credentials-groups-get
func (svc *ObjectStorageCredentialsGroupService) List(ctx context.Context, projectID string) (res CredentialsGroupResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(listPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// Create creates a new credentialsGroup
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/create-credentials-group-v1-project-projectId-credentials-group-post
func (svc *ObjectStorageCredentialsGroupService) Create(ctx context.Context, projectID, displayName string) (err error) {
	if err = ValidateDisplayName(displayName); err != nil {
		err = validate.WrapError(err)
		return
	}

	o := CredentialsGroup{DisplayName: displayName}
	body, _ := json.Marshal(o)
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(createPath, projectID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, nil)
	return
}

// Delete deletes a credentialsGroup
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/delete-credentials-group-v1-project-projectId-credentials-group-groupId-delete
func (svc *ObjectStorageCredentialsGroupService) Delete(ctx context.Context, projectID, credentialsGroupId string) (err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(deletePath, projectID, credentialsGroupId), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, nil)
	return err
}
