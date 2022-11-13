// package users is used to manange Postgres Flex users

package users_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/postgres-flex"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/postgres-flex/users"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPathList   = consts.API_PATH_POSTGRES_FLEX_USERS
	apiPathCreate = consts.API_PATH_POSTGRES_FLEX_USER
	apiPathGet    = consts.API_PATH_POSTGRES_FLEX_USER
)

func TestPostgresUsersService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	postgres := postgres.New(c)

	projectID := "abc"
	instanceID := "efg"

	mux.HandleFunc(fmt.Sprintf(apiPathList, projectID, instanceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"count": 0,
			"items": [
			  {
				"id": "string",
				"username": "string"
			  }
			]
		  }`)
	})

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes users.ListResponse
		wantErr bool
	}{
		{"ok", args{context.Background(), projectID, instanceID}, users.ListResponse{Items: []users.UserListItem{{"string", "string"}}}, false},
		{"bad", args{nil, projectID, instanceID}, users.ListResponse{Items: []users.UserListItem{{"string", "string"}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := postgres.Users.List(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresUsersService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("PostgresUsersService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestPostgresUsersService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	postgres := postgres.New(c)

	projectID := "abc"
	instanceID := "efg"
	userID := "123"

	mux.HandleFunc(fmt.Sprintf(apiPathGet, projectID, instanceID, userID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"item": {
			  "database": "string",
			  "hostname": "string",
			  "id": "123",
			  "port": 0,
			  "roles": [
				"string"
			  ],
			  "username": "string"
			}
		  }`)
	})

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		userID     string
	}
	tests := []struct {
		name    string
		args    args
		wantRes users.GetResponse
		wantErr bool
	}{
		{"ok", args{context.Background(), projectID, instanceID, userID}, users.GetResponse{Item: users.UserGetItem{
			Database: "string",
			Hostname: "string",
			Username: "string",
			ID:       userID,
			Roles:    []string{"string"},
		}}, false},
		{"nil ctx", args{nil, projectID, instanceID, userID}, users.GetResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := postgres.Users.Get(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresUsersService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("PostgresUsersService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestPostgresUsersService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	postgres := postgres.New(c)

	projectID := "abc"
	instanceID := "efg"
	userID := "123"

	mux.HandleFunc(fmt.Sprintf(apiPathGet, projectID, instanceID, userID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"item": {
			  "database": "string",
			  "hostname": "string",
			  "id": "123",
			  "password": "string",
			  "port": 0,
			  "roles": [
				"string"
			  ],
			  "uri": "string",
			  "username": "string"
			}
		  }`)
	})

	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		userID     string
		username   string
		database   string
		roles      []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes users.CreateResponse
		wantErr bool
	}{
		{"ok", args{
			context.Background(),
			projectID,
			instanceID,
			userID,
			"string",
			"string",
			[]string{"string"},
		}, users.CreateResponse{
			Item: users.User{
				Database: "string",
				Hostname: "string",
				ID:       userID,
				Password: "string",
				Port:     0,
				Roles:    []string{"string"},
				URI:      "string",
				Username: "string",
			},
		}, false},

		{"nil ctx", args{ctx: nil}, users.CreateResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := postgres.Users.Create(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.userID, tt.args.username, tt.args.database, tt.args.roles)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresUsersService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("PostgresUsersService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
