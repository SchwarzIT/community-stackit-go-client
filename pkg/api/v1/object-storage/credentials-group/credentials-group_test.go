package credentialsgroup_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/internal/clients"
	credentialsgroup "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage/credentials-group"
	"golang.org/x/exp/slices"

	"github.com/SchwarzIT/community-stackit-go-client"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	listPath                 = consts.API_PATH_OBJECT_STORAGE_CREDENTIALS_LIST
	createPath               = consts.API_PATH_OBJECT_STORAGE_CREDENTIALS_CREATE
	deletePath               = consts.API_PATH_OBJECT_STORAGE_CREDENTIALS_DELETE
	skipAcceptanceTestList   = true
	skipAcceptanceTestCreate = true
	skipAcceptanceTestDelete = true
)

func TestAccProjectsService_Create(t *testing.T) {
	if skipAcceptanceTestCreate {
		t.Skip()
	}

	c, err := clients.LocalClient()
	if err != nil {
		t.Error(err)
	}
	s := objectstorage.New(c).CredentialsGroup
	projectID := "74c00201-a5fd-4e0d-afd6-f895bef7a4da"
	credentialsGroupName := "my-credentialsGroup"

	type args struct {
		ctx         context.Context
		projectID   string
		displayName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, credentialsGroupName}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Create(tt.args.ctx, tt.args.projectID, tt.args.displayName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageCredentialsGroupService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}

func TestAccProjectsService_List(t *testing.T) {
	if skipAcceptanceTestList {
		t.Skip()
	}

	c, err := clients.LocalClient()
	if err != nil {
		t.Error(err)
	}
	s := objectstorage.New(c).CredentialsGroup
	projectID := "74c00201-a5fd-4e0d-afd6-f895bef7a4da"
	credentialsGroupName := "my-credentialsGroup"
	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageCredentialsGroupService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if slices.IndexFunc(gotRes.CredentialsGroups, func(c credentialsgroup.CredentialsGroup) bool { return c.DisplayName == credentialsGroupName }) == -1 {
				t.Errorf("ObjectStorageCredentialsGroupService.List() = %v, want to contain %v", gotRes, credentialsGroupName)
			}
		})
	}

}

func TestAccProjectsService_Delete(t *testing.T) {
	if skipAcceptanceTestDelete {
		t.Skip()
	}

	c, err := clients.LocalClient()
	if err != nil {
		t.Error(err)
	}
	s := objectstorage.New(c).CredentialsGroup
	projectID := "74c00201-a5fd-4e0d-afd6-f895bef7a4da"
	credentialsGroupName := "my-credentialsGroup"

	type args struct {
		ctx         context.Context
		projectID   string
		displayName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, credentialsGroupName}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.List(tt.args.ctx, tt.args.projectID)
			if err != nil {
				t.Error(err)
			}
			i := slices.IndexFunc(
				resp.CredentialsGroups,
				func(c credentialsgroup.CredentialsGroup) bool { return c.DisplayName == tt.args.displayName },
			)
			if i == -1 {
				t.Fatalf("credentials group %s not found", tt.args.displayName)
			}
			credentialsGroupId := resp.CredentialsGroups[i].CredentialsGroupId
			err = s.Delete(tt.args.ctx, tt.args.projectID, credentialsGroupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageCredentialsGroupService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestObjectStorageCredentialsGroupsService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).CredentialsGroup

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	credentialsGroupName := "my-credentialsGroup"
	want := credentialsgroup.CredentialsGroupResponse{
		Project: projectID,
		CredentialsGroups: []credentialsgroup.CredentialsGroup{
			{CredentialsGroupId: credentialsGroupName},
		},
	}
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	mux.HandleFunc(fmt.Sprintf(listPath, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes credentialsgroup.CredentialsGroupResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID}, want, false},
		{"ctx is canceled", args{ctx, projectID}, want, true},
		{"project not found", args{context.Background(), "some-id"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageCredentialsGroupService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("ObjectStorageCredentialsGroupService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestObjectStorageCredentialsGroupsService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).CredentialsGroup

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	credentialsGroupName := "my-credentialsGroup"
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	mux.HandleFunc(fmt.Sprintf(createPath, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	type args struct {
		ctx         context.Context
		projectID   string
		displayName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, credentialsGroupName}, false},
		{"ctx is canceled", args{ctx, projectID, credentialsGroupName}, true},
		{"project not found", args{context.Background(), "my-project", credentialsGroupName}, true},
		{"display name invalid", args{context.Background(), projectID, ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Create(tt.args.ctx, tt.args.projectID, tt.args.displayName); (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageCredentialsGroupService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestObjectStorageCredentialsGroupsService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).CredentialsGroup

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	credentialsGroupName := "my-credentialsGroup"
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	mux.HandleFunc(fmt.Sprintf(deletePath, projectID, credentialsGroupName), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	type args struct {
		ctx                  context.Context
		projectID            string
		credentialsGroupName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, credentialsGroupName}, false},
		{"ctx is canceled", args{ctx, projectID, credentialsGroupName}, true},
		{"project not found", args{context.Background(), "my-project", credentialsGroupName}, true},
		{"credentialsGroup not found", args{context.Background(), projectID, "some-credentialsGroup"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Delete(tt.args.ctx, tt.args.projectID, tt.args.credentialsGroupName); (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageCredentialsGroupService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
