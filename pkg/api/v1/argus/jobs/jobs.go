// Package jobs is used to create and manage scrape jobs in an Argus instance
// Therefore, it can only be used after an Argus instance has been created

package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPath    = consts.API_PATH_ARGUS_JOBS
	apiPathJob = consts.API_PATH_ARGUS_JOBS_WITH_JOB_NAME
)

// New returns a new handler for the service
func New(c common.Client) *JobsService {
	return &JobsService{
		Client: c,
	}
}

// JobsService is the service that handles
// CRUD functionality for Argus metric collection jobs
type JobsService common.Service

// ListJobsResponse is a struct representing the api response for listing jobs
type ListJobsResponse struct {
	Message string `json:"message,omitempty"`
	Data    []Job  `json:"data,omitempty"`
}

// GetJobResponse is a struct representing the api response for a single job
type GetJobResponse struct {
	Message string `json:"message,omitempty"`
	Data    Job    `json:"data,omitempty"`
}

// Job is a struct representing a job
type Job struct {
	StaticConfigs         []StaticConfig         `json:"staticConfigs"`
	JobName               string                 `json:"jobName,omitempty"`
	Scheme                string                 `json:"scheme,omitempty"`
	ScrapeInterval        string                 `json:"scrapeInterval,omitempty"`
	ScrapeTimeout         string                 `json:"scrapeTimeout,omitempty"`
	MetricsPath           string                 `json:"metricsPath,omitempty"`
	SampleLimit           int                    `json:"sampleLimit,omitempty"`
	BasicAuth             *BasicAuth             `json:"basicAuth,omitempty"`
	OAuth2                *OAuth2                `json:"oauth2,omitempty"`
	TLSConfig             *TLSConfig             `json:"tlsConfig,omitempty"`
	BearerToken           string                 `json:"bearerToken,omitempty"`
	MetricsRelabelConfigs []MetricsRelabelConfig `json:"metricsRelabelConfigs,omitempty"`
	Params                map[string]interface{} `json:"params,omitempty"`
	ServiceDiscoryConfigs []ServiceDiscoryConfig `json:"httpSdConfigs,omitempty"`
	HonorLabels           bool                   `json:"honorLabels,omitempty"`
	HonorTimeStamps       bool                   `json:"honorTimeStamps,omitempty"`
}

// StaticConfig holds targets for scraping
type StaticConfig struct {
	Targets []string          `json:"targets,omitempty"`
	Labels  map[string]string `json:"labels,omitempty"`
}

// BasicAuth holds basic auth data
type BasicAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// OAuth2 holds OAuth2 data
type OAuth2 struct {
	ClientID     string    `json:"clientId,omitempty"`
	ClientSecret string    `json:"clientSecret,omitempty"`
	TokenURL     string    `json:"tokenUrl,omitempty"`
	Scopes       []string  `json:"scopes,omitempty"`
	TLSConfig    TLSConfig `json:"tlsConfig,omitempty"`
}

// TLSConfig determines if the insecureSkipVerify is enabled/disabled
type TLSConfig struct {
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

// MetricsRelabelConfig configuration for metric relabeling
type MetricsRelabelConfig struct {
	SourceLabels []string `json:"sourceLabels,omitempty"`
	Separator    string   `json:"separator,omitempty"`
	TargetLabel  string   `json:"targetLabel,omitempty"`
	Regex        string   `json:"regex,omitempty"`
	Modulus      int      `json:"modulus,omitempty"`
	Replacement  string   `json:"replacement,omitempty"`
	Action       string   `json:"action,omitempty"`
}

// ServiceDiscoryConfig is the configuration for service discovery
type ServiceDiscoryConfig struct {
	URL             string    `json:"url,omitempty"`
	RefreshInterval string    `json:"refreshInterval,omitempty"`
	BasicAuth       BasicAuth `json:"basicAuth,omitempty"`
	TLSConfig       TLSConfig `json:"tlsConfig,omitempty"`
	OAuth2          OAuth2    `json:"oauth2,omitempty"`
}

// List returns a list of argus jobs
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_scrapeconfigs_list
func (svc *JobsService) List(ctx context.Context, projectID, instanceID string) (res ListJobsResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPath, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Get returns an argus job by job name
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_scrapeconfigs_read
func (svc *JobsService) Get(ctx context.Context, projectID, instanceID, jobName string) (res GetJobResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathJob, projectID, instanceID, jobName), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Create creates an argus scrape job and return a list of jobs
// it also returns a wait service to verify the creation
// the wait service can be triggered using `Wait()` and it returns the job (GetJobResponse) and an error if occurred
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_scrapeconfigs_create
func (svc *JobsService) Create(ctx context.Context, projectID, instanceID string, job Job) (res ListJobsResponse, w *wait.Handler, err error) {
	if err = job.Validate(); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildCreateRequest(job)
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPath, projectID, instanceID), body)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	w = wait.New(svc.waitForCreation(ctx, projectID, instanceID, job.JobName))
	w.SetTimeout(10 * time.Minute)
	return
}

func (svc *JobsService) waitForCreation(ctx context.Context, projectID, instanceID, jobName string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID, jobName)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, false, nil
			}
			return nil, false, err
		}
		return s, false, nil
	}
}

func (svc *JobsService) buildCreateRequest(job Job) ([]byte, error) {
	return json.Marshal(job)
}

// Update updates an argus scrape job
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_scrapeconfigs_create
func (svc *JobsService) Update(ctx context.Context, projectID, instanceID string, job Job) (res ListJobsResponse, err error) {
	if err = job.Validate(); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildCreateRequest(job)
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPathJob, projectID, instanceID, job.JobName), body)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Delete deletes an argus scrape job and returns the list of jobs, a wait handler and error
// the wait service is triggered with `.Wait()` and returns nil and error if occurred
// See also https://api.stackit.schwarz/argus-monitoring-service/openapi.v1.html#operation/v1_projects_instances_scrapeconfigs_delete_configs
func (svc *JobsService) Delete(ctx context.Context, projectID, instanceID, jobName string) (res ListJobsResponse, w *wait.Handler, err error) {
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPathJob, projectID, instanceID, jobName), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	w = wait.New(svc.waitForDeletion(ctx, projectID, instanceID, jobName))
	w.SetTimeout(10 * time.Minute)
	return
}

func (svc *JobsService) waitForDeletion(ctx context.Context, projectID, instanceID, jobName string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		if _, err = svc.Get(ctx, projectID, instanceID, jobName); err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	}
}
