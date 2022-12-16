// package credentials is used to manage DSA instance credentials

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
	apiPathList   = consts.API_PATH_DSA_CREDENTIALS
	apiPathCreate = consts.API_PATH_DSA_CREDENTIALS
	apiPathGet    = consts.API_PATH_DSA_CREDENTIAL
	apiPathDelete = consts.API_PATH_DSA_CREDENTIAL
)

// New returns a new handler for the service
func New(c common.Client) *DSACredentialsService {
	return &DSACredentialsService{
		Client: c,
	}
}

// DSACredentialsService is the service that retrieves the DSA options
type DSACredentialsService common.Service

// ListResponse is the APIs response for listing all instance credentials
type ListResponse struct {
	CredentialsList []CredentialsListItem `json:"credentialsList,omitempty"`
}

// CredentialsListItem is a single credential ID from ListResponse
type CredentialsListItem struct {
	ID string `json:"id,omitempty"`
}

// GetResponse is the response struct for GET call
type GetResponse struct {
	ID  string        `json:"id,omitempty"`
	URI string        `json:"uri,omitempty"`
	Raw RawCredential `json:"raw,omitempty"`
}

// CreateResponse is the response struct for POST call
type CreateResponse GetResponse

// RawCredential contains the full credential information
type RawCredential struct {
	Credential      Credential `json:"credentials,omitempty"`
	SyslogDrainURL  string     `json:"syslogDrainUrl,omitempty"`
	RouteServiceURL string     `json:"routeServiceUrl,omitempty"`
}

// Credential holds the credential information
type Credential struct {
	Port     int      `json:"port,omitempty"`
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
	Hosts    []string `json:"hosts,omitempty"`
	Host     string   `json:"host,omitempty"`
	URI      string   `json:"uri,omitempty"`
	Cacrt    string   `json:"cacrt,omitempty"`
	Scheme   string   `json:"scheme,omitempty"`
}

// DeleteResponse is the delete request response struct
type DeleteResponse struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all instance credentials
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Credentials.get
func (svc *DSACredentialsService) List(ctx context.Context, projectID, instanceID string) (res ListResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathList, projectID, instanceID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Get returns a signle instance credential
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Credentials.get
func (svc *DSACredentialsService) Get(ctx context.Context, projectID, instanceID, credentialID string) (res GetResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathGet, projectID, instanceID, credentialID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Create creates a new instance credential
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Credentials.Post
func (svc *DSACredentialsService) Create(ctx context.Context, projectID, instanceID string) (res CreateResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPathCreate, projectID, instanceID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Delete deletes an instance credential
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#operation/Credentials.delete
func (svc *DSACredentialsService) Delete(ctx context.Context, projectID, instanceID, credentialID string) (res DeleteResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPathDelete, projectID, instanceID, credentialID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}
