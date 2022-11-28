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
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPathList   = consts.API_PATH_MONGO_DB_FLEX_INSTANCES
	apiPathCreate = consts.API_PATH_MONGO_DB_FLEX_INSTANCES
	apiPathGet    = consts.API_PATH_MONGO_DB_FLEX_INSTANCE
	apiPathUpdate = consts.API_PATH_MONGO_DB_FLEX_INSTANCE
)

func TestMongoDBInstancesService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"

	mux.HandleFunc(fmt.Sprintf(apiPathList, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"count": 1,
			"items": [
			  {
				"id": "string",
				"name": "string",
				"projectId": "abc"
			  }
			]
		  }`)
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
			Count: 1,
			Items: []instances.ListResponseItem{
				{
					ID:        "string",
					Name:      "string",
					ProjectID: projectID,
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mongo.Instances.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDBInstancesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MongoDBInstancesService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const instance_details_response = `{
	"item": {
	  "acl": {
		"items": [
		  "string"
		]
	  },
	  "backupSchedule": "string",
	  "flavor": {
		"cpu": 0,
		"description": "string",
		"id": "string",
		"memory": 0
	  },
	  "id": "def",
	  "name": "string",
	  "projectId": "abc",
	  "replicas": 0,
	  "status": "string",
	  "storage": {
		"class": "string",
		"size": 0
	  },
	  "users": [
		{
		  "database": "string",
		  "hostname": "string",
		  "id": "string",
		  "password": "string",
		  "port": 0,
		  "roles": [
			"string"
		  ],
		  "uri": "string",
		  "username": "string"
		}
	  ],
	  "version": "string"
	}
  }`

func buildInstance(projectID, instanceID string) instances.Instance {
	return instances.Instance{
		ACL: instances.ACL{
			Items: []string{"string"},
		},
		BackupSchedule: "string",
		Flavor: instances.Flavor{
			CPU:         0,
			Description: "string",
			ID:          "string",
			Memory:      0,
		},
		ID:        instanceID,
		Name:      "string",
		ProjectID: projectID,
		Replicas:  0,
		Status:    "string",
		Storage: instances.Storage{
			Class: "string",
			Size:  0,
		},
		Users: []instances.User{
			{
				Database: "string",
				Hostname: "string",
				ID:       "string",
				Password: "string",
				Port:     0,
				Roles:    []string{"string"},
				URI:      "string",
				Username: "string",
			},
		},
		Version: "string",
	}
}

func TestMongoDBInstancesService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"
	instanceID := "def"

	mux.HandleFunc(fmt.Sprintf(apiPathGet, projectID, instanceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, instance_details_response)
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
		{"ok", args{context.Background(), projectID, instanceID}, instances.GetResponse{
			Item: buildInstance(projectID, instanceID),
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mongo.Instances.Get(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDBInstancesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MongoDBInstancesService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMongoDBInstancesService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"
	instaceID := "def"
	instanceName := "string"

	instance := buildInstance(projectID, instaceID)

	mux.HandleFunc(fmt.Sprintf(apiPathCreate, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": "def"
		  }`)
	})

	type args struct {
		ctx            context.Context
		projectID      string
		instanceName   string
		flavorID       string
		storage        instances.Storage
		version        string
		replicas       int
		backupSchedule string
		labels         map[string]string
		options        map[string]string
		acl            instances.ACL
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
			context.Background(), projectID, instanceName,
			instance.Flavor.ID, instance.Storage, instance.Version,
			instance.Replicas, instance.BackupSchedule, map[string]string{},
			map[string]string{}, instance.ACL,
		}, instances.CreateResponse{ID: instaceID}, false, true},
	}
	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotW, err := mongo.Instances.Create(tt.args.ctx, tt.args.projectID, tt.args.instanceName, tt.args.flavorID, tt.args.storage, tt.args.version, tt.args.replicas, tt.args.backupSchedule, tt.args.labels, tt.args.options, tt.args.acl)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDBInstancesService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("MongoDBInstancesService.Create() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.useWait {
				process = gotW
			}
		})
	}

	baseDuration := time.Millisecond * 200

	ctx0, td0 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td0()

	ctx1, td1 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 3*baseDuration)
	defer td2()

	ctx3, td3 := context.WithTimeout(context.Background(), 4*baseDuration)
	defer td3()

	mux.HandleFunc(fmt.Sprintf(apiPathGet, projectID, instaceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}

		if ctx0.Err() == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if ctx1.Err() == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if ctx2.Err() == nil {
			fmt.Fprint(w, `{"item":{"status":"PROCESSING"}}`)
			return
		}
		if ctx3.Err() == nil {
			fmt.Fprint(w, `{"item":{"status":"FAILED"}}`)
			return
		}

		fmt.Fprint(w, `{"item":{"status":"READY"}}`)
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
		if !strings.Contains(err.Error(), "received status FAILED from server") {
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

func TestMongoDBInstancesService_Update(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"
	instaceID := "def"

	instance := buildInstance(projectID, instaceID)

	mux.HandleFunc(fmt.Sprintf(apiPathUpdate, projectID, instaceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, instance_details_response)
	})

	type args struct {
		ctx            context.Context
		projectID      string
		instanceID     string
		name           string
		flavorID       string
		storage        instances.Storage
		varsion        string
		replicas       int
		backupSchedule string
		labels         map[string]string
		options        map[string]string
		acl            instances.ACL
	}
	tests := []struct {
		name    string
		args    args
		wantRes instances.UpdateResponse
		wantErr bool
	}{
		{"nil ctx", args{ctx: nil}, instances.UpdateResponse{}, true},
		{"ok", args{
			context.Background(), projectID, instaceID, "string",
			instance.Flavor.ID, instances.Storage{Class: "string", Size: 0}, "string", 0, instance.BackupSchedule, map[string]string{},
			map[string]string{}, instance.ACL,
		}, instances.UpdateResponse{Item: instance}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, _, err := mongo.Instances.Update(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.name, tt.args.flavorID, tt.args.storage, tt.args.varsion, tt.args.replicas, tt.args.backupSchedule, tt.args.labels, tt.args.options, tt.args.acl)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDBInstancesService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("MongoDBInstancesService.Update() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMongoDBInstancesService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"
	instaceID := "def"

	baseDuration := time.Millisecond * 200

	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td2()

	mux.HandleFunc(fmt.Sprintf(apiPathUpdate, projectID, instaceID), func(w http.ResponseWriter, r *http.Request) {
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
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if ctx2.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{}`)
				return
			}
			w.WriteHeader(http.StatusNotFound)
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
				t.Errorf("MongoDBInstancesService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.useWait {
				process = gotW
			}
		})
	}

	process.SetThrottle(baseDuration)
	if _, err := process.Wait(); err == nil {
		t.Error("expected first run to return an error. got nil instead")
		return
	}

	time.Sleep(baseDuration)
	if _, err := process.Wait(); err != nil {
		t.Error(err)
	}

}
