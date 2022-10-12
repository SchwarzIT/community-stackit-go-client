package buckets_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SchwarzIT/community-stackit-go-client"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage/buckets"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// constants
const (
	apiPath     = consts.API_PATH_OBJECT_STORAGE_BUCKET
	apiPathList = consts.API_PATH_OBJECT_STORAGE_BUCKETS
)

// getStatus is used in test arguments to cover the waitHandler cases
type getStatus int

const (
	notFound = iota
	ok
	internalServerError
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
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1()
	cancelledCtx, cancelTmp := context.WithCancel(context.TODO())
	cancelTmp()

	type args struct {
		ctx        context.Context
		projectID  string
		bucketName string
	}
	tests := []struct {
		name      string
		args      args
		getStatus getStatus
		wantErr   bool
	}{
		{name: "all ok", args: args{context.Background(), projectID, bucketName}, getStatus: ok},
		{name: "ctx is canceled", args: args{cancelledCtx, projectID, bucketName}, getStatus: ok, wantErr: true},
		{name: "project not found", args: args{context.Background(), "my-project", bucketName}, getStatus: ok, wantErr: true},
		{name: "bucket not found", args: args{context.Background(), projectID, "some-bucket"}, getStatus: ok, wantErr: true},
		{name: "bucket name invalid", args: args{context.Background(), projectID, "b"}, getStatus: ok, wantErr: true},
		{name: "err from get", args: args{context.Background(), projectID, bucketName}, getStatus: internalServerError, wantErr: true},
		{name: "not found from get", args: args{ctx1, projectID, bucketName}, getStatus: notFound, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf(apiPath, projectID, bucketName), func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodGet:
					w.Header().Set("Content-Type", "application/json")
					switch tt.getStatus {
					case ok:
						w.WriteHeader(http.StatusOK)
						b, _ := json.Marshal(buckets.Bucket{})
						fmt.Fprint(w, string(b))
					case notFound:
						w.WriteHeader(http.StatusNotFound)
					case internalServerError:
						w.WriteHeader(http.StatusInternalServerError)
					}
					return
				case http.MethodPost:
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					return
				default:
					t.Errorf("expected method %s, got %s", http.MethodPost, r.Method)
				}
			})

			process, err := s.Create(tt.args.ctx, tt.args.projectID, tt.args.bucketName)
			hasErr := err != nil
			if !hasErr {
				process.SetThrottle(1 * time.Second)
				_, err = process.Wait()
			}
			hasErr = hasErr || err != nil
			if hasErr != tt.wantErr {
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
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1()
	cancelledCtx, cancelTmp := context.WithCancel(context.TODO())
	cancelTmp()

	type args struct {
		ctx        context.Context
		projectID  string
		bucketName string
	}
	tests := []struct {
		name      string
		args      args
		getStatus getStatus
		wantErr   bool
	}{
		{name: "all ok", args: args{ctx1, projectID, bucketName}},
		{name: "ctx is canceled", args: args{cancelledCtx, projectID, bucketName}, wantErr: true},
		{name: "project not found", args: args{context.Background(), "my-project", bucketName}, wantErr: true},
		{name: "bucket not found", args: args{context.Background(), projectID, "some-bucket"}, wantErr: true},
		{name: "error from get in wait", args: args{context.Background(), projectID, bucketName}, getStatus: internalServerError, wantErr: true},
		{name: "ok from get in wait", args: args{ctx1, projectID, bucketName}, getStatus: ok, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf(apiPath, projectID, bucketName), func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodGet:
					w.Header().Set("Content-Type", "application/json")
					switch tt.getStatus {
					case notFound:
						w.WriteHeader(http.StatusNotFound)
					case ok:
						w.WriteHeader(http.StatusOK)
						b, _ := json.Marshal(buckets.Bucket{})
						fmt.Fprint(w, string(b))
					case internalServerError:
						w.WriteHeader(http.StatusInternalServerError)
					}
					return
				case http.MethodDelete:
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					return
				default:
					t.Errorf("expected method %s, got %s", http.MethodPost, r.Method)
				}
			})
			process, err := s.Delete(tt.args.ctx, tt.args.projectID, tt.args.bucketName)
			hasErr := err != nil
			if !hasErr {
				process.SetThrottle(1 * time.Second)
				_, err = process.Wait()
			}
			hasErr = hasErr || err != nil
			if hasErr != tt.wantErr {
				t.Errorf("ObjectStorageBucketsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
