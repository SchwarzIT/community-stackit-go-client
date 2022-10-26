package projects_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	client "github.com/SchwarzIT/community-stackit-go-client"
	resourcemanager "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-manager"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-manager/projects"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

func TestProjectsService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := resourcemanager.New(c).Projects

	containerID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	want := projects.ProjectResponse{
		Name:        "my-project",
		ContainerID: containerID,
	}

	mux.HandleFunc(fmt.Sprintf("/resource-manager/v2/projects/%s", containerID), func(w http.ResponseWriter, r *http.Request) {
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
		wantRes projects.ProjectResponse
		wantErr bool
	}{
		{"resource not found", args{context.Background(), "not-found"}, projects.ProjectResponse{}, true},
		{"ctx is nil", args{nil, containerID}, projects.ProjectResponse{}, true},
		{"all ok", args{context.Background(), containerID}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := p.Get(tt.args.ctx, tt.args.containerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ProjectsService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestProjectsService_GetLifecycleState(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := resourcemanager.New(c).Projects

	containerID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	want := projects.ProjectResponse{
		Name:           "my-project",
		ContainerID:    containerID,
		LifecycleState: "CREATED",
	}

	mux.HandleFunc(fmt.Sprintf("/resource-manager/v2/projects/%s", containerID), func(w http.ResponseWriter, r *http.Request) {
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
		wantRes projects.ProjectResponse
		wantErr bool
	}{
		{"resource not found", args{context.Background(), "not-found"}, projects.ProjectResponse{}, true},
		{"ctx is nil", args{nil, containerID}, projects.ProjectResponse{}, true},
		{"all ok", args{context.Background(), containerID}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := p.GetLifecycleState(tt.args.ctx, tt.args.containerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectsService.GetLifecycleState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes.LifecycleState {
				t.Errorf("ProjectsService.GetLifecycleState() = %v, want %v", gotRes, tt.wantRes.LifecycleState)
			}
		})
	}
}

func TestProjectsService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := resourcemanager.New(c).Projects

	containerID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	member := []projects.ProjectMember{
		{Subject: "user@domain", Role: "project.owner"},
	}
	memberNotFound := []projects.ProjectMember{
		{Subject: "notfound@domain", Role: "project.owner"},
	}

	res := fmt.Sprintf(`{
		"name": "my-project",
		"parent": {
		  "id": "54066bf4-1aff-4f7b-9f83-fb23c348fee3",
		  "containerId": "container-name-Rq86WY1",
		  "type": "ORGANIZATION"
		},
		"containerId": "%s",
		"lifecycleState": "CREATING",
		"labels": {
		  "billingReference": "T-0123456E",
		  "scope": "PUBLIC"
		},
		"updateTime": "2021-08-24T14:15:22Z",
		"creationTime": "2021-08-24T14:15:22Z"
	  }`, containerID)

	want := projects.ProjectResponse{
		Name: "my-project",
		Parent: projects.Parent{
			ID:          "54066bf4-1aff-4f7b-9f83-fb23c348fee3",
			ContainerID: "container-name-Rq86WY1",
			Type:        "ORGANIZATION",
		},
		ContainerID: containerID,
		Labels: map[string]string{
			"billingReference": "T-0123456E",
			"scope":            "PUBLIC",
		},
		LifecycleState: "CREATING",
		UpdateTime:     "2021-08-24T14:15:22Z",
		CreationTime:   "2021-08-24T14:15:22Z",
	}

	mux.HandleFunc("/resource-manager/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")

		var pr projects.CreateProjectRequest
		err := json.NewDecoder(r.Body).Decode(&pr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, v := range pr.Members {
			if v.Subject == "notfound@domain" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, `{}`)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, res)
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	longLabels := make(map[string]string, 101)
	for x := 0; x < 101; x++ {
		longLabels[strconv.Itoa(x)] = "x"
	}
	type args struct {
		ctx     context.Context
		name    string
		labels  map[string]string
		members []projects.ProjectMember
	}
	tests := []struct {
		name    string
		args    args
		wantRes projects.ProjectResponse
		wantErr bool
		useWait bool
	}{
		{"no owner", args{context.Background(), "my-project", map[string]string{}, []projects.ProjectMember{}}, want, true, false},
		{"no billing", args{context.Background(), "my-project", map[string]string{}, member}, want, true, false},
		{"bad project name", args{context.Background(), "my project!", map[string]string{}, member}, want, true, false},
		{"all ok", args{context.Background(), "my-project", map[string]string{"billingReference": "T-0123456B", "scope": "PUBLIC"}, member}, want, false, true},
		{"too many labels", args{context.Background(), "my-project", longLabels, member}, want, true, false},
		{"no scope label", args{context.Background(), "my-project", map[string]string{"billingReference": "T-0123456B"}, member}, want, true, false},
		{"ctx is canceled", args{ctx, "my-project", map[string]string{"billingReference": "T-0123456B", "scope": "PUBLIC"}, member}, want, true, false},
		{"user not found", args{context.Background(), "my-project", map[string]string{"billingReference": "T-0123456B", "scope": "PUBLIC"}, memberNotFound}, want, true, false},
	}

	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, w, err := p.Create(tt.args.ctx, tt.args.name, tt.args.labels, tt.args.members...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectsService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("ProjectsService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.useWait {
				process = w
			}
		})
	}

	baseDuration := 200 * time.Millisecond
	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	ctx2, td2 := context.WithTimeout(context.Background(), 2*baseDuration)
	defer td2()

	ctx3, td3 := context.WithTimeout(context.Background(), 3*baseDuration)
	defer td3()

	ctx4, td4 := context.WithTimeout(context.Background(), 4*baseDuration)
	defer td4()

	mux.HandleFunc(fmt.Sprintf("/resource-manager/v2/projects/%s", want.ContainerID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("wrong method %s", r.Method)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if ctx1.Err() == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if ctx2.Err() == nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if ctx3.Err() == nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"lifecycleState": "CREATING"}`)
			return
		}

		if ctx4.Err() == nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"lifecycleState": "ACTIVE"}`)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"lifecycleState": "DELETING"}`)
	})

	process.SetThrottle(baseDuration)
	// first test run: expect Wait to exit with an error
	if _, err := process.Wait(); err == nil {
		t.Error("expected error but got nil")
	}

	time.Sleep(baseDuration)
	// 2nd test: expect Forbidden, creating, and exit with success after retry
	if _, err := process.Wait(); err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	time.Sleep(baseDuration)
	// last test run: expect error because of DELETING status
	if _, err := process.Wait(); err == nil {
		t.Error("expected error but got nil")
	}
}

func TestProjectsService_Update(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := resourcemanager.New(c).Projects

	containerID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	want := projects.ProjectResponse{
		Name:        "my-project",
		ContainerID: containerID,
	}

	mux.HandleFunc("/resource-manager/v2/projects/"+containerID, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")

		var pr projects.CreateProjectRequest
		err := json.NewDecoder(r.Body).Decode(&pr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx         context.Context
		containerID string
		name        string
		parentID    string
		labels      map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantRes projects.ProjectResponse
		wantErr bool
	}{
		{"no billing", args{context.Background(), containerID, "my-project", consts.SCHWARZ_ORGANIZATION_ID, map[string]string{}}, want, true},
		{"bad project name", args{context.Background(), containerID, "my project!", consts.SCHWARZ_ORGANIZATION_ID, map[string]string{}}, want, true},
		{"all ok", args{context.Background(), containerID, "my-project", consts.SCHWARZ_ORGANIZATION_ID, map[string]string{"billingReference": "T-0123456B", "scope": "PUBLIC"}}, want, false},
		{"no scope", args{context.Background(), containerID, "my-project", consts.SCHWARZ_ORGANIZATION_ID, map[string]string{"billingReference": "T-0123456B"}}, want, true},
		{"bad billing", args{context.Background(), containerID, "my-project", consts.SCHWARZ_ORGANIZATION_ID, map[string]string{"billingReference": "T#$%0123456B", "scope": "PUBLIC"}}, want, true},
		{"ctx is canceled", args{ctx, containerID, "my-project", consts.SCHWARZ_ORGANIZATION_ID, map[string]string{"billingReference": "T-0123456B", "scope": "PUBLIC"}}, want, true},
		{"project not found", args{context.Background(), "something", "my-project", consts.SCHWARZ_ORGANIZATION_ID, map[string]string{"billingReference": "T-0123456B", "scope": "PUBLIC"}}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := p.Update(tt.args.ctx, tt.args.containerID, tt.args.name, tt.args.parentID, tt.args.labels)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectsService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("ProjectsService.Update() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestProjectsService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := resourcemanager.New(c).Projects

	containerID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	baseDuration := 200 * time.Millisecond
	ctx1, td1 := context.WithTimeout(context.Background(), 1*baseDuration)
	defer td1()

	mux.HandleFunc("/resource-manager/v2/projects/"+containerID, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			return
		}

		if r.Method == http.MethodGet {

			w.Header().Set("Content-Type", "application/json")

			if ctx1.Err() == nil {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, `{"lifecycleState": "DELETING"}`)
				return
			}

			w.WriteHeader(http.StatusNotFound)
			return

		}
		t.Error("wrong method")
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx         context.Context
		containerID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		useWait bool
	}{
		{"all ok", args{context.Background(), containerID}, false, true},
		{"ctx is canceled", args{ctx, containerID}, true, false},
	}

	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := p.Delete(tt.args.ctx, tt.args.containerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.useWait {
				process = w
			}
		})
	}

	process.SetThrottle(baseDuration)
	// 1nd test: expect success after retry
	if _, err := process.Wait(); err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}

func TestProjectsService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	p := resourcemanager.New(c).Projects

	want := projects.ProjectsResponse{
		Iteams: []projects.ProjectResponse{},
		Offset: 0,
		Limit:  50,
	}

	want2 := projects.ProjectsResponse{
		Iteams: []projects.ProjectResponse{},
		Offset: 2,
		Limit:  50,
	}

	mux.HandleFunc("/resource-manager/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	mux.HandleFunc("/resource-manager/v2/projects?offset=2&limit=50", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want2)
		fmt.Fprint(w, string(b))
	})
	mux.HandleFunc("/resource-manager/v2/projects?offset=0&limit=50&containerIds=my-container-123", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx               context.Context
		containerParentID string
		filters           map[string]string
		containerIDs      []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes projects.ProjectsResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), "container-parent-1234", map[string]string{}, nil}, want, false},
		{"all ok 2", args{context.Background(), "container-parent-1234", map[string]string{"offset": "2"}, nil}, want, false},
		{"all ok 3", args{context.Background(), "", map[string]string{"offset": "2"}, []string{"my-container-123"}}, want, false},
		{"no parent, no project ids", args{context.Background(), "", map[string]string{}, nil}, want, true},
		{"bad offset", args{context.Background(), "container-parent-1234", map[string]string{"offset": "abc"}, nil}, want, true},
		{"bad offset 2", args{context.Background(), "container-parent-1234", map[string]string{"offset": "-1"}, nil}, want, true},
		{"bad limit", args{context.Background(), "container-parent-1234", map[string]string{"limit": "abc"}, nil}, want, true},
		{"bad limit 2", args{context.Background(), "container-parent-1234", map[string]string{"limit": "101"}, nil}, want, true},
		{"bad creation time start", args{context.Background(), "container-parent-1234", map[string]string{"creation-time-start": "abc"}, nil}, want, true},
		{"ctx canceled", args{ctx, "container-parent-1234", map[string]string{}, nil}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := p.List(tt.args.ctx, tt.args.containerParentID, tt.args.filters, tt.args.containerIDs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("ProjectsService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
