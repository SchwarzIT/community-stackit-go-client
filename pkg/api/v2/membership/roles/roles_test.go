package roles_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/roles"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath         = consts.API_PATH_MEMBERSHIP_V2_ROLES
	apiPathWithType = consts.API_PATH_MEMBERSHIP_V2_ROLES_WITH_RESOURCE_TYPE
)

func TestRolesService_AddCustom(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	r := membership.New(c).Roles

	resourceID := "some-resource-id-123"
	wantJson := []byte(`{
		"resourceType": "project",
		"roles": [
		  {
			"name": "project.special-owner",
			"description": "An owner of the project",
			"permissions": [
			  {
				"name": "organization.projects.read"
			  }
			]
		  }
		]
	  }`)

	var want roles.Roles
	if err := json.Unmarshal(wantJson, &want); err != nil {
		t.Error(err)
	}

	mux.HandleFunc(fmt.Sprintf(apiPath, resourceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(wantJson))
	})

	type args struct {
		ctx          context.Context
		resourceID   string
		resourceType string
		roles        []roles.Role
	}
	tests := []struct {
		name     string
		args     args
		wantRes  roles.Roles
		wantErr  bool
		runEqual bool
	}{
		{"all ok", args{context.Background(), resourceID, "project", want.Roles}, want, false, true},
		{"wrong type", args{context.Background(), resourceID, "something", want.Roles}, want, true, true},
		{"wrong id", args{context.Background(), "abc", "project", want.Roles}, want, true, false},
		{"bad ctx", args{nil, resourceID, "project", want.Roles}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := r.AddCustom(tt.args.ctx, tt.args.resourceID, tt.args.resourceType, tt.args.roles...)
			if (err != nil) != tt.wantErr && !tt.runEqual {
				t.Errorf("RolesService.AddCustom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && tt.runEqual {
				t.Errorf("RolesService.AddCustom() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestRolesService_RemoveCustom(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	r := membership.New(c).Roles

	resourceID := "some-resource-id-123"
	wantJson := []byte(`{
		"resourceType": "project",
		"roles": [
		  {
			"name": "project.specialowner"
		  }
		]
	  }`)

	var want roles.Roles
	if err := json.Unmarshal(wantJson, &want); err != nil {
		t.Error(err)
	}
	mux.HandleFunc(fmt.Sprintf(apiPath+"/remove", resourceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(wantJson))
	})

	type args struct {
		ctx          context.Context
		resourceID   string
		resourceType string
		roles        []roles.Role
	}
	tests := []struct {
		name     string
		args     args
		wantR    roles.Roles
		wantErr  bool
		runEqual bool
	}{
		{"all ok", args{context.Background(), resourceID, "project", []roles.Role{}}, want, false, true},
		{"wrong type", args{context.Background(), resourceID, "something", []roles.Role{}}, want, true, true},
		{"wrong id", args{context.Background(), "abc", "project", []roles.Role{}}, want, true, false},
		{"bad ctx", args{nil, resourceID, "project", []roles.Role{}}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := r.RemoveCustom(tt.args.ctx, tt.args.resourceID, tt.args.resourceType, tt.args.roles...)
			if (err != nil) != tt.wantErr && !tt.runEqual {
				t.Errorf("RolesService.RemoveCustom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) && tt.runEqual {
				t.Errorf("RolesService.RemoveCustom() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestRolesService_GetByResource(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	r := membership.New(c).Roles

	resourceID := "my-project-aUze1quX"
	resourceType := "project"
	wantJson := []byte(`{
		"resourceId": "my-project-aUze1quX",
		"resourceType": "project",
		"roles": [
		  {
			"name": "project.owner",
			"description": "An owner of the project",
			"permissions": [
			  {
				"name": "organization.projects.read",
				"description": "Can read projects of an organization"
			  }
			]
		  }
		]
	  }`)

	var want roles.Roles
	if err := json.Unmarshal(wantJson, &want); err != nil {
		t.Error(err)
	}
	mux.HandleFunc(fmt.Sprintf(apiPathWithType, resourceType, resourceID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(wantJson))
	})

	type args struct {
		ctx          context.Context
		resourceID   string
		resourceType string
	}
	tests := []struct {
		name     string
		args     args
		wantR    roles.Roles
		wantErr  bool
		runEqual bool
	}{
		{"all ok", args{context.Background(), resourceID, resourceType}, want, false, true},
		{"nil ctx", args{nil, resourceID, resourceType}, want, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := r.GetByResource(tt.args.ctx, tt.args.resourceID, tt.args.resourceType)
			if (err != nil) != tt.wantErr {
				t.Errorf("RolesService.GetByResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) && tt.runEqual {
				t.Errorf("RolesService.GetByResource() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
