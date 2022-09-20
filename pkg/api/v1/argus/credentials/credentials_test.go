package credentials_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/credentials"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath = consts.API_PATH_ARGUS_CREDENTIALS
)

func prep(t *testing.T, path, projectID, instanceID, want, method string) (*credentials.CredentialsService, func()) {
	c, mux, teardown, _ := client.MockServer()
	a := credentials.New(c)

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

func TestInstanceCredentialsService_List(t *testing.T) {
	pid := "1234"
	iid := "1234"
	c, teardown := prep(t, "", pid, iid, list_response, http.MethodGet)
	defer teardown()

	var want credentials.CredentialList
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
		wantRes credentials.CredentialList
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid}, want, false, true},
		{"nil ctx", args{nil, pid, iid}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := c.List(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstanceCredentialsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("InstanceCredentialsService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	list_response = `{
		"message": "Successfully received all technical users",
		"credentials": [
		  {
			"name": "test",
			"id": "test",
			"credentialsInfo": {
			  "username": "test"
			}
		  }
		]
	  }`
)

func TestCredentialsService_Get(t *testing.T) {
	pid := "1234"
	iid := "1234"
	username := "testing_9449de83-64ac-45dc-9967-e7c75bbdca70_4d92d3d9-d5c2-4c0b-98ad-950878101d9e"
	c, teardown := prep(t, fmt.Sprintf("/%s", username), pid, iid, get_response, http.MethodGet)
	defer teardown()

	var want credentials.Credential
	if err := json.Unmarshal([]byte(get_response), &want); err != nil {
		t.Error(err)
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		username   string
	}
	tests := []struct {
		name    string
		args    args
		wantRes credentials.Credential
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid, username}, want, false, true},
		{"nil ctx", args{nil, pid, iid, username}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := c.Get(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("CredentialsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("CredentialsService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	get_response = `{
		"message": "The technical credentials were successfully deleted",
		"credentialsInfo": {
		  "username": "testing_9449de83-64ac-45dc-9967-e7c75bbdca70_4d92d3d9-d5c2-4c0b-98ad-950878101d9e"
		},
		"name": "testing_9449de83-64ac-45dc-9967-e7c75bbdca70_4d92d3d9-d5c2-4c0b-98ad-950878101d9e",
		"id": "testing_9449de83-64ac-45dc-9967-e7c75bbdca70_4d92d3d9-d5c2-4c0b-98ad-950878101d9e"
	  }`
)

func TestCredentialsService_Create(t *testing.T) {
	pid := "1234"
	iid := "1234"
	c, teardown := prep(t, "", pid, iid, create_response, http.MethodPost)
	defer teardown()

	var want credentials.CreateResponse
	if err := json.Unmarshal([]byte(create_response), &want); err != nil {
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
		wantRes credentials.CreateResponse
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid}, want, false, true},
		{"nil ctx", args{nil, pid, iid}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := c.Create(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CredentialsService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("CredentialsService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	create_response = `{
		"message": "Successfully created api user",
		"credentials": {
		  "username": "test",
		  "password": "1fasAELDB234ddeDAfdasfel787oplpj"
		}
	  }`
)

func TestCredentialsService_Delete(t *testing.T) {
	pid := "1234"
	iid := "1234"
	username := "testing_9449de83-64ac-45dc-9967-e7c75bbdca70_4d92d3d9-d5c2-4c0b-98ad-950878101d9e"
	c, teardown := prep(t, fmt.Sprintf("/%s", username), pid, iid, delete_response, http.MethodDelete)
	defer teardown()

	var want credentials.Credential
	if err := json.Unmarshal([]byte(delete_response), &want); err != nil {
		t.Error(err)
	}

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		username   string
	}
	tests := []struct {
		name    string
		args    args
		wantRes credentials.Credential
		wantErr bool
		compare bool
	}{
		{"all ok", args{context.Background(), pid, iid, username}, want, false, true},
		{"nil ctx", args{nil, pid, iid, username}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := c.Delete(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("CredentialsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.compare {
				t.Errorf("CredentialsService.Delete() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	delete_response = `{
		"message": "The technical credentials were successfully deleted"
	  }`
)
