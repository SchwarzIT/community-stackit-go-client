package projects_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/projects"
)

const (
	apiPath      = consts.API_PATH_OBJECT_STORAGE_PROJECT
	apiPathForce = consts.API_PATH_OBJECT_STORAGE_PROJECT_FORCE_DELETE
)

func TestStorageProjectService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	want := projects.ObjectStorageProjectResponse{
		ProjectID: projectID,
		Scope:     "PUBLIC",
	}

	mux.HandleFunc(fmt.Sprintf(apiPath, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes projects.ObjectStorageProjectResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID}, want, false},
		{"ctx is canceled", args{ctx, projectID}, want, true},
		{"project not found", args{context.Background(), "my-project"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.Projects.Get(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageProjectService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("StorageProjectService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestStorageProjectService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	want := projects.ObjectStorageProjectResponse{
		ProjectID: projectID,
		Scope:     "PUBLIC",
	}

	mux.HandleFunc(fmt.Sprintf(apiPath, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes projects.ObjectStorageProjectResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID}, want, false},
		{"ctx is canceled", args{ctx, projectID}, want, true},
		{"project not found", args{context.Background(), "my-project"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.Projects.Create(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageProjectService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("StorageProjectService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestStorageProjectService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	mux.HandleFunc(fmt.Sprintf(apiPath, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

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
		{"ctx is canceled", args{ctx, projectID}, true},
		{"project not found", args{context.Background(), "my-project"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Projects.Delete(tt.args.ctx, tt.args.projectID); (err != nil) != tt.wantErr {
				t.Errorf("StorageProjectService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageProjectService_ForceDelete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	mux.HandleFunc(fmt.Sprintf(apiPathForce, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

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
		{"ctx is canceled", args{ctx, projectID}, true},
		{"project not found", args{context.Background(), "my-project"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Projects.ForceDelete(tt.args.ctx, tt.args.projectID); (err != nil) != tt.wantErr {
				t.Errorf("StorageProjectService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
