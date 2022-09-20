package projects_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/internal/clients"
	p "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/projects"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

var (
	skipAcceptanceTestGet    = true
	skipAcceptanceTestCreate = true
	skipAcceptanceTestDelete = true
	skipAcceptanceTestUpdate = true
)

// Requests

func TestProjectsService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	projects := p.New(c)

	mux.HandleFunc("/resource-management/v1/projects/abc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, err := json.Marshal(p.ProjectsResBody{
			ProjectID: "123",
			Name:      "I-am_the-law",
			Labels: p.ProjectsLabelsResBody{
				"T-1234567B",
			},
			Parent: p.ProjectsParentResBody{
				ID: "987",
			},
		})
		if err != nil {
			log.Fatalf("json response marshal: %v", err)
		}
		fmt.Fprint(w, string(b))
	})

	got, err := projects.Get(context.Background(), "abc")
	if err != nil {
		t.Fatal(err)
	}

	want := p.Project{
		ID:               "123",
		Name:             "I-am_the-law",
		BillingReference: "T-1234567B",
		OrganizationID:   "987",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestAccProjectsService_Get(t *testing.T) {
	if skipAcceptanceTestGet {
		t.Skip()
	}

	c, err := clients.LocalClient()
	if err != nil {
		t.Error(err)
	}
	projects := p.New(c)

	want := "6ea70fa9-af49-4550-80ad-8317788b4c4d"
	got, err := projects.Get(context.Background(), want)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != want {
		t.Errorf("got = %v, want %v", got.ID, want)
	}
}

func TestProjectsService_GetLifecycleState(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	projects := p.New(c)
	want := "some-state"

	mux.HandleFunc("/resource-management/v1/projects/abc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, err := json.Marshal(p.ProjectsResBody{
			LifecycleState: want,
		})
		if err != nil {
			log.Fatalf("json response marshal: %v", err)
		}
		fmt.Fprint(w, string(b))
	})

	got, err := projects.GetLifecycleState(context.Background(), "abc")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestProjectsService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	projects := p.New(c)

	mux.HandleFunc(fmt.Sprintf("/resource-management/v1/organizations/%s/projects", consts.SCHWARZ_ORGANIZATION_ID), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, err := json.Marshal(p.ProjectsResBody{
			ProjectID: "123",
			Name:      "Fancy-new-project",
			Labels: p.ProjectsLabelsResBody{
				"T-9876543B",
			},
			Parent: p.ProjectsParentResBody{
				ID: consts.SCHWARZ_ORGANIZATION_ID,
			},
		})
		if err != nil {
			log.Fatalf("json response marshal: %v", err)
		}
		fmt.Fprint(w, string(b))
	})

	role := p.ProjectRole{
		Name: "project.owner",
		Users: []p.ProjectRoleMember{
			{
				ID: "0d3a2fb9-1472-4284-9655-d7aae2cd5bd5",
			},
		},
	}
	got, err := projects.Create(context.Background(), "Fancy-new-project", "T-9876543B", role)
	if err != nil {
		t.Fatal(err)
	}

	want := p.Project{
		ID:               "123",
		Name:             "Fancy-new-project",
		BillingReference: "T-9876543B",
		OrganizationID:   consts.SCHWARZ_ORGANIZATION_ID,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestAccProjectsService_Create(t *testing.T) {
	if skipAcceptanceTestCreate {
		t.Skip()
	}

	c, err := clients.LocalClient()
	if err != nil {
		t.Error(err)
	}
	projects := p.New(c)

	want := p.Project{
		Name:             "test-odj-stackit-client-at-0",
		BillingReference: "T-0012253B",
	}
	role := p.ProjectRole{
		Name: "project.owner",
		Users: []p.ProjectRoleMember{
			{
				ID: "0d3a2fb9-1472-4284-9655-d7aae2cd5bd5",
			},
		},
	}
	got, err := projects.Create(context.Background(), want.Name, want.BillingReference, role)
	if err != nil {
		t.Fatal(err)
	}

	if got.Name != want.Name || got.BillingReference != want.BillingReference {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestProjectsService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	projects := p.New(c)

	mux.HandleFunc("/resource-management/v1/projects/abc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err := projects.Delete(context.Background(), "abc"); err != nil {
		t.Errorf("delete project: %v", err)
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
	projects := p.New(c)
	if err := projects.Delete(context.Background(), "0d3a2fb9-1472-4284-9655-d7aae2cd5bd5"); err != nil {
		t.Fatal(err)
	}
}

func TestProjectsService_Update(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	projects := p.New(c)

	mux.HandleFunc("/resource-management/v1/projects/123", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	err := projects.Update(context.Background(), "123", "Fancy-new-project", "T-9876543B")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccProjectsService_Update(t *testing.T) {
	if skipAcceptanceTestUpdate {
		t.Skip()
	}
	c, err := clients.LocalClient()
	if err != nil {
		t.Error(err)
	}
	projects := p.New(c)
	want := p.Project{
		ID:               "6ea70fa9-af49-4550-80ad-8317788b4c4d",
		Name:             "my-odj-test-project",
		BillingReference: "T-0012253B",
	}

	if err := projects.Update(context.Background(), want.ID, want.Name, want.BillingReference); err != nil {
		t.Fatal(err)
	}
}
