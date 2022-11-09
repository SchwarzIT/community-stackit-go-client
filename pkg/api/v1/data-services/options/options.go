// package options is used to retrieve various options used for configuring DSA

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
	apiPathOfferings = consts.API_PATH_DSA_OFFERINGS
)

// New returns a new handler for the service
func New(c common.Client) *DSAOptionsService {
	return &DSAOptionsService{
		Client: c,
	}
}

// DSAOptionsService is the service that retrieves the DSA options
type DSAOptionsService common.Service

// OfferingsResponse is the APIs response for available offerings
type OfferingsResponse struct {
	Offerings []Offer `json:"offerings,omitempty"`
}

// Offering represents a single DSA offer
type Offer struct {
	Name             string `json:"name"`
	Version          string `json:"version"`
	Latest           bool   `json:"latest"`
	DocumentationURL string `json:"documentationUrl"`
	Description      string `json:"description"`
	QuotaCount       int    `json:"quotaCount"`
	ImageURL         string `json:"imageUrl"`
	Schema           Schema `json:"schema"`
	Plans            []Plan `json:"plans"`
}

// Schema is an ofer schema struct
type Schema struct {
	Create ActionSetup `json:"create"`
	Update ActionSetup `json:"update"`
}

// ActionSetup is the setup of action such as create or update
type ActionSetup struct {
	Parameters map[string]string `json:"parameters"`
}

// Plan is a single plan an offer provides
type Plan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Free        bool   `json:"free"`
}

// GetVersions returns all available DSA offerings
// See also https://api.stackit.schwarz/data-services/openapi.v1.html#tag/Offerings
func (svc *DSAOptionsService) GetOfferings(ctx context.Context, projectID string) (res OfferingsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathOfferings, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}
