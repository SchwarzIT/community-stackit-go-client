// package costs is used for retrieving cost related information from stackit

package costs

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

const (
	apiPath         = consts.API_PATH_COSTS_WITH_PARAMS
	apiPathProjects = consts.API_PATH_COSTS_PROJECT_WITH_PARAMS

	timeFormat = "2006-01-02"
)

// CostsService is the client costs service
type CostsService common.Service

// New creates a new CostsService, which can be used to retrieve cost
// informations for customer accounts and specific projects.
func New(client common.Client) *CostsService {
	return &CostsService{
		Client: client,
	}
}

// responses

// CustomerAccountResponse is the response for customer account costs
type CustomerAccountResponse []ProjectDetailResponse

// ProjectCostsResponse is the response for project costs
type ProjectCostsResponse struct {
	ProjectID     string  `json:"projectId"`
	ProjectName   string  `json:"projectName"`
	TotalCharge   float64 `json:"totalCharge"`
	TotalDiscount float64 `json:"totalDiscount"`
}

// Date represents a date range
type Date struct {
	From string `json:"start"`
	To   string `json:"end"`
}

// ReportDatapoint is a report data struct
type ReportDatapoint struct {
	Date     Date    `json:"timePeriod"`
	Charge   float64 `json:"charge"`
	Discount float64 `json:"discount"`
	Quantity int64   `json:"quantity"`
}

// ServiceResponse is the costs for service
type ServiceResponse struct {
	TotalQuantity       int64             `json:"totalQuantity"`
	TotalCharge         float64           `json:"totalCharge"`
	TotalDiscount       float64           `json:"totalDiscount"`
	UnitLabel           string            `json:"unitLabel"`
	Sku                 string            `json:"sku"`
	ServiceName         string            `json:"serviceName"`
	ServiceCategoryName string            `json:"serviceCategoryName"`
	ReportData          []ReportDatapoint `json:"reportData"`
}

// ProjectDetailResponse is the costs for project and its services
type ProjectDetailResponse struct {
	ProjectCostsResponse
	Services []ServiceResponse `json:"services"`
}

// GetCustomerAccountCosts fetches cost data for a certain customer account in
// a specific time frame. The response is a list of project cost summaries.
// It provides a general cost overview for a certain customer account.
// See https://api.stackit.schwarz/costs-service/openapi.v1.html#operation/get-costs-reports-customer-account
func (svc *CostsService) GetCustomerAccountCosts(
	ctx context.Context,
	customerAccountID string,
	from, to time.Time,
	granularity, depth string,
) (*CustomerAccountResponse, error) {
	path := fmt.Sprintf(apiPath,
		customerAccountID,
		from.Format(timeFormat),
		to.Format(timeFormat),
		granularity,
		depth,
	)
	req, err := svc.Client.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var resp CustomerAccountResponse
	_, err = svc.Client.Do(req, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetProjectCosts fetches detailed cost data for a project in a
// specific time frame. The response contains all costs broken down into
// services used by the project.
// See https://api.stackit.schwarz/costs-service/openapi.v1.html#operation/get-costs-project
func (svc *CostsService) GetProjectCosts(
	ctx context.Context,
	customerAccountID,
	projectID string,
	from, to time.Time,
	granularity, depth string,
) (*ProjectDetailResponse, error) {
	err := validate.ProjectID(projectID)
	if err != nil {
		return nil, validate.WrapError(err)
	}

	path := fmt.Sprintf(apiPathProjects,
		customerAccountID,
		projectID,
		from.Format(timeFormat),
		to.Format(timeFormat),
		granularity,
		depth,
	)

	req, err := svc.Client.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var resp ProjectDetailResponse
	_, err = svc.Client.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
