// package credentials implements calls for managing Argus instance credentials
// this package should be used after an Argus instance has already been created

package credentials

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_ARGUS_CREDENTIALS
)

// New returns a new handler for the service
func New(c common.Client) *CredentialsService {
	return &CredentialsService{
		Client: c,
	}
}

// CredentialList is the api response struct for credential list
type CredentialList struct {
	Message     string       `json:"message,omitempty"`
	Credentials []Credential `json:"credentials"`
}

// Credential is an item from credential list credentials
type Credential struct {
	Message         string `json:"message,omitempty"`
	Name            string `json:"name,omitempty"`
	ID              string `json:"id,omitempty"`
	CredentialsInfo struct {
		Username string `json:"username,omitempty"`
	} `json:"credentialsInfo,omitempty"`
}

// CredentialsInfo contains credential specific data
type CredentialsInfo struct {
	Username string `json:"username,omitempty"`
}

type CreateResponse struct {
	Message     string `json:"message,omitempty"`
	Credentials struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"credentials,omitempty"`
}

// CredentialsService is the service that handles
// CRUD functionality for Argus instance credentials
type CredentialsService common.Service

// List returns a list of argus instance credentials
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_credentials_list
func (svc *CredentialsService) List(ctx context.Context, projectID, instanceID string) (res CredentialList, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Get returns information about credential by username
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_credentials_read
func (svc *CredentialsService) Get(ctx context.Context, projectID, instanceID, username string) (res Credential, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath+"/%s", projectID, instanceID, username), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Create creates a new credential
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_credentials_create
func (svc *CredentialsService) Create(ctx context.Context, projectID, instanceID string) (res CreateResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPath, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Delete deletes a credential
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_credentials_delete
func (svc *CredentialsService) Delete(ctx context.Context, projectID, instanceID, username string) (res Credential, err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPath+"/%s", projectID, instanceID, username), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}
