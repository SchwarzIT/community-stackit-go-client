// package options is used to retrieve various options used for configuring MongoDB Flex
// Such as available versions and storage size

package options

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPathVersions = consts.API_PATH_MONGO_DB_FLEX_VERSIONS
	apiPathStorage  = consts.API_PATH_MONGO_DB_FLEX_STORAGE
)

// New returns a new handler for the service
func New(c common.Client) *MongoDBOptionsService {
	return &MongoDBOptionsService{
		Client: c,
	}
}

// MongoDBOptionsService is the service that retrieves the provider options
type MongoDBOptionsService common.Service

// ListVersionsResponse is the APIs response for available versions
type ListVersionsResponse struct {
	Versions []string `json:"versions,omitempty"`
}

// GetStorageResponse is the API response for the storage range for a StorageClass
type GetStorageResponse struct {
	StorageClasses []string `json:"storageClasses,omitempty"`
	StorageRange   struct {
		Max int `json:"max,omitempty"`
		Min int `json:"min,omitempty"`
	} `json:"storageRange,omitempty"`
}

// ListVersions returns all available MongoDB Flex versions
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#/paths/~1projects~1{projectId}~1versions/get
func (svc *MongoDBOptionsService) ListVersions(ctx context.Context, projectID string) (res ListVersionsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathVersions, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}

// GetStorage returns available storage reange for a given flavor
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#/paths/~1projects~1{projectId}~1versions/get
func (svc *MongoDBOptionsService) GetStorage(ctx context.Context, projectID, flavor string) (res GetStorageResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathStorage, projectID, flavor), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}
