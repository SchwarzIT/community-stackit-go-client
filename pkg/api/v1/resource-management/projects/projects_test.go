package projects_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/resource-management/projects"
)

func TestProjectService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := projects.New(c)
	want := projects.ProjectGetResponse{
		ProjectID: "abc",
		Name:      "I-am_the-law",
		Labels: projects.ProjectsLabels{
			"T-1234567B",
		},
		Parent: projects.ProjectsParent{
			ID: "987",
		},
	}
	mux.HandleFunc("/resource-management/v1/projects/abc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, err := json.Marshal(want)
		if err != nil {
			log.Fatalf("json response marshal: %v", err)
		}
		fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes projects.ProjectGetResponse
		wantErr bool
	}{
		{"ok", args{context.Background(), "abc"}, want, false},
		{"nil ctx", args{nil, "abc"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := p.Get(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("ProjectService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
