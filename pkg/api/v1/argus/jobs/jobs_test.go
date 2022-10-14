package jobs_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/jobs"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPath    = consts.API_PATH_ARGUS_JOBS
	apiPathJob = "/%s"
)

func prep(t *testing.T, path, projectID, instanceID, want, method string) (*jobs.JobsService, *http.ServeMux, func()) {
	c, mux, teardown, _ := client.MockServer()
	a := jobs.New(c)

	mux.HandleFunc(fmt.Sprintf(apiPath+path, projectID, instanceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(want))
	})

	return a, mux, teardown
}

func TestJobsService_List(t *testing.T) {
	pid := "1234"
	iid := "5678"
	a, _, teardown := prep(t, "", pid, iid, list_response, http.MethodGet)
	defer teardown()

	var want jobs.ListJobsResponse
	if err := json.Unmarshal([]byte(list_response), &want); err != nil {
		t.Error(err)
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes jobs.ListJobsResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid}, want, false, true},
		{"nil ctx", args{nil, pid, iid}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := a.List(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("JobsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("JobsService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	list_response = `{
		"message": "Successfully got scrape config",
		"data": [
		  {
			"jobName": "test",
			"scheme": "https",
			"scrapeInterval": "5m",
			"scrapeTimeout": "1m",
			"staticConfigs": [
			  {
				"targets": [
				  "example.com"
				]
			  }
			],
			"metricsPath": "/metrics"
		  }
		]
	  }`
)

func TestJobsService_Get(t *testing.T) {
	pid := "1234"
	iid := "5678"
	name := "job-name-1"
	a, _, teardown := prep(t, fmt.Sprintf(apiPathJob, name), pid, iid, get_response, http.MethodGet)
	defer teardown()

	var want jobs.GetJobResponse
	if err := json.Unmarshal([]byte(get_response), &want); err != nil {
		t.Error(err)
	}
	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		jobName    string
	}
	tests := []struct {
		name    string
		args    args
		wantRes jobs.GetJobResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid, name}, want, false, true},
		{"nil ctx", args{nil, pid, iid, name}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := a.Get(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.jobName)
			if (err != nil) != tt.wantErr {
				t.Errorf("JobsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("JobsService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const get_response = `{
	"message": "Successfully got scrape target",
	"data": {
	  "jobName": "test",
	  "scheme": "https",
	  "scrapeInterval": "5m",
	  "scrapeTimeout": "1m",
	  "staticConfigs": [
		{
		  "targets": [
			"example.com"
		  ]
		}
	  ],
	  "metricsPath": "/metrics"
	}
  }`

func TestJobsService_Create(t *testing.T) {
	pid := "1234"
	iid := "5678"
	a, mux, teardown := prep(t, "", pid, iid, create_response, http.MethodPost)
	defer teardown()

	baseDuration := 200 * time.Millisecond
	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td2()

	mux.HandleFunc(fmt.Sprintf(apiPath+apiPathJob, pid, iid, "my-job"), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if ctx1.Err() == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if ctx2.Err() == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, get_response)
	})

	var want jobs.ListJobsResponse
	if err := json.Unmarshal([]byte(create_response), &want); err != nil {
		t.Error(err)
	}

	job := jobs.Job{
		StaticConfigs:  []jobs.StaticConfig{{Targets: []string{"abc"}}},
		JobName:        "my-job",
		Scheme:         "http",
		ScrapeInterval: "1m",
		ScrapeTimeout:  "5s",
		MetricsPath:    "/",
	}

	job2 := jobs.Job{
		StaticConfigs:  []jobs.StaticConfig{{Targets: []string{"abc"}}},
		JobName:        "my_job",
		Scheme:         "http",
		ScrapeInterval: "1m",
		ScrapeTimeout:  "5s",
		MetricsPath:    "/",
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		job        jobs.Job
	}
	tests := []struct {
		name    string
		args    args
		wantRes jobs.ListJobsResponse
		wantErr bool
		compare bool
		useWait bool
	}{
		{"all ok", args{context.Background(), pid, iid, job}, want, false, true, true},
		{"bad job", args{context.Background(), pid, iid, job2}, want, true, false, false},
		{"nil ctx", args{nil, pid, iid, job}, want, true, false, false},
	}

	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, w, err := a.Create(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("JobsService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("JobsService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.useWait {
				process = w
			}
		})
	}

	process.SetThrottle(baseDuration)
	// first test run: expect Wait to exit with an error
	if _, err := process.Wait(); err == nil {
		t.Error("expected error but got nil")
	}

	time.Sleep(baseDuration)
	// 2nd test: expect Not found, and exit with success after retry
	if _, err := process.Wait(); err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}

const create_response = `{
	"message": "Scrape target successfully created",
	"data": [
	  {
		"jobName": "test",
		"scheme": "https",
		"scrapeInterval": "5m",
		"scrapeTimeout": "1m",
		"staticConfigs": [
		  {
			"targets": [
			  "example.com"
			]
		  }
		],
		"metricsPath": "/metrics"
	  }
	]
  }`

func TestJobsService_Update(t *testing.T) {
	pid := "1234"
	iid := "5678"
	a, _, teardown := prep(t, "/my-job", pid, iid, create_response, http.MethodPut)
	defer teardown()

	var want jobs.ListJobsResponse
	if err := json.Unmarshal([]byte(create_response), &want); err != nil {
		t.Error(err)
	}

	job := jobs.Job{
		StaticConfigs:  []jobs.StaticConfig{{Targets: []string{"abc"}}},
		JobName:        "my-job",
		Scheme:         "http",
		ScrapeInterval: "1m",
		ScrapeTimeout:  "5s",
		MetricsPath:    "/",
	}

	job2 := jobs.Job{
		StaticConfigs:  []jobs.StaticConfig{{Targets: []string{"abc"}}},
		JobName:        "my_job",
		Scheme:         "http",
		ScrapeInterval: "1m",
		ScrapeTimeout:  "5s",
		MetricsPath:    "/",
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		job        jobs.Job
	}
	tests := []struct {
		name    string
		args    args
		wantRes jobs.ListJobsResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid, job}, want, false, true},
		{"bad job", args{context.Background(), pid, iid, job2}, want, true, false},
		{"nil ctx", args{nil, pid, iid, job}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := a.Update(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("JobsService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("JobsService.Update() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestJobsService_Delete(t *testing.T) {
	pid := "1234"
	iid := "5678"
	jid := "my-job"

	c, mux, teardown, _ := client.MockServer()
	a := jobs.New(c)
	defer teardown()

	baseDuration := 200 * time.Millisecond
	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td2()

	mux.HandleFunc(fmt.Sprintf(apiPath+apiPathJob, pid, iid, jid), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, delete_response)
			return
		}

		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}

		w.Header().Set("Content-Type", "application/json")
		if ctx1.Err() == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if ctx2.Err() == nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, get_response)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})

	var want jobs.ListJobsResponse
	if err := json.Unmarshal([]byte(delete_response), &want); err != nil {
		t.Error(err)
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		jobName    string
	}
	tests := []struct {
		name    string
		args    args
		wantRes jobs.ListJobsResponse
		wantErr bool
		compare bool
		useWait bool
	}{
		{"all ok", args{context.Background(), pid, iid, jid}, want, false, true, true},
		{"bad job", args{context.Background(), pid, iid, "something"}, want, true, false, false},
		{"nil ctx", args{nil, pid, iid, jid}, want, true, false, false},
	}
	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, w, err := a.Delete(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.jobName)
			if (err != nil) != tt.wantErr {
				t.Errorf("JobsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("JobsService.Delete() = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.useWait {
				process = w
			}
		})
	}

	process.SetThrottle(baseDuration)
	// first test run: expect Wait to exit with an error
	if _, err := process.Wait(); err == nil {
		t.Error("expected error but got nil")
	}

	time.Sleep(baseDuration)
	// 2nd test: expect found, retrey & not found -> exist successfully
	if _, err := process.Wait(); err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}

const delete_response = `{
	"data": [
	  {
		"jobName": "test",
		"scheme": "https",
		"scrapeInterval": "5m",
		"scrapeTimeout": "1m",
		"staticConfigs": [
		  {
			"targets": [
			  "example.com"
			]
		  }
		],
		"metricsPath": "/metrics"
	  }
	],
	"message": "Job has been deleted successfully"
  }`
