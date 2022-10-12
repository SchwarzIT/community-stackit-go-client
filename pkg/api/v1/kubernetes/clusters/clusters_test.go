package clusters_test

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
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes/clusters"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

var (
	k_bad_0 = clusters.Kubernetes{}
	k_bad_1 = clusters.Kubernetes{Version: "a$.b%@^*&"}
	k_ok    = clusters.Kubernetes{Version: "1.2.3"}

	np_bad_0 = []clusters.NodePool{}
	np_bad_1 = []clusters.NodePool{{
		Machine: clusters.Machine{Type: "type"},
	}}
	np_bad_2 = []clusters.NodePool{{
		Machine: clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
	}}
	np_bad_3 = []clusters.NodePool{{
		Machine: clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum: 7,
		Maximum: 3,
	}}
	np_bad_4 = []clusters.NodePool{{
		Machine: clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum: 0,
		Maximum: 3,
	}}
	np_bad_5 = []clusters.NodePool{{
		Machine: clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum: 1,
		Maximum: 101,
	}}
	np_bad_6 = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  10,
		MaxSurge: 0,
	}}
	np_bad_7 = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  10,
		MaxSurge: 2,
	}}
	np_bad_8 = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  10,
		MaxSurge: 2,
		Volume:   clusters.Volume{Size: 30},
	}}
	np_bad_9 = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  10,
		MaxSurge: 2,
		Volume:   clusters.Volume{Size: 30},
		Taints:   []clusters.Taint{{Effect: "random"}},
	}}
	np_bad_10 = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  10,
		MaxSurge: 2,
		Volume:   clusters.Volume{Size: 30},
		Taints:   []clusters.Taint{{Effect: consts.SKE_CLUSTERS_TAINT_EFFECT_NO_EXEC}},
		CRI:      clusters.CRI{Name: "dockers"},
	}}
	np_bad_11 = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  20,
		MaxSurge: 2,
		Volume:   clusters.Volume{Size: 30},
		Taints:   []clusters.Taint{{Effect: consts.SKE_CLUSTERS_TAINT_EFFECT_NO_EXEC, Key: "something"}},
		CRI:      clusters.CRI{Name: "containerd"},
	}}
	np_bad_12 = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  200,
		MaxSurge: 2,
		Volume:   clusters.Volume{Size: 30},
		Taints:   []clusters.Taint{{Effect: consts.SKE_CLUSTERS_TAINT_EFFECT_NO_EXEC, Key: "something"}},
		CRI:      clusters.CRI{Name: "containerd"},
	}}
	np_ok = []clusters.NodePool{{
		Machine:  clusters.Machine{Type: "type", Image: clusters.MachineImage{Version: "abc"}},
		Minimum:  1,
		Maximum:  10,
		MaxSurge: 2,
		Volume:   clusters.Volume{Size: 30},
		Taints:   []clusters.Taint{{Effect: consts.SKE_CLUSTERS_TAINT_EFFECT_NO_EXEC, Key: "something"}},
		CRI:      clusters.CRI{Name: "containerd"},
	}}

	m_bad_1 = &clusters.Maintenance{}
	m_bad_2 = &clusters.Maintenance{
		clusters.MaintenanceAutoUpdate{},
		clusters.MaintenanceTimeWindow{
			Start: "some date..",
		},
	}
	m_bad_3 = &clusters.Maintenance{
		clusters.MaintenanceAutoUpdate{},
		clusters.MaintenanceTimeWindow{
			End: "some date..",
		},
	}
	m_ok = &clusters.Maintenance{
		clusters.MaintenanceAutoUpdate{},
		clusters.MaintenanceTimeWindow{
			Start: "some date..",
			End:   "some other date",
		},
	}

	h_ok    = &clusters.Hibernation{}
	h_bad_1 = &clusters.Hibernation{Schedules: []clusters.HibernationScedule{{
		Start: "something",
	}}}
	h_bad_2 = &clusters.Hibernation{Schedules: []clusters.HibernationScedule{{
		End: "something",
	}}}

	e_ok    = &clusters.Extensions{}
	e_bad_1 = &clusters.Extensions{Argus: &clusters.ArgusExtension{Enabled: true, ArgusInstanceID: ""}}
)

