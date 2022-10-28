package members_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/members"
)

func TestNew(t *testing.T) {
	c, _, teardown, _ := client.MockServer()
	defer teardown()
	type args struct {
		c common.Client
	}
	tests := []struct {
		name string
		args args
		want *members.MembersService
	}{
		{"test directly", args{c}, members.New(c)},
		{"test through client", args{c}, &c.Membership.Members},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := members.New(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembersService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	m := members.New(c)
	type args struct {
		ctx          context.Context
		resourceID   string
		resourceType string
	}

	want := members.ResourceMembers{
		ResourceID:   "test",
		ResourceType: "project",
		Members: []members.Member{
			{
				Subject: "user@domain",
				Role:    "my-role",
			},
		},
	}

	testUUID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	mux.HandleFunc("/membership/v2/project/"+testUUID+"/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	tests := []struct {
		name        string
		args        args
		wantMembers members.ResourceMembers
		wantErr     bool
	}{
		{"resource not found", args{context.Background(), "not-found", "project"}, members.ResourceMembers{}, true},
		{"ctx is nil", args{nil, testUUID, "project"}, members.ResourceMembers{}, true},
		{"invalid resource type", args{context.Background(), testUUID, "projects"}, want, true},
		{"resource found", args{context.Background(), testUUID, "project"}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMembers, err := m.Get(tt.args.ctx, tt.args.resourceID, tt.args.resourceType)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMembers, tt.wantMembers) && !tt.wantErr {
				t.Errorf("MembersService.Get() = %v, want %v", gotMembers, tt.wantMembers)
			}
		})
	}
}

