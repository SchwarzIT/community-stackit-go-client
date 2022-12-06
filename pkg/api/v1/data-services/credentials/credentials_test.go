// package credentials is used to manage DSA instance credentials

package credentials_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/credentials"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

const (
	apiPathList   = consts.API_PATH_DSA_CREDENTIALS
	apiPathCreate = consts.API_PATH_DSA_CREDENTIALS
	apiPathGet    = consts.API_PATH_DSA_CREDENTIAL
	apiPathDelete = consts.API_PATH_DSA_CREDENTIAL

	broker     = 0
	projectID  = "example"
	instanceID = "example"
	credID     = "example"
)

const (
	list_response = `{
		"credentialsList": [
		  {
			"id": "string"
		  }
		]
	  }`
)

func TestDSACredentialsService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	dsa := dataservices.New(c, broker, c.GetConfig().BaseUrl.String())

	mux.HandleFunc(fmt.Sprintf(apiPathList, projectID, instanceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, list_response)
	})

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes credentials.ListResponse
		wantErr bool
	}{
		{"nil ctx", args{ctx: nil}, credentials.ListResponse{}, true},
		{"ok", args{context.Background(), projectID, instanceID}, credentials.ListResponse{
			CredentialsList: []credentials.CredentialsListItem{
				{
					ID: "string",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := dsa.Credentials.List(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSACredentialsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("DSACredentialsService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	get_response = `{
		"id": "string",
		"uri": "string",
		"raw": {
		  "credentials": {
			"host": "string",
			"port": 0,
			"hosts": [
			  "string"
			],
			"username": "string",
			"password": "string",
			"cacrt": "string",
			"scheme": "string",
			"uri": "string"
		  },
		  "syslogDrainUrl": "string",
		  "routeServiceUrl": "string"
		}
	  }`
)

func TestDSACredentialsService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	dsa := dataservices.New(c, broker, c.GetConfig().BaseUrl.String())

	mux.HandleFunc(fmt.Sprintf(apiPathGet, projectID, instanceID, credID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, get_response)
	})

	type args struct {
		ctx          context.Context
		projectID    string
		instanceID   string
		credentialID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes credentials.GetResponse
		wantErr bool
	}{
		{"nil ctx", args{ctx: nil}, credentials.GetResponse{}, true},
		{"ok", args{context.Background(), projectID, instanceID, credID}, credentials.GetResponse{
			ID:  "string",
			URI: "string",
			Raw: credentials.RawCredential{
				Credential: credentials.Credential{
					Host:     "string",
					Hosts:    []string{"string"},
					Username: "string",
					Password: "string",
					Cacrt:    "string",
					Scheme:   "string",
					URI:      "string",
				},
				SyslogDrainURL:  "string",
				RouteServiceURL: "string",
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := dsa.Credentials.Get(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.credentialID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSACredentialsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("DSACredentialsService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestDSACredentialsService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	dsa := dataservices.New(c, broker, c.GetConfig().BaseUrl.String())

	mux.HandleFunc(fmt.Sprintf(apiPathCreate, projectID, instanceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
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
		wantRes credentials.CreateResponse
		wantErr bool
	}{
		{"nil ctx", args{ctx: nil}, credentials.CreateResponse{}, true},
		{"ok", args{context.Background(), projectID, instanceID}, credentials.CreateResponse{
			ID:  "string",
			URI: "string",
			Raw: credentials.RawCredential{
				Credential: credentials.Credential{
					Host:     "string",
					Hosts:    []string{"string"},
					Username: "string",
					Password: "string",
					Cacrt:    "string",
					Scheme:   "string",
					URI:      "string",
				},
				SyslogDrainURL:  "string",
				RouteServiceURL: "string",
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := dsa.Credentials.Create(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSACredentialsService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("DSACredentialsService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

const (
	delete_response = `{
		"error": "string",
		"description": "string"
	  }`
)

func TestDSACredentialsService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	dsa := dataservices.New(c, broker, c.GetConfig().BaseUrl.String())

	mux.HandleFunc(fmt.Sprintf(apiPathDelete, projectID, instanceID, credID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, delete_response)
	})
	type args struct {
		ctx          context.Context
		projectID    string
		instanceID   string
		credentialID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes credentials.DeleteResponse
		wantErr bool
	}{
		{"nil ctx", args{ctx: nil}, credentials.DeleteResponse{}, true},
		{"ok", args{context.Background(), projectID, instanceID, credID}, credentials.DeleteResponse{
			Error:       "string",
			Description: "string",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := dsa.Credentials.Delete(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.credentialID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSACredentialsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("DSACredentialsService.Delete() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