func TestKubernetesClusterService_CreateOrUpdate(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	clusterName := "cname"

	want := clusters.Cluster{
		Name:        clusterName,
		Kubernetes:  k_ok,
		Nodepools:   np_ok,
		Maintenance: m_ok,
		Hibernation: h_ok,
	}

	get1 := clusters.Cluster{
		Status: &clusters.Status{
			Aggregated: consts.SKE_CLUSTER_STATUS_HEALTHY,
		},
	}

	get2 := clusters.Cluster{
		Status: &clusters.Status{
			Aggregated: consts.SKE_CLUSTER_STATUS_CREATING,
		},
	}

	ctx1, cancel1 := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel1()

	ctx2, cancel2 := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel2()

	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
			return
		}
		if r.Method == http.MethodGet {
			if ctx1.Err() == nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			if ctx2.Err() == nil {
				b, _ := json.Marshal(get2)
				fmt.Fprint(w, string(b))
				return
			}

			b, _ := json.Marshal(get1)
			fmt.Fprint(w, string(b))
			return
		}
		t.Error("wrong method")
	}
	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName, fn)
	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/C_N-AME", fn)

	ctx := context.Background()
	ctx_bad, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx           context.Context
		projectID     string
		clusterName   string
		clusterConfig clusters.Kubernetes
		nodePools     []clusters.NodePool
		maintenance   *clusters.Maintenance
		hibernation   *clusters.Hibernation
		extensions    *clusters.Extensions
	}
	tests := []struct {
		name    string
		args    args
		wantRes clusters.Cluster
		wantErr bool
		useWait bool
	}{
		{"ctx is canceled", args{ctx_bad, projectID, clusterName, k_ok, np_ok, m_ok, h_ok, e_ok}, want, true, false},
		{"bad project ID", args{ctx, "something", clusterName, k_ok, np_ok, m_ok, h_ok, e_ok}, want, true, false},
		{"bad cluster name", args{ctx, projectID, "C_N-AME", k_ok, np_ok, m_ok, h_ok, e_ok}, want, true, false},
		{"all ok", args{ctx, projectID, clusterName, k_ok, np_ok, m_ok, h_ok, e_ok}, want, false, true},
		{"all ok 2", args{ctx, projectID, clusterName, k_ok, np_ok, nil, h_ok, e_ok}, want, false, false},
		{"all ok 3", args{ctx, projectID, clusterName, k_ok, np_ok, nil, nil, e_ok}, want, false, false},
		{"all ok 4", args{ctx, projectID, clusterName, k_ok, np_ok, nil, nil, nil}, want, false, false},

		{"kube bad 0", args{ctx, projectID, clusterName, k_bad_0, np_ok, m_ok, h_ok, e_ok}, want, true, false},
		{"kube bad 1", args{ctx, projectID, clusterName, k_bad_1, np_ok, m_ok, h_ok, e_ok}, want, true, false},

		{"np bad 0", args{ctx, projectID, clusterName, k_ok, np_bad_0, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 1", args{ctx, projectID, clusterName, k_ok, np_bad_1, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 2", args{ctx, projectID, clusterName, k_ok, np_bad_2, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 3", args{ctx, projectID, clusterName, k_ok, np_bad_3, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 4", args{ctx, projectID, clusterName, k_ok, np_bad_4, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 5", args{ctx, projectID, clusterName, k_ok, np_bad_5, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 6", args{ctx, projectID, clusterName, k_ok, np_bad_6, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 7", args{ctx, projectID, clusterName, k_ok, np_bad_7, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 8", args{ctx, projectID, clusterName, k_ok, np_bad_8, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 9", args{ctx, projectID, clusterName, k_ok, np_bad_9, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 10", args{ctx, projectID, clusterName, k_ok, np_bad_10, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 11", args{ctx, projectID, clusterName, k_ok, np_bad_11, m_ok, h_ok, e_ok}, want, true, false},
		{"np bad 12", args{ctx, projectID, clusterName, k_ok, np_bad_12, m_ok, h_ok, e_ok}, want, true, false},

		{"maintenance bad 1", args{ctx, projectID, clusterName, k_ok, np_ok, m_bad_1, h_ok, e_ok}, want, true, false},
		{"maintenance bad 2", args{ctx, projectID, clusterName, k_ok, np_ok, m_bad_2, h_ok, e_ok}, want, true, false},
		{"maintenance bad 3", args{ctx, projectID, clusterName, k_ok, np_ok, m_bad_3, h_ok, e_ok}, want, true, false},

		{"h bad 1", args{ctx, projectID, clusterName, k_ok, np_ok, m_ok, h_bad_1, e_ok}, want, true, false},
		{"h bad 2", args{ctx, projectID, clusterName, k_ok, np_ok, m_ok, h_bad_2, e_ok}, want, true, false},

		{"e bad 1", args{ctx, projectID, clusterName, k_ok, np_ok, m_ok, h_ok, e_bad_1}, want, true, false},
	}

	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test create
			{
				gotRes, w, err := s.Clusters.CreateOrUpdate(tt.args.ctx, tt.args.projectID, tt.args.clusterName, tt.args.clusterConfig, tt.args.nodePools, tt.args.maintenance, tt.args.hibernation, tt.args.extensions)
				if (err != nil) != tt.wantErr {
					t.Errorf("KubernetesClusterService.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
					t.Errorf("KubernetesClusterService.Create() = %v, want %v", gotRes, tt.wantRes)
				}

				if tt.useWait {
					process = w
				}
			}
		})
	}

	process.SetThrottle(1 * time.Second)

	// during the initial call, the server is supposed to return an error
	// that should force the Wait() to exit immediately
	if _, err := process.Wait(); err == nil {
		t.Errorf("expected error during first process wait, got nil instead")
	}

	// after 1s the server should return a retryable error and than
	// change to the wanted response status
	// meaning - Wait() should exit without any error
	time.Sleep(1 * time.Second)
	if _, err := process.Wait(); err != nil {
		t.Errorf("unexpected error during 2nd process wait: %s", err)
	}
}

func TestKubernetesClusterService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	clusterName := "cname"
	want := clusters.ClusterList{
		Items: []clusters.Cluster{{
			Name:        clusterName,
			Kubernetes:  k_ok,
			Nodepools:   np_ok,
			Maintenance: m_ok,
			Hibernation: h_ok,
		}},
	}

	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	ctx := context.Background()
	ctx_bad, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes clusters.ClusterList
		wantErr bool
	}{
		{"ctx is canceled", args{ctx_bad, projectID}, want, true},
		{"bad project ID", args{ctx, "something"}, want, true},
		{"all ok", args{ctx, projectID}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.Clusters.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesClusterService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("KubernetesClusterService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestKubernetesClusterService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	clusterName := "cname"

	want := clusters.Cluster{
		Name:        clusterName,
		Kubernetes:  k_ok,
		Nodepools:   np_ok,
		Maintenance: m_ok,
		Hibernation: h_ok,
	}

	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	ctx := context.Background()
	ctx_bad, cancel := context.WithCancel(context.TODO())
	cancel()
	type args struct {
		ctx         context.Context
		projectID   string
		clusterName string
	}
	tests := []struct {
		name    string
		args    args
		wantRes clusters.Cluster
		wantErr bool
	}{
		{"ctx is canceled", args{ctx_bad, projectID, clusterName}, want, true},
		{"bad project ID", args{ctx, "something", clusterName}, want, true},
		{"bad cluster name", args{ctx, projectID, "something"}, want, true},
		{"all ok", args{ctx, projectID, clusterName}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.Clusters.Get(tt.args.ctx, tt.args.projectID, tt.args.clusterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesClusterService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("KubernetesClusterService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestKubernetesClusterService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	clusterName := "cname"

	get1 := clusters.Cluster{
		Status: &clusters.Status{
			Aggregated: consts.SKE_CLUSTER_STATUS_HEALTHY,
		},
	}

	ctx1, cancel1 := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel1()

	ctx2, cancel2 := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel2()

	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			if ctx1.Err() == nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if ctx2.Err() == nil {
				w.WriteHeader(http.StatusOK)
				b, _ := json.Marshal(get1)
				fmt.Fprint(w, string(b))
				return
			}

			w.WriteHeader(http.StatusNotFound)
			return
		}
		t.Error("wrong method")
	})

	ctx := context.Background()
	ctx_bad, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx         context.Context
		projectID   string
		clusterName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		useWait bool
	}{
		{"ctx is canceled", args{ctx_bad, projectID, clusterName}, true, false},
		{"bad project ID", args{ctx, "something", clusterName}, true, false},
		{"bad cluster name", args{ctx, projectID, "something"}, true, false},
		{"all ok", args{ctx, projectID, clusterName}, false, true},
	}
	var process *wait.Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := s.Clusters.Delete(tt.args.ctx, tt.args.projectID, tt.args.clusterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesClusterService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.useWait {
				process = w
			}
		})
	}

	// during the initial call, the server is supposed to return an error
	// that should force the Wait() to exist immediately
	if _, err := process.Wait(); err == nil {
		t.Error("expected error in first attempt, but got nil instead")
	}

	// after 1s the server should return a status of an active cluster
	// and after another wait run return status Not Found -> cluster deleted
	// meaning - Wait() should exit without any error
	time.Sleep(1 * time.Second)
	if _, err := process.Wait(); err != nil {
		t.Errorf("unexpected error during 2nd wait: %v", err)
	}
}

func TestKubernetesClusterService_Triggers(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	clusterName := "cname"

	want := clusters.Cluster{
		Name:        clusterName,
		Kubernetes:  k_ok,
		Nodepools:   np_ok,
		Maintenance: m_ok,
		Hibernation: h_ok,
	}

	fixedRespFn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	}

	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName+"/hibernate", fixedRespFn)
	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName+"/maintenance", fixedRespFn)
	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName+"/reconcile", fixedRespFn)
	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName+"/wakeup", fixedRespFn)

	ctx := context.Background()
	ctx_bad, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx         context.Context
		projectID   string
		clusterName string
	}
	tests := []struct {
		name    string
		args    args
		wantRes clusters.Cluster
		wantErr bool
	}{
		{"ctx is canceled", args{ctx_bad, projectID, clusterName}, want, true},
		{"bad project ID", args{ctx, "something", clusterName}, want, true},
		{"bad cluster name", args{ctx, projectID, "something"}, want, true},
		{"all ok", args{ctx, projectID, clusterName}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Hibernate
			{
				gotRes, err := s.Clusters.Hibernate(tt.args.ctx, tt.args.projectID, tt.args.clusterName)
				if (err != nil) != tt.wantErr {
					t.Errorf("KubernetesClusterService.Hibernate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
					t.Errorf("KubernetesClusterService.Hibernate() = %v, want %v", gotRes, tt.wantRes)
				}
			}
			// maintenance
			{
				gotRes, err := s.Clusters.Maintenance(tt.args.ctx, tt.args.projectID, tt.args.clusterName)
				if (err != nil) != tt.wantErr {
					t.Errorf("KubernetesClusterService.Maintenance() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
					t.Errorf("KubernetesClusterService.Maintenance() = %v, want %v", gotRes, tt.wantRes)
				}
			}
			// Reconcile
			{
				gotRes, err := s.Clusters.Reconcile(tt.args.ctx, tt.args.projectID, tt.args.clusterName)
				if (err != nil) != tt.wantErr {
					t.Errorf("KubernetesClusterService.Reconcile() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
					t.Errorf("KubernetesClusterService.Reconcile() = %v, want %v", gotRes, tt.wantRes)
				}
			}
			// Wakeup
			{
				gotRes, err := s.Clusters.Wakeup(tt.args.ctx, tt.args.projectID, tt.args.clusterName)
				if (err != nil) != tt.wantErr {
					t.Errorf("KubernetesClusterService.Wakeup() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
					t.Errorf("KubernetesClusterService.Wakeup() = %v, want %v", gotRes, tt.wantRes)
				}
			}
		})
	}
}

func TestKubernetesClusterService_GetCredential(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	clusterName := "cname"

	want := clusters.Credentials{}

	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName+"/credentials", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	ctx := context.Background()
	ctx_bad, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx         context.Context
		projectID   string
		clusterName string
	}
	tests := []struct {
		name    string
		args    args
		wantRes clusters.Credentials
		wantErr bool
	}{
		{"ctx is canceled", args{ctx_bad, projectID, clusterName}, want, true},
		{"bad project ID", args{ctx, "something", clusterName}, want, true},
		{"bad cluster name", args{ctx, projectID, "something"}, want, true},
		{"all ok", args{ctx, projectID, clusterName}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.Clusters.GetCredential(tt.args.ctx, tt.args.projectID, tt.args.clusterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesClusterService.GetCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("KubernetesClusterService.GetCredential() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestKubernetesClusterService_RotateCredentials(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := kubernetes.New(c)

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	clusterName := "cname"
	mux.HandleFunc("/ske/v1/projects/"+projectID+"/clusters/"+clusterName+"/rotate-credentials", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	ctx_bad, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx         context.Context
		projectID   string
		clusterName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ctx is canceled", args{ctx_bad, projectID, clusterName}, true},
		{"bad project ID", args{ctx, "something", clusterName}, true},
		{"bad cluster name", args{ctx, projectID, "something"}, true},
		{"all ok", args{ctx, projectID, clusterName}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Clusters.RotateCredentials(tt.args.ctx, tt.args.projectID, tt.args.clusterName); (err != nil) != tt.wantErr {
				t.Errorf("KubernetesClusterService.RotateCredentials() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
