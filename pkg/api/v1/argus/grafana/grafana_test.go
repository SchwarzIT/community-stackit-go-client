package grafana_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/grafana"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_ARGUS_GRAFANA_CONFIGS
)

func prep(t *testing.T, path, projectID, instanceID, want, method string) (*grafana.GrafanaService, func()) {
	c, mux, teardown, _ := client.MockServer()
	a := grafana.New(c)

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

func TestGrafanaService_GetConfig(t *testing.T) {
	pid := "1234"
	iid := "5678"
	a, teardown := prep(t, "", pid, iid, get_config_response, http.MethodGet)
	defer teardown()

	var want grafana.GetConfigResponse
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
		wantRes grafana.GetConfigResponse
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
				t.Errorf("GrafanaService.GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("GrafanaService.GetConfig() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	get_config_response = `{
		"message": "Successfully got grafana config",
		"publicReadAccess": false
	  }`
)

func TestGrafanaService_UpdateConfig(t *testing.T) {
	pid := "1234"
	iid := "5678"
	a, teardown := prep(t, "", pid, iid, update_response, http.MethodPut)
	defer teardown()

	var want grafana.UpdateConfigResponse
	if err := json.Unmarshal([]byte(update_response), &want); err != nil {
		t.Error(err)
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		config     grafana.Config
	}
	tests := []struct {
		name    string
		args    args
		wantRes grafana.UpdateConfigResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid, grafana.Config{}}, want, false, true},
		{"nil ctx", args{nil, pid, iid, grafana.Config{}}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := a.UpdateConfig(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrafanaService.UpdateConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("GrafanaService.UpdateConfig() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	update_response = `{
		"message": "Successfully updated grafana config"
	  }`
)
