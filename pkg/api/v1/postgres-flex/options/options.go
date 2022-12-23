// package options is used to retrieve various options used for configuring Postgres Flex
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
	apiPathVersions = consts.API_PATH_POSTGRES_FLEX_VERSIONS
	apiPathStorage  = consts.API_PATH_POSTGRES_FLEX_STORAGES
	apiPathFlavors  = consts.API_PATH_POSTGRES_FLEX_FLAVORS
)

// New returns a new handler for the service
func New(c common.Client) *PostgresOptionsService {
	return &PostgresOptionsService{
		Client: c,
	}
}

// PostgresOptionsService is the service that retrieves the provider options
type PostgresOptionsService common.Service

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
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	CPU         int    `json:"cpu,omitempty"`
	Memory      int    `json:"memory,omitempty"`
}

// GetVersions returns all available Postgres Flex versions
// See also https://api.stackit.schwarz/postgres-flex-service/openapi.html#tag/versions
func (svc *PostgresOptionsService) GetVersions(ctx context.Context, projectID string) (res VersionsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathVersions, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// GetFlavors returns available flavors
// See also https://api.stackit.schwarz/postgres-flex-service/openapi.html#tag/flavors
func (svc *PostgresOptionsService) GetFlavors(ctx context.Context, projectID string) (res GetFlavorsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathFlavors, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// GetStorageClasses returns available storage reange for a given flavor
// See also https://api.stackit.schwarz/postgres-flex-service/openapi.html#tag/storage
func (svc *PostgresOptionsService) GetStorageClasses(ctx context.Context, projectID, flavorID string) (res GetStorageResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathStorage, projectID, flavorID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}
