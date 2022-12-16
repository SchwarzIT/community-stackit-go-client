// package grafana implements calls for getting & updating grafana configuration for an Argus Instance
// therefore, it can only be used after an Argus Instance has been created

package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_ARGUS_GRAFANA_CONFIGS
)

// New returns a new handler for the service
func New(c common.Client) *GrafanaService {
	return &GrafanaService{
		Client: c,
	}
}

// GrafanaService is the service that handles
// CRUD functionality for Argus Grafana
type GrafanaService common.Service

type Config struct {
	PublicReadAccess bool   `json:"publicReadAccess"`
	OAuth            *OAuth `json:"genericOauth,omitempty"`
}

type GetConfigResponse struct {
	Message          string `json:"message,omitempty"`
	PublicReadAccess bool   `json:"publicReadAccess,omitempty"`
	OAuth            *OAuth `json:"genericOauth,omitempty"`
}

type OAuth struct {
	Enabled           bool   `json:"enabled,omitempty"`
	APIURL            string `json:"apiUrl,omitempty"`
	AuthURL           string `json:"authUrl,omitempty"`
	Scopes            string `json:"scopes,omitempty"`
	TokenURL          string `json:"tokenUrl,omitempty"`
	OauthClientID     string `json:"oauthClientId,omitempty"`
	OauthClientSecret string `json:"oauthClientSecret,omitempty"`

	// If therole_attribute_path property does not return a role, then the user is assigned the Viewer role by default.
	// You can disable the role assignment by setting role_attribute_strict = true.
	// It denies user access if no role or an invalid role is returned.
	RoleAttributeStrict bool `json:"roleAttributeStrict,omitempty"`

	// Grafana checks for the presence of a role using the JMESPath specified via the role_attribute_path
	// configuration option. The JMESPath is applied to the id_token first. If there is no match, then the UserInfo
	// endpoint specified via the api_url configuration option is tried next.
	// The result after evaluation of the role_attribute_path JMESPath expression should be a valid Grafana role,
	// for example, Viewer, Editor or Admin For example:
	// contains(roles[*], 'grafana-admin') && 'Admin' || contains(roles[*], 'grafana-editor') && 'Editor' || contains(roles[*], 'grafana-viewer') && 'Viewer'
	RoleAttributePath string `json:"roleAttributePath,omitempty"`
}

type UpdateConfigResponse struct {
	Message string `json:"message,omitempty"`
}

// GetConfig returns argus grafana config
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#tag/grafana-configs
func (svc *GrafanaService) GetConfig(ctx context.Context, projectID, instanceID string) (res GetConfigResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// UpdateConfig updates argus grafana config
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_grafana-configs_update
func (svc *GrafanaService) UpdateConfig(ctx context.Context, projectID, instanceID string, config Config) (res UpdateConfigResponse, err error) {
	body, _ := json.Marshal(config)
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPath, projectID, instanceID), body)
	if err != nil {
		return
	}
	_, err = svc.Client.LegacyDo(req, &res)
	return
}
