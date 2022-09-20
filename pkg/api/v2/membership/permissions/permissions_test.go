package permissions_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/permissions"
)

const (
	apiPath       = "/membership/v2/permissions"
	apiPathMember = "/membership/v2/users/%s/permissions"
)

func TestPermissionsService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := membership.New(c).Permissions

	wantJson := []byte(`{
		"permissions": [
		  {
			"name": "organization.projects.read",
			"description": "Can read projects of an organization"
		  }
		]
	  }`)

	var want permissions.PermissionList
	if err := json.Unmarshal(wantJson, &want); err != nil {
		t.Error(err)
	}

	mux.HandleFunc(apiPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(wantJson))
	})

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantP   permissions.PermissionList
		wantErr bool
	}{
		{"all ok", args{context.Background()}, want, false},
		{"no ctx", args{nil}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := p.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermissionsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) && !tt.wantErr {
				t.Errorf("PermissionsService.List() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestPermissionsService_GetByEmail(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := membership.New(c).Permissions
	email := "user@mail.schwarz"
	wantJson := []byte(`{
		"items": [
		  {
			"resourceId": "my-project-aUze1quX",
			"resourceType": "project",
			"subject": "user@mail.schwarz",
			"role": "project.owner",
			"condition": {
			  "expiresAt": "2019-08-24T14:15:22Z"
			}
		  }
		],
		"limit": 50,
		"offset": 0
	  }`)

	var want permissions.UserPermissions
	if err := json.Unmarshal(wantJson, &want); err != nil {
		t.Error(err)
	}

	mux.HandleFunc(fmt.Sprintf(apiPathMember, email), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(wantJson))
	})

	type args struct {
		ctx    context.Context
		email  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		wantUp  permissions.UserPermissions
		wantErr bool
	}{
		{"all ok", args{context.Background(), email, 50, 0}, want, false},
		{"bad pagination", args{context.Background(), email, 0, 0}, want, true},
		{"bad pagination 2", args{context.Background(), email, 1, -1}, want, true},
		{"nil ctx", args{nil, email, 50, 0}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUp, err := p.GetByEmail(tt.args.ctx, tt.args.email, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermissionsService.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUp, tt.wantUp) && !tt.wantErr {
				t.Errorf("PermissionsService.GetByEmail() = %v, want %v", gotUp, tt.wantUp)
			}
		})
	}
}