func TestMembersService_Add(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	m := members.New(c)

	testUUID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	wantMembers := []members.Member{
		{
			Subject: "user@domain",
			Role:    "project.admin",
		},
	}

	want := members.ResourceMembers{
		ResourceID:   testUUID,
		ResourceType: "project",
		Members:      wantMembers,
	}

	mux.HandleFunc(fmt.Sprintf("/membership/v2/%s/members", testUUID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx          context.Context
		resourceID   string
		resourceType string
		members      []members.Member
	}

	tests := []struct {
		name    string
		args    args
		wantRes members.ResourceMembers
		wantErr bool
	}{
		{"resource not found", args{context.Background(), "not-found", "project", []members.Member{}}, members.ResourceMembers{}, true},
		{"ctx is nil", args{nil, "resource-id-string", "project", []members.Member{}}, members.ResourceMembers{}, true},
		{"members is nil", args{nil, testUUID, "project", nil}, members.ResourceMembers{}, true},
		{"bad resource type", args{nil, testUUID, "projects", wantMembers}, want, true},
		{"all ok", args{context.Background(), testUUID, "project", wantMembers}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := m.Add(tt.args.ctx, tt.args.resourceID, tt.args.resourceType, tt.args.members)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MembersService.Add() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMembersService_Replace(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	m := members.New(c)

	wantMembers := []members.Member{
		{
			Subject: "user@domain",
			Role:    "project.admin",
		},
	}

	testUUID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	want := members.ResourceMembers{
		ResourceID:   testUUID,
		ResourceType: "project",
		Members:      wantMembers,
	}

	wantNoMembers := members.ResourceMembers{
		ResourceID:   testUUID,
		ResourceType: "project",
		Members:      []members.Member{},
	}

	mux.HandleFunc(fmt.Sprintf("/membership/v2/%s/members", testUUID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var p members.AddMembersRequest
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		b, _ := json.Marshal(members.ResourceMembers{
			ResourceID:   testUUID,
			ResourceType: p.ResourceType,
			Members:      p.Members,
		})
		fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx          context.Context
		resourceID   string
		resourceType string
		members      []members.Member
	}
	tests := []struct {
		name    string
		args    args
		wantRes members.ResourceMembers
		wantErr bool
	}{
		{"resource not found", args{context.Background(), "not-found", "project", []members.Member{}}, members.ResourceMembers{}, true},
		{"ctx is nil", args{nil, "resource-id-string", "project", []members.Member{}}, members.ResourceMembers{}, true},
		{"members is nil", args{nil, testUUID, "project", nil}, members.ResourceMembers{}, true},
		{"bad resource type", args{context.Background(), testUUID, "projects", wantMembers}, want, true},
		{"all ok", args{context.Background(), testUUID, "project", wantMembers}, want, false},
		{"all ok - no members", args{context.Background(), testUUID, "project", []members.Member{}}, wantNoMembers, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := m.Replace(tt.args.ctx, tt.args.resourceID, tt.args.resourceType, tt.args.members)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersService.Replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MembersService.Replace() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMembersService_Remove(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	m := members.New(c)

	testUUID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	initialMembers := []members.Member{
		{
			Subject: "user-1@domain",
			Role:    "project.admin",
		},
		{
			Subject: "user-2@domain",
			Role:    "project.admin",
		},
	}

	toRemove := []members.Member{
		{
			Subject: "user-1@domain",
			Role:    "project.admin",
		},
	}

	want := members.ResourceMembers{
		ResourceID:   testUUID,
		ResourceType: "project",
		Members: []members.Member{
			{
				Subject: "user-2@domain",
				Role:    "project.admin",
			},
		},
	}

	mux.HandleFunc(fmt.Sprintf("/membership/v2/%s/members/remove", testUUID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var p members.AddMembersRequest
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		processed := []members.Member{}
		for _, k := range initialMembers {
			for _, j := range p.Members {
				if k.Subject == j.Subject && k.Role == j.Role {
					continue
				}
				processed = append(processed, k)
			}
		}

		b, _ := json.Marshal(members.ResourceMembers{
			ResourceID:   testUUID,
			ResourceType: p.ResourceType,
			Members:      processed,
		})
		fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx          context.Context
		resourceID   string
		resourceType string
		members      []members.Member
	}
	tests := []struct {
		name    string
		args    args
		wantRes members.ResourceMembers
		wantErr bool
	}{
		{"resource not found", args{context.Background(), "not-found", "project", []members.Member{}}, members.ResourceMembers{}, true},
		{"ctx is nil", args{nil, "resource-id-string", "project", []members.Member{}}, members.ResourceMembers{}, true},
		{"members is nil", args{nil, testUUID, "project", nil}, members.ResourceMembers{}, true},
		{"bad resource type", args{context.Background(), testUUID, "projects", toRemove}, want, true},
		{"all ok", args{context.Background(), testUUID, "project", toRemove}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := m.Remove(tt.args.ctx, tt.args.resourceID, tt.args.resourceType, tt.args.members)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersService.Replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MembersService.Replace() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMembersService_Validate(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	m := members.New(c)

	resourceID := "my-project-aUze1quX"
	resourceType := "project"
	wantJson := []byte(`{
		"resourceType": "project",
		"members": [
		  {
			"subject": "user@mail.schwarz",
			"role": "project.owner",
			"condition": {
			  "expiresAt": "2019-08-24T14:15:22Z"
			}
		  }
		]
	  }`)

	var want members.ResourceMembers
	if err := json.Unmarshal(wantJson, &want); err != nil {
		t.Error(err)
	}
	mux.HandleFunc(fmt.Sprintf("/membership/v2/%s/members/validate", resourceID), func(w http.ResponseWriter, r *http.Request) {
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
		members      []members.Member
	}
	tests := []struct {
		name    string
		args    args
		wantRes members.ResourceMembers
		wantErr bool
	}{
		{"all ok", args{context.Background(), resourceID, resourceType, want.Members}, want, false},
		{"bad ctx", args{nil, resourceID, resourceType, want.Members}, want, true},
		{"bad type", args{nil, resourceID, "abc", want.Members}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := m.Validate(tt.args.ctx, tt.args.resourceID, tt.args.resourceType, tt.args.members)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersService.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MembersService.Validate() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
