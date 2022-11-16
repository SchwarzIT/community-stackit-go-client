// package options is used to retrieve various options used for configuring MongoDB Flex
// Such as available versions and storage size

package options_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/mongodb-flex/options"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPathVersions = consts.API_PATH_MONGO_DB_FLEX_VERSIONS
	apiPathStorage  = consts.API_PATH_MONGO_DB_FLEX_STORAGES
	apiPathFlavors  = consts.API_PATH_MONGO_DB_FLEX_FLAVORS
)

func TestMongoDBOptionsService_GetVersions(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"

	mux.HandleFunc(fmt.Sprintf(apiPathVersions, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"versions": [
			  "string"
			]
		  }`)
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes options.VersionsResponse
		wantErr bool
	}{
		{"ok", args{context.Background(), projectID}, options.VersionsResponse{Versions: []string{"string"}}, false},
		{"nil ctx", args{nil, projectID}, options.VersionsResponse{Versions: []string{"string"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mongo.Options.GetVersions(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDBOptionsService.GetVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MongoDBOptionsService.GetVersions() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMongoDBOptionsService_GetFlavors(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"

	mux.HandleFunc(fmt.Sprintf(apiPathFlavors, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"flavors": [{
				"categories": [
					"string"
				],
				"cpu": 1,
				"memory": 2,
				"id": "string",
				"description": "string"
			}]
		  }`)
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes options.GetFlavorsResponse
		wantErr bool
	}{
		{"ok", args{context.Background(), projectID}, options.GetFlavorsResponse{
			Flavors: []options.Flavor{
				{
					Categories:  []string{"string"},
					CPU:         1,
					Memory:      2,
					ID:          "string",
					Description: "string",
				},
			},
		}, false},
		{"nil ctx", args{ctx: nil}, options.GetFlavorsResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mongo.Options.GetFlavors(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDBOptionsService.GetFlavors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MongoDBOptionsService.GetFlavors() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMongoDBOptionsService_GetStorage(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	mongo := mongodb.New(c)

	projectID := "abc"
	flavorID := "efd"

	mux.HandleFunc(fmt.Sprintf(apiPathStorage, projectID, flavorID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"storageClasses": [
			  "string"
			],
			"storageRange": {
			  "max": 20,
			  "min": 10
			}
		  }`)
	})

	type args struct {
		ctx       context.Context
		projectID string
		flavorID  string
	}
	tests := []struct {
		name    string
		args    args
		wantRes options.GetStorageResponse
		wantErr bool
	}{
		{"ok", args{context.Background(), projectID, flavorID}, options.GetStorageResponse{
			StorageClasses: []string{"string"},
			StorageRange: options.StorageRange{
				Max: 20,
				Min: 10,
			},
		}, false},
		{"nil ctx", args{ctx: nil}, options.GetStorageResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := mongo.Options.GetStorageClasses(tt.args.ctx, tt.args.projectID, tt.args.flavorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDBOptionsService.GetStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("MongoDBOptionsService.GetStorage() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
