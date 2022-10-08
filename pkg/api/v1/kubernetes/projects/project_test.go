package projects_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes/projects"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

func TestKubernetesProjectsService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	want := projects.KubernetesProjectsResponse{
		ProjectID: projectID,
		State:     consts.SKE_PROJECT_STATUS_UNSPECIFIED,
	}

	mux.HandleFunc("/ske/v1/projects/"+projectID, func(w http.ResponseWriter, r *http.Request) {
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
		wantRes projects.KubernetesProjectsResponse
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
				t.Errorf("KubernetesProjectsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("KubernetesProjectsService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestKubernetesProjectsService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	want := projects.KubernetesProjectsResponse{
		ProjectID: projectID,
		State:     consts.SKE_PROJECT_STATUS_UNSPECIFIED,
	}

	mux.HandleFunc("/ske/v1/projects/"+projectID, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
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
		wantRes projects.KubernetesProjectsResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID}, want, false},
		{"ctx is canceled", args{ctx, projectID}, want, true},
		{"project not found", args{context.Background(), "my-project"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, _, err := s.Projects.Create(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesProjectsService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("KubernetesProjectsService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestKubernetesProjectsService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	mux.HandleFunc("/ske/v1/projects/"+projectID, func(w http.ResponseWriter, r *http.Request) {
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
			if _, err := s.Projects.Delete(tt.args.ctx, tt.args.projectID); (err != nil) != tt.wantErr {
				t.Errorf("KubernetesProjectsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
