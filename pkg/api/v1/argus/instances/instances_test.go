package instances_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath               = consts.API_PATH_ARGUS_INSTANCES
	apiPathWithInstanceID = "/%s"
)

func prep(t *testing.T, path, projectID string, want []byte, method string) (*argus.ArgusService, func()) {
	c, mux, teardown, _ := client.MockServer()
	a := argus.New(c)

	mux.HandleFunc(fmt.Sprintf(apiPath+path, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(want))
	})

	return a, teardown
}

func TestInstancesService_List(t *testing.T) {
	projectID := "123"
	want := []byte(`{
		"message": "Successfully got instances",
		"instances": [
		  {
			"id": "test-awesome",
			"planName": "Observability-Basic-EU01",
			"instance": "9449de83-64ac-45dc-9967-e7c75bbdca70",
			"name": "testing",
			"status": "CREATE_SUCCEEDED",
			"serviceName": "STACKIT Argus"
		  }
		]
	  }`)

	var wantRes instances.InstanceList
	if err := json.Unmarshal(want, &wantRes); err != nil {
		t.Error(err)
	}
	svc, teardown := prep(t, "", projectID, want, http.MethodGet)
	defer teardown()

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.InstanceList
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), projectID}, wantRes, false, true},
		{"project not found", args{context.Background(), "random"}, wantRes, true, false},
		{"nil ctx", args{nil, projectID}, wantRes, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svc.Instances.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("InstancesService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestInstancesService_Get(t *testing.T) {
	projectID := "123"
	instanceID := "9449de83-64ac-45dc-9967-e7c75bbdca70"
	want := []byte(`{
		"message": "Successfully got instance",
		"dashboardUrl": "https://portal-dev.stackit.cloud/projects/775eee9d-8565-48ab-9dcc-64a6ca55043a/service/9449de83-64ac-45dc-9967-e7c75bbdca70/argus-dashboard/instances/9449de83-64ac-45dc-9967-e7c75bbdca70/overview",
		"isUpdatable": true,
		"name": "testing",
		"parameters": {},
		"id": "9449de83-64ac-45dc-9967-e7c75bbdca70",
		"serviceName": "STACKIT Argus",
		"planId": "a9d5b2df-82dd-40d6-91e9-d551f2de3dda",
		"planName": "Observability-Basic-EU01",
		"planSchema": "{}",
		"status": "CREATE_SUCCEEDED",
		"instance": {
		  "instance": "9449de83-64ac-45dc-9967-e7c75bbdca70",
		  "cluster": "stackit",
		  "grafanaUrl": "https://ui.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70",
		  "dashboardUrl": "dashboard.example.com",
		  "grafanaPlugins": [],
		  "name": "test",
		  "grafanaAdminPassword": "asdf313kles23450des0asdf313kles2",
		  "grafanaAdminUser": "admin",
		  "metricsRetentionTimeRaw": 14,
		  "MetricsRetentionTime5m": 0,
		  "MetricsRetentionTime1h": 0,
		  "metricsUrl": "https://storage.api.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70",
		  "pushMetricsUrl": "https://push.metrics.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70/api/v1/receive",
		  "grafanaPublicReadAccess": false,
		  "targetsUrl": "https://metrics.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70",
		  "alertingUrl": "https://alerting.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70",
		  "plan": {
			"planId": "7b1fbd9c-9acd-42ce-95d3-0f6822d6cabe",
			"id": "7b1fbd9c-9acd-42ce-95d3-0f6822d6cabe",
			"description": "Small Plan",
			"name": "SmallPlan",
			"bucketSize": 20,
			"grafanaGlobalUsers": 10,
			"grafanaGlobalOrgs": 2,
			"grafanaGlobalDashboards": 20,
			"alertRules": 1000,
			"targetNumber": 2,
			"samplesPerScrape": 1000,
			"grafanaGlobalSessions": 10,
			"amount": 49,
			"alertReceivers": 10,
			"alertMatchers": 10,
			"tracesStorage": 20,
			"logsStorage": 20,
			"logsAlert": 20,
			"isFree": false,
			"isPublic": true,
			"schema": ""
		  },
		  "logsUrl": "https://logs.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70",
		  "logsPushUrl": "https://logs.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70/loki/api/v1/push",
		  "jaegerTracesUrl": "9449de83-64ac-gj.traces.stackit.argus.eu01.cloud:443",
		  "otlpTracesUrl": "9449de83-64ac-op.traces.stackit.argus.eu01.cloud:443",
		  "zipkinSpansUrl": "https://9449de83-64ac-zk.traces.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70",
		  "jaegerUiUrl": "https://9449de83-64ac-jui.traces.stackit.argus.eu01.cloud/instances/9449de83-64ac-45dc-9967-e7c75bbdca70"
		}
	  }`)

	var wantRes instances.Instance
	if err := json.Unmarshal(want, &wantRes); err != nil {
		t.Error(err)
	}
	svc, teardown := prep(t, fmt.Sprintf(apiPathWithInstanceID, instanceID), projectID, want, http.MethodGet)
	defer teardown()

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.Instance
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), projectID, instanceID}, wantRes, false, true},
		{"project not found", args{context.Background(), "random", instanceID}, wantRes, true, false},
		{"nil ctx", args{nil, projectID, instanceID}, wantRes, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svc.Instances.Get(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("InstancesService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestInstancesService_Create(t *testing.T) {
	projectID := "597976c4-d4c1-44d6-9f43-213df3da1799"
	projectID2 := "697976c4-d4c1-44d6-9f43-213df3da1799"
	want := []byte(`{
		"message": "Successfully created instance",
		"instanceId": "597976c4-d4c1-44d6-9f43-213df3da1799",
		"dashboardUrl": "https://portal-dev.stackit.cloud/projects/775eee9d-8565-48ab-9dcc-64a6ca55043a/service/597976c4-d4c1-44d6-9f43-213df3da1799/argus-dashboard/instances/597976c4-d4c1-44d6-9f43-213df3da1799/overview"
	  }`)

	wantRes := instances.CreateOrUpdateResponse{
		Message:      "Successfully created instance",
		InstanceID:   "597976c4-d4c1-44d6-9f43-213df3da1799",
		DashboardURL: "https://portal-dev.stackit.cloud/projects/775eee9d-8565-48ab-9dcc-64a6ca55043a/service/597976c4-d4c1-44d6-9f43-213df3da1799/argus-dashboard/instances/597976c4-d4c1-44d6-9f43-213df3da1799/overview",
	}

	svc, teardown := prep(t, "", projectID, want, http.MethodPost)
	defer teardown()

	type args struct {
		ctx          context.Context
		projectID    string
		instanceName string
		planID       string
		params       map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.CreateOrUpdateResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), projectID, "name-123", "plan-123", map[string]string{}}, wantRes, false, true},
		{"nil ctx", args{nil, projectID, "name-123", "plan-123", map[string]string{}}, wantRes, true, false},
		{"wrong project uuid", args{context.Background(), "random", "name-123", "plan-123", map[string]string{}}, wantRes, true, false},
		{"project not found", args{context.Background(), projectID2, "name-123", "plan-123", map[string]string{}}, wantRes, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svc.Instances.Create(tt.args.ctx, tt.args.projectID, tt.args.instanceName, tt.args.planID, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("InstancesService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestInstancesService_Update(t *testing.T) {
	projectID := "597976c4-d4c1-44d6-9f43-213df3da1799"
	instanceID := "597976c4-d4c1-44d6-9f43-213df3da1799"
	projectID2 := "697976c4-d4c1-44d6-9f43-213df3da1799"
	want := []byte(`{
		"message": "Successfully created instance",
		"instanceId": "597976c4-d4c1-44d6-9f43-213df3da1799",
		"dashboardUrl": "https://portal-dev.stackit.cloud/projects/775eee9d-8565-48ab-9dcc-64a6ca55043a/service/597976c4-d4c1-44d6-9f43-213df3da1799/argus-dashboard/instances/597976c4-d4c1-44d6-9f43-213df3da1799/overview"
	  }`)

	wantRes := instances.CreateOrUpdateResponse{
		Message:      "Successfully created instance",
		InstanceID:   "597976c4-d4c1-44d6-9f43-213df3da1799",
		DashboardURL: "https://portal-dev.stackit.cloud/projects/775eee9d-8565-48ab-9dcc-64a6ca55043a/service/597976c4-d4c1-44d6-9f43-213df3da1799/argus-dashboard/instances/597976c4-d4c1-44d6-9f43-213df3da1799/overview",
	}

	svc, teardown := prep(t, fmt.Sprintf(apiPathWithInstanceID, instanceID), projectID, want, http.MethodPut)
	defer teardown()

	type args struct {
		ctx          context.Context
		projectID    string
		instanceID   string
		instanceName string
		planID       string
		params       map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.CreateOrUpdateResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), projectID, instanceID, "name-123", "plan-123", map[string]string{}}, wantRes, false, true},
		{"nil ctx", args{nil, projectID, instanceID, "name-123", "plan-123", map[string]string{}}, wantRes, true, false},
		{"wrong project uuid", args{context.Background(), "random", instanceID, "name-123", "plan-123", map[string]string{}}, wantRes, true, false},
		{"bad name", args{context.Background(), projectID, instanceID, "", "plan-123", map[string]string{}}, wantRes, true, false},
		{"bad plan", args{context.Background(), projectID, instanceID, "name-123", "", map[string]string{}}, wantRes, true, false},
		{"project not found", args{context.Background(), projectID2, instanceID, "name-123", "plan-123", map[string]string{}}, wantRes, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svc.Instances.Update(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.instanceName, tt.args.planID, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("InstancesService.Update() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestInstancesService_Delete(t *testing.T) {
	projectID := "597976c4-d4c1-44d6-9f43-213df3da1799"
	instanceID := "597976c4-d4c1-44d6-9f43-213df3da1799"
	projectID2 := "697976c4-d4c1-44d6-9f43-213df3da1799"
	want := []byte(`{
		"message": "Successfully deleted instance"
	  }`)

	wantRes := instances.Instance{
		Message: "Successfully deleted instance",
	}

	svc, teardown := prep(t, fmt.Sprintf(apiPathWithInstanceID, instanceID), projectID, want, http.MethodDelete)
	defer teardown()

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.Instance
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), projectID, instanceID}, wantRes, false, true},
		{"nil ctx", args{nil, projectID, instanceID}, wantRes, true, false},
		{"wrong project uuid", args{context.Background(), "random", instanceID}, wantRes, true, false},
		{"project not found", args{context.Background(), projectID2, instanceID}, wantRes, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := svc.Instances.Delete(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("InstancesService.Delete() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
