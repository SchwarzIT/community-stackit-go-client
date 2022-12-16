// package options is used to retrieve various options used for configuring a SKE cluster
// Such as available Kubernetes versions, machine types and more

package options

import (
	"context"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_SKE_OPTIONS
)

// New returns a new handler for the service
func New(c common.Client) *KubernetesOptionsService {
	return &KubernetesOptionsService{
		Client: c,
	}
}

// KubernetesOptionsService is the service that retrieves the provider options
type KubernetesOptionsService common.Service

// ProviderOptions is the api's provider options response struct
type ProviderOptions struct {
	KubernetesVersions []KubernetesVersion `json:"kubernetesVersions"`
	MachineTypes       []MachineType       `json:"machineTypes"`
	MachineImages      []MachineImage      `json:"machineImages"`
	VolumeTypes        []VolumeType        `json:"volumeTypes"`
	AvailabilityZones  []AvailabilityZone  `json:"availabilityZones"`
}

type KubernetesVersion struct {
	Version        string            `json:"version"`
	State          string            `json:"state"`
	ExpirationDate string            `json:"expirationDate"`
	FeatureGates   map[string]string `json:"featureGates"`
}

type MachineType struct {
	Name   string `json:"name"`
	CPU    int    `json:"cpu"`
	Memory int    `json:"memory"`
}

type MachineImage struct {
	Name     string `json:"name"`
	Versions []struct {
		Version        string `json:"version"`
		State          string `json:"state"`
		ExpirationDate string `json:"expirationDate"`
		CRI            []struct {
			Name string `json:"name"`
		} `json:"cri"`
	} `json:"versions"`
}

type VolumeType struct {
	Name string `json:"name"`
}

type AvailabilityZone struct {
	Name string `json:"name"`
}

// List returns all of the SKE provider options
// See also https://api.stackit.schwarz/ske-service/openapi.v1.html#tag/ProviderOptions
func (svc *KubernetesOptionsService) List(ctx context.Context) (res ProviderOptions, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, apiPath, nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}
