package metrics_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/metrics"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_ARGUS_METRICS_RETENTION
)

func prep(t *testing.T, path, projectID, instanceID, want, method string) (*metrics.MetricsService, func()) {
	c, mux, teardown, _ := client.MockServer()
	a := metrics.New(c)

	mux.HandleFunc(fmt.Sprintf(apiPath+path, projectID, instanceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(want))
	})

	return a, teardown
}

func TestMetricsService_GetConfig(t *testing.T) {
	pid := "1234"
	iid := "5678"
	a, teardown := prep(t, "", pid, iid, get_config_response, http.MethodGet)
	defer teardown()

	var want metrics.GetConfigResponse
	if err := json.Unmarshal([]byte(get_config_response), &want); err != nil {
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
		wantRes metrics.GetConfigResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid}, want, false, true},
		{"nil ctx", args{nil, pid, iid}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := a.GetConfig(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricsService.GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("MetricsService.GetConfig() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	get_config_response = `{
		"metricsRetentionTimeRaw": "string",
		"MetricsRetentionTime5m": "string",
		"MetricsRetentionTime1h": "string"
	  }`
)

func TestMetricsService_UpdateConfig(t *testing.T) {
	pid := "1234"
	iid := "5678"
	a, teardown := prep(t, "", pid, iid, update_response, http.MethodPut)
	defer teardown()

	var want metrics.UpdateConfigResponse
	if err := json.Unmarshal([]byte(update_response), &want); err != nil {
		t.Error(err)
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		config     metrics.Config
	}
	tests := []struct {
		name    string
		args    args
		wantRes metrics.UpdateConfigResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid, metrics.Config{"9360h", "720h", "24h"}}, want, false, true},
		{"missing data", args{context.Background(), pid, iid, metrics.Config{}}, want, true, false},
		{"nil ctx", args{nil, pid, iid, metrics.Config{"9360h", "720h", "24h"}}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := a.UpdateConfig(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricsService.UpdateConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("MetricsService.UpdateConfig() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	update_response = `{
		"message": "Successfully updated metric storage retention"
	  }`
)
