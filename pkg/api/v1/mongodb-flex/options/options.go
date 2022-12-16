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
	apiPathStorage  = consts.API_PATH_MONGO_DB_FLEX_STORAGES
	apiPathFlavors  = consts.API_PATH_MONGO_DB_FLEX_FLAVORS
)

// New returns a new handler for the service
func New(c common.Client) *MongoDBOptionsService {
	return &MongoDBOptionsService{
		Client: c,
	}
}

// MongoDBOptionsService is the service that retrieves the provider options
type MongoDBOptionsService common.Service

// VersionsResponse is the APIs response for available versions
type VersionsResponse struct {
	Versions []string `json:"versions,omitempty"`
}

// GetStorageResponse is the API response for the storage range for a StorageClass
type GetStorageResponse struct {
	StorageClasses []string     `json:"storageClasses,omitempty"`
	StorageRange   StorageRange `json:"storageRange,omitempty"`
}

// StorageRange represents the storage size range
type StorageRange struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}

// GetFlavorsResponse is the server response to GetFlavors call
type GetFlavorsResponse struct {
	Flavors []Flavor `json:"flavors,omitempty"`
}

// Flavor represents a single flavor
type Flavor struct {
	ID          string   `json:"id,omitempty"`
	Description string   `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	CPU         int      `json:"cpu,omitempty"`
	Memory      int      `json:"memory,omitempty"`
}

// GetVersions returns all available MongoDB Flex versions
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#/paths/~1projects~1{projectId}~1versions/get
func (svc *MongoDBOptionsService) GetVersions(ctx context.Context, projectID string) (res VersionsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathVersions, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// GetFlavors returns available flavors
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#/paths/~1projects~1{projectId}~1flavors/get
func (svc *MongoDBOptionsService) GetFlavors(ctx context.Context, projectID string) (res GetFlavorsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathFlavors, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// GetStorageClasses returns available storage reange for a given flavor
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#/paths/~1projects~1{projectId}~1versions/get
func (svc *MongoDBOptionsService) GetStorageClasses(ctx context.Context, projectID, flavorID string) (res GetStorageResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathStorage, projectID, flavorID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}
