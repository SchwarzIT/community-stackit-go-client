// package instances is used to manange MongoDB Flex instances

package instances_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	client "github.com/SchwarzIT/community-stackit-go-client"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPathList   = consts.API_PATH_DSA_INSTANCES
	apiPathCreate = consts.API_PATH_DSA_INSTANCES
	apiPathGet    = consts.API_PATH_DSA_INSTANCE
	apiPathUpdate = consts.API_PATH_DSA_INSTANCE
	broker        = "example"
)

const (
	list_response = `{
		"instances": [
		  {
			"instanceId": "string",
			"name": "string",
			"planId": "string",
			"dashboardUrl": "string",
			"cfGuid": "string",
			"cfSpaceGuid": "string",
			"organizationGuid": "string",
			"imageUrl": "string",
			"parameters": {},
			"lastOperation": {
			  "type": "create",
			  "state": "in progress",
			  "description": "string"
			}
		  }
		]
	  }`
)

func TestDSAInstancesService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := dataservices.New(c, broker)

	projectID := "abc"

	mux.HandleFunc(fmt.Sprintf(apiPathList, broker, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, list_response)
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.ListResponse
		wantErr bool
	}{
		{"nil ctx", args{ctx: nil}, instances.ListResponse{}, true},
		{"ok", args{context.Background(), projectID}, instances.ListResponse{
			Instances: []instances.Instance{
				{
					InstanceID:       "string",
					Name:             "string",
					PlanID:           "string",
					DashboardURL:     "string",
					CFGUID:           "string",
					CFSpaceGUID:      "string",
					OrganizationGUID: "string",
					ImageURL:         "string",
					Parameters:       map[string]string{},
					LastOperation: instances.LastOperation{
						Type:        "create",
						State:       "in progress",
						Description: "string",
					},
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mongo.Instances.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSAInstancesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("DSAInstancesService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const get_response = `{
	"instanceId": "string",
	"name": "string",
	"planId": "string",
	"dashboardUrl": "string",
	"cfGuid": "string",
	"cfSpaceGuid": "string",
	"organizationGuid": "string",
	"imageUrl": "string",
	"parameters": {},
	"lastOperation": {
	  "type": "create",
	  "state": "in progress",
	  "description": "string"
	}
  }`

func buildInstance(projectID, instanceID string) instances.Instance {
	return instances.Instance{}
}

func TestDSAInstancesService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := dataservices.New(c, broker)

	projectID := "abc"
	instanceID := "string"

	mux.HandleFunc(fmt.Sprintf(apiPathGet, broker, projectID, instanceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, get_response)
	})

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.GetResponse
		wantErr bool
	}{
		{"nil ctx", args{ctx: nil}, instances.GetResponse{}, true},
		{"ok", args{context.Background(), projectID, "string"}, instances.GetResponse{
			InstanceID:       "string",
			Name:             "string",
			PlanID:           "string",
			DashboardURL:     "string",
			CFGUID:           "string",
			CFSpaceGUID:      "string",
			OrganizationGUID: "string",
			ImageURL:         "string",
			Parameters:       map[string]string{},
			LastOperation: instances.LastOperation{
				Type:        "create",
				State:       "in progress",
				Description: "string",
			}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mongo.Instances.Get(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSAInstancesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("DSAInstancesService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestDSAInstancesService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := dataservices.New(c, broker)

	projectID := "abc"

	mux.HandleFunc(fmt.Sprintf(apiPathCreate, broker, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"instanceId": "string"
		  }`)
	})

	type args struct {
		ctx          context.Context
		projectID    string
		instanceName string
		planID       string
		parametes    map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.CreateResponse
		wantErr bool
		useWait bool
	}{
		{"nil ctx", args{ctx: nil}, instances.CreateResponse{}, true, false},
		{"ok", args{
			context.Background(), projectID, "string", "string", map[string]string{}}, instances.CreateResponse{InstanceID: "string"}, false, true},
	}
	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotW, err := mongo.Instances.Create(tt.args.ctx, tt.args.projectID, tt.args.instanceName, tt.args.planID, tt.args.parametes)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSAInstancesService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("DSAInstancesService.Create() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.useWait {
				process = gotW
			}
		})
	}

	baseDuration := time.Millisecond * 200

	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td2()

	ctx3, td3 := context.WithTimeout(context.Background(), 3*baseDuration)
	defer td3()

	mux.HandleFunc(fmt.Sprintf(apiPathGet, broker, projectID, "string"), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}

		if ctx1.Err() == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if ctx2.Err() == nil {
			fmt.Fprint(w, `{"lastOperation":{"state":"in progress"}}`)
			return
		}
		if ctx3.Err() == nil {
			fmt.Fprint(w, `{"lastOperation":{"state":"failed"}}`)
			return
		}

		fmt.Fprint(w, `{"lastOperation":{"state":"succeeded"}}`)
	})

	process.SetThrottle(baseDuration)
	if _, err := process.Wait(); err == nil {
		t.Error("expected first call to return an error. got nil instead")
		return
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err == nil {
		t.Error("expected 2nd call to return an error. got nil instead")
		return
	} else {
		if !strings.Contains(err.Error(), "received failed status from DSA instance") {
			t.Errorf("unexpected error: %v", err)
			return
		}
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err != nil {
		t.Error(err)
		return
	}
}

func TestDSAInstancesService_Update(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := dataservices.New(c, broker)

	projectID := "abc"
	instaceID := "string"

	instance := buildInstance(projectID, instaceID)

	baseDuration := time.Millisecond * 200

	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td2()

	ctx3, td3 := context.WithTimeout(context.Background(), 3*baseDuration)
	defer td3()

	ctx4, td4 := context.WithTimeout(context.Background(), 4*baseDuration)
	defer td4()

	mux.HandleFunc(fmt.Sprintf(apiPathUpdate, broker, projectID, instaceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{
			"error": "string",
			"description": "string"
		  }`)
			return
		}

		if r.Method == http.MethodGet {
			if ctx1.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"in progress", "type":"create"}}`)
				return
			}
			if ctx2.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"in progress", "type":"update"}}`)
				return
			}
			if ctx3.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"succeeded", "type":"update"}}`)
				return
			}
			if ctx4.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"failed", "type":"update"}}`)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Error("wrong method")
	})

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		planID     string
		parametes  map[string]string
	}
	tests := []struct {
		name     string
		args     args
		wantRes  instances.UpdateResponse
		wantErr  bool
		wantWait bool
	}{
		{"nil ctx", args{ctx: nil}, instances.UpdateResponse{}, true, false},
		{"ok", args{context.Background(), projectID, instaceID, instance.PlanID, instance.Parameters}, instances.UpdateResponse{
			Error:       "string",
			Description: "string",
		}, false, true},
	}

	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, w, err := mongo.Instances.Update(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.planID, tt.args.parametes)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSAInstancesService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("DSAInstancesService.Update() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.wantWait {
				process = w
			}
		})
	}

	process.SetThrottle(baseDuration)
	if _, err := process.Wait(); err != nil {
		t.Error(err)
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err == nil {
		t.Error("expected 2nd run to return an error. got nil instead")
		return
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err == nil {
		t.Error("expected last run to return an error. got nil instead")
		return
	}
}

func TestDSAInstancesService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := dataservices.New(c, broker)

	projectID := "abc"
	instaceID := "def"

	baseDuration := time.Millisecond * 200

	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td2()

	ctx3, td3 := context.WithTimeout(context.Background(), 3*baseDuration)
	defer td3()

	ctx4, td4 := context.WithTimeout(context.Background(), 4*baseDuration)
	defer td4()

	ctx5, td5 := context.WithTimeout(context.Background(), 5*baseDuration)
	defer td5()

	mux.HandleFunc(fmt.Sprintf(apiPathUpdate, broker, projectID, instaceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{
				"executionTime": "string"
			  }`)
			return
		}

		if r.Method == http.MethodGet {
			if ctx1.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"in progress", "type":"create"}}`)
				return
			}
			if ctx2.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"in progress", "type":"delete"}}`)
				return
			}
			if ctx3.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"succeeded", "type":"delete"}}`)
				return
			}
			if ctx4.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lastOperation":{"state":"failed", "type":"delete"}}`)
				return
			}
			if ctx5.Err() == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Error("wrong method")
	})

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		useWait bool
	}{
		{"nil ctx", args{ctx: nil}, true, false},
		{"ok", args{context.Background(), projectID, instaceID}, false, true},
	}
	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotW, err := mongo.Instances.Delete(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSAInstancesService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.useWait {
				process = gotW
			}
		})
	}

	process.SetThrottle(baseDuration)
	if _, err := process.Wait(); err != nil {
		t.Error(err)
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err == nil {
		t.Error("expected 2nd run to return an error. got nil instead")
		return
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err != nil {
		t.Error(err)
		return
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err == nil {
		t.Error("expected last run to return an error. got nil instead")
		return
	}

}
