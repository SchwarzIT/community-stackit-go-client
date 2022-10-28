package organizations_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	resourceManagement "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-management"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-management/organizations"
)

const (
	apiPath = "/resource-management/v2/organizations/%s"
)

func TestOrganizationsService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := resourceManagement.New(c).Organizations

	containerID := "my-container-id-b18796aa7e78"
	want := organizations.OrganizationResponse{
		ContainerID: containerID,
	}

	mux.HandleFunc(fmt.Sprintf(apiPath, containerID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx         context.Context
		containerID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes organizations.OrganizationResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), containerID}, want, false},
		{"nil ctx", args{nil, containerID}, want, true},
		{"not found", args{context.Background(), "abc"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := p.Get(tt.args.ctx, tt.args.containerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("OrganizationsService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
