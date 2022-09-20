package options_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes/options"
)

var optionsExample = `{
	"kubernetesVersions": [
	  {
		"version": "string",
		"state": "string",
		"expirationDate": "2019-08-24T14:15:22Z",
		"featureGates": {
		  "property1": "string",
		  "property2": "string"
		}
	  }
	],
	"machineTypes": [
	  {
		"name": "string",
		"cpu": 0,
		"memory": 0
	  }
	],
	"machineImages": [
	  {
		"name": "string",
		"versions": [
		  {
			"version": "string",
			"state": "string",
			"expirationDate": "2019-08-24T14:15:22Z",
			"cri": [
			  {
				"name": "docker"
			  }
			]
		  }
		]
	  }
	],
	"volumeTypes": [
	  {
		"name": "string"
	  }
	],
	"availabilityZones": [
	  {
		"name": "string"
	  }
	]
  }`

func TestKubernetesOptionsService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	o := options.New(c)

	var want options.ProviderOptions
	if err := json.Unmarshal([]byte(optionsExample), &want); err != nil {
		t.Errorf(err.Error())
	}
	mux.HandleFunc("/ske/v1/provider-options", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, optionsExample)
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantRes options.ProviderOptions
		wantErr bool
	}{
		{"all ok", args{context.Background()}, want, false},
		{"ctx is canceled", args{ctx}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := o.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("KubernetesOptionsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("KubernetesOptionsService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
