// package plans is used for retrieving information about Argus plans

package plans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_ARGUS_PLANS
)

// New returns a new handler for the service
func New(c common.Client) *PlansService {
	return &PlansService{
		Client: c,
	}
}

// PlansService is the service that handles
// CRUD functionality for Argus Plans
type PlansService common.Service

type PlanList struct {
	Message string `json:"message"`
	Plans   []Plan `json:"plans"`
}

type Plan struct {
	PlanID                  string  `json:"planId,omitempty"`
	ID                      string  `json:"id,omitempty"`
	Description             string  `json:"description,omitempty"`
	Name                    string  `json:"name,omitempty"`
	BucketSize              int     `json:"bucketSize,omitempty"`
	GrafanaGlobalUsers      int     `json:"grafanaGlobalUsers,omitempty"`
	GrafanaGlobalOrgs       int     `json:"grafanaGlobalOrgs,omitempty"`
	GrafanaGlobalDashboards int     `json:"grafanaGlobalDashboards,omitempty"`
	AlertRules              int     `json:"alertRules,omitempty"`
	TargetNumber            int     `json:"targetNumber,omitempty"`
	SamplesPerScrape        int     `json:"samplesPerScrape,omitempty"`
	GrafanaGlobalSessions   int     `json:"grafanaGlobalSessions,omitempty"`
	Amount                  float32 `json:"amount,omitempty"`
	AlertReceivers          int     `json:"alertReceivers,omitempty"`
	AlertMatchers           int     `json:"alertMatchers,omitempty"`
	TracesStorage           int     `json:"tracesStorage,omitempty"`
	LogsStorage             int     `json:"logsStorage,omitempty"`
	LogsAlert               int     `json:"logsAlert,omitempty"`
	IsFree                  bool    `json:"isFree,omitempty"`
	IsPublic                bool    `json:"isPublic,omitempty"`
	Schema                  string  `json:"schema,omitempty"`
}

// List returns a list of argus plans
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_plans_list
func (svc *PlansService) List(ctx context.Context, projectID string) (res PlanList, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.Do(req, &res)
	return
}
