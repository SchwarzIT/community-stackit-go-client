package projects_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes/projects"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
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

	ctx2, cancel2 := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel2()

	ctx3, cancel3 := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel3()

	state1 := projects.KubernetesProjectsResponse{
		ProjectID: projectID,
		State:     consts.SKE_PROJECT_STATUS_UNSPECIFIED,
	}

	state2 := projects.KubernetesProjectsResponse{
		ProjectID: projectID,
		State:     consts.SKE_PROJECT_STATUS_CREATED,
	}

	state3 := projects.KubernetesProjectsResponse{
		ProjectID: projectID,
		State:     consts.SKE_PROJECT_STATUS_FAILED,
	}

	mux.HandleFunc("/ske/v1/projects/"+projectID, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
			return
		}
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			if ctx2.Err() == nil {
				b, _ := json.Marshal(state1)
				fmt.Fprint(w, string(b))
				return
			}
			if ctx3.Err() == nil {
				b, _ := json.Marshal(state2)
				fmt.Fprint(w, string(b))
				return
			}

			b, _ := json.Marshal(state3)
			fmt.Fprint(w, string(b))
			return
		}
		t.Error("wrong method")
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name     string
		args     args
		wantRes  projects.KubernetesProjectsResponse
		wantErr  bool
		goodWait bool
	}{
		{"all ok", args{context.Background(), projectID}, want, false, true},
		{"ctx is canceled", args{ctx, projectID}, want, true, false},
		{"project not found", args{context.Background(), "my-project"}, want, true, false},
	}

	var goodWait, badWait *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, w, err := s.Projects.Create(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesProjectsService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("KubernetesProjectsService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.goodWait {
				goodWait = w
			} else {
				badWait = w
			}
		})
	}

	c.SetRetry(nil)
	goodWait.SetThrottle(20 * time.Millisecond)
	badWait.SetThrottle(20 * time.Millisecond)

	if _, err := goodWait.Wait(); err != nil {
		t.Errorf("returned: %v", err)
	}
	if _, err := badWait.Wait(); err == nil {
		t.Error("expected error but got nil")
	}

	time.Sleep(2 * time.Second)
	if _, err := goodWait.Wait(); err == nil {
		t.Error("expected error but got nil [2]")
	}

}

func TestKubernetesProjectsService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	ctx2, cancel2 := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel2()

	ctx3, cancel3 := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel3()

	state1 := projects.KubernetesProjectsResponse{
		ProjectID: projectID,
		State:     consts.SKE_PROJECT_STATUS_DELETING,
	}

	mux.HandleFunc("/ske/v1/projects/"+projectID, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")

			if ctx2.Err() == nil {
				w.WriteHeader(http.StatusOK)
				b, _ := json.Marshal(state1)
				fmt.Fprint(w, string(b))
				return
			}

			if ctx3.Err() == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		t.Error("wrong method")
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		goodWait bool
	}{
		{"all ok", args{context.Background(), projectID}, false, true},
		{"ctx is canceled", args{ctx, projectID}, true, false},
		{"project not found", args{context.Background(), "my-project"}, true, false},
	}

	var goodWait, badWait *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := s.Projects.Delete(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesProjectsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.goodWait {
				goodWait = w
			} else {
				badWait = w
			}
		})
	}

	c.SetRetry(nil)
	goodWait.SetThrottle(20 * time.Millisecond)
	badWait.SetThrottle(20 * time.Millisecond)

	if _, err := goodWait.Wait(); err != nil {
		t.Errorf("returned: %v", err)
	}
	time.Sleep(2 * time.Second)
	if _, err := goodWait.Wait(); err == nil {
		t.Error("expected error but got nil")
	}

}
