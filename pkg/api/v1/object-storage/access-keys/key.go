// package keys handles management of Object Storage credentials

package keys

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	credentialsGroup "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage/credentials-group"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPathList             = consts.API_PATH_OBJECT_STORAGE_KEYS
	apiPathListWithParams   = consts.API_PATH_OBJECT_STORAGE_KEYS_WITH_PARAMS
	apiPathCreate           = consts.API_PATH_OBJECT_STORAGE_KEY
	apiPathCreateWithParams = consts.API_PATH_OBJECT_STORAGE_KEY_WITH_PARAMS
	apiPathDelete           = consts.API_PATH_OBJECT_STORAGE_WITH_KEY_ID
	apiPathDeleteWithParams = consts.API_PATH_OBJECT_STORAGE_WITH_KEY_ID_WITH_PARAMS
)

// New returns a new handler for the service
func New(c common.Client) *ObjectStorageAccessKeysService {
	return &ObjectStorageAccessKeysService{
		Client: c,
	}
}

// ObjectStorageAccessKeysService is the service that handles
// enabling / disabling AccessKeys for a project
type ObjectStorageAccessKeysService common.Service

// AccessKeyListResponse is the api list response struct for a given project ID
type AccessKeyListResponse struct {
	Project    string             `json:"project"`
	AccessKeys []AccessKeyDetails `json:"accessKeys"`
}

// AccessKeyCreateResponse is the struct representing a creation response
type AccessKeyCreateResponse struct {
	Project         string `json:"project"`
	DisplayName     string `json:"displayName"`
	KeyID           string `json:"keyId"`
	Expires         string `json:"expires"`
	AccessKey       string `json:"accessKey"`
	SecretAccessKey string `json:"secretAccessKey"`
}

// AccessKeyDetails is the minial information of an access key
type AccessKeyDetails struct {
	DisplayName string `json:"displayName"`
	KeyID       string `json:"keyId"`
	Expires     string `json:"expires"`
}

// Expiry represents the date and time in which the key will expire in
type Expiry struct {
	Expires string `json:"expires,omitempty"`
}

// List returns a list of access keys assigned to a given Project ID
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/get_access_keys_v1_project__projectId__access_keys_get
func (svc *ObjectStorageAccessKeysService) List(ctx context.Context, projectID, credentialsGroupID string) (res AccessKeyListResponse, err error) {
	if err = credentialsGroup.ValidateCredentialsGroupID(credentialsGroupID); err != nil && credentialsGroupID != "" {
		err = validate.WrapError(err)
		return
	}
	path := fmt.Sprintf(apiPathListWithParams, projectID, credentialsGroupID)
	if credentialsGroupID == "" {
		path = fmt.Sprintf(apiPathList, projectID)
	}
	req, err := svc.Client.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Create creates an Access Keys
// If expires is empty, the key will not expire
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/create_project_v1_project__projectId__post
func (svc *ObjectStorageAccessKeysService) Create(ctx context.Context, projectID, expires, credentialsGroupID string) (res AccessKeyCreateResponse, err error) {
	if err = credentialsGroup.ValidateCredentialsGroupID(credentialsGroupID); err != nil && credentialsGroupID != "" {
		err = validate.WrapError(err)
		return
	}
	path := fmt.Sprintf(apiPathCreateWithParams, projectID, credentialsGroupID)
	if credentialsGroupID == "" {
		path = fmt.Sprintf(apiPathCreate, projectID)
	}
	body, _ := json.Marshal(Expiry{
		Expires: expires,
	})
	req, err := svc.Client.Request(ctx, http.MethodPost, path, body)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Delete deletes an Access Keys by ID
// See also https://api.stackit.schwarz/object-storage-service/openapi.v1.html#operation/delete_access_key_v1_project__projectId__access_key__keyId__delete
func (svc *ObjectStorageAccessKeysService) Delete(ctx context.Context, projectID, keyID, credentialsGroupID string) (err error) {
	if err = credentialsGroup.ValidateCredentialsGroupID(credentialsGroupID); err != nil && credentialsGroupID != "" {
		err = validate.WrapError(err)
		return
	}
	path := fmt.Sprintf(apiPathDeleteWithParams, projectID, keyID, credentialsGroupID)
	if credentialsGroupID == "" {
		path = fmt.Sprintf(apiPathDelete, projectID, keyID)
	}
	req, err := svc.Client.Request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, nil)
	return
}
