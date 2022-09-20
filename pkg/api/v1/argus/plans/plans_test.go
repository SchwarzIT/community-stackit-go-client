package plans_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/plans"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

const (
	apiPath = consts.API_PATH_ARGUS_PLANS
)

func prep(t *testing.T, path, projectID string, want []byte) (*argus.ArgusService, plans.PlanList, func()) {
	c, mux, teardown, _ := client.MockServer()
	a := argus.New(c)

	mux.HandleFunc(fmt.Sprintf(apiPath+path, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(want))
	})

	var wantRes plans.PlanList
	if err := json.Unmarshal(want, &wantRes); err != nil {
		t.Error(err)
	}

	return a, wantRes, teardown
}

func TestPlansService_List(t *testing.T) {
	projectID := "1234"
	want := []byte(`{
		"message": "Successfully got plans",
		"plans": [
		  {
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
			"schema": "{}"
		  }
		]
	  }`)

	a, wantRes, teardown := prep(t, "", projectID, want)
	defer teardown()

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes plans.PlanList
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID}, wantRes, false},
		{"nil ctx", args{nil, projectID}, wantRes, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := a.Plans.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlansService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("PlansService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
