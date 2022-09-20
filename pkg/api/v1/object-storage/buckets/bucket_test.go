package buckets_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage/buckets"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPath     = consts.API_PATH_OBJECT_STORAGE_BUCKET
	apiPathList = consts.API_PATH_OBJECT_STORAGE_BUCKETS
)

func TestObjectStorageBucketsService_List(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).Buckets

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	want := buckets.BucketListResponse{
		Project: projectID,
		Buckets: []buckets.Bucket{},
	}
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	mux.HandleFunc(fmt.Sprintf(apiPathList, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes buckets.BucketListResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID}, want, false},
		{"ctx is canceled", args{ctx, projectID}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.List(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageBucketsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("ObjectStorageBucketsService.List() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestObjectStorageBucketsService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).Buckets

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	bucketName := "my-bucket"
	want := buckets.BucketResponse{
		Project: projectID,
		Bucket: buckets.Bucket{
			Name: bucketName,
		},
	}
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	mux.HandleFunc(fmt.Sprintf(apiPath, projectID, bucketName), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	type args struct {
		ctx        context.Context
		projectID  string
		bucketName string
	}
	tests := []struct {
		name    string
		args    args
		wantRes buckets.BucketResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, bucketName}, want, false},
		{"ctx is canceled", args{ctx, projectID, bucketName}, want, true},
		{"bucket not found", args{context.Background(), projectID, "some-bucket"}, want, true},
		{"project not found", args{context.Background(), "some-id", bucketName}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.Get(tt.args.ctx, tt.args.projectID, tt.args.bucketName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageBucketsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("ObjectStorageBucketsService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestObjectStorageBucketsService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).Buckets

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	bucketName := "my-bucket"
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	mux.HandleFunc(fmt.Sprintf(apiPath, projectID, bucketName), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	type args struct {
		ctx        context.Context
		projectID  string
		bucketName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, bucketName}, false},
		{"ctx is canceled", args{ctx, projectID, bucketName}, true},
		{"project not found", args{context.Background(), "my-project", bucketName}, true},
		{"bucket not found", args{context.Background(), projectID, "some-bucket"}, true},
		{"bucket name invalid", args{context.Background(), projectID, "b"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Create(tt.args.ctx, tt.args.projectID, tt.args.bucketName); (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageBucketsService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestObjectStorageBucketsService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).Buckets

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	bucketName := "my-bucket"
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	mux.HandleFunc(fmt.Sprintf(apiPath, projectID, bucketName), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	type args struct {
		ctx        context.Context
		projectID  string
		bucketName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, bucketName}, false},
		{"ctx is canceled", args{ctx, projectID, bucketName}, true},
		{"project not found", args{context.Background(), "my-project", bucketName}, true},
		{"bucket not found", args{context.Background(), projectID, "some-bucket"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Delete(tt.args.ctx, tt.args.projectID, tt.args.bucketName); (err != nil) != tt.wantErr {
				t.Errorf("ObjectStorageBucketsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
