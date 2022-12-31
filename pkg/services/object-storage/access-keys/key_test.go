package keys_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage"
	key "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/access-keys"
)

const (
	apiPathList      = consts.API_PATH_OBJECT_STORAGE_KEYS
	apiPathCreate    = consts.API_PATH_OBJECT_STORAGE_KEY
	apiPathDelete    = consts.API_PATH_OBJECT_STORAGE_WITH_KEY_ID
	credentialsGroup = "cd5e788d-5b7b-4ab9-a20d-e790205dabcd"
)

func TestStorageObjectKeyService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).AccessKeys

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"

	want := key.AccessKeyListResponse{
		Project:    projectID,
		AccessKeys: []key.AccessKeyDetails{},
	}

	mux.HandleFunc(fmt.Sprintf(apiPathList, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(b))
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx              context.Context
		projectID        string
		credentialsGroup string
	}
	tests := []struct {
		name    string
		args    args
		wantRes key.AccessKeyListResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, credentialsGroup}, want, false},
		{"all ok 2", args{context.Background(), projectID, ""}, want, false},
		{"ctx is canceled", args{ctx, projectID, credentialsGroup}, want, true},
		{"project not found", args{context.Background(), "my-project", credentialsGroup}, want, true},
		{"invalid credentials group", args{context.Background(), projectID, "invalid"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.List(tt.args.ctx, tt.args.projectID, tt.args.credentialsGroup)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageObjectKeyService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("StorageObjectKeyService.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestStorageObjectKeyService_Create(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).AccessKeys

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	expires := "2020-09-04T00:00:00.000Z"
	keyName := "my-key"
	keyID := "abcdefg"

	want := key.AccessKeyCreateResponse{
		Project:         projectID,
		DisplayName:     keyName,
		KeyID:           keyID,
		Expires:         expires,
		AccessKey:       "123456",
		SecretAccessKey: "098765",
	}

	mux.HandleFunc(fmt.Sprintf(apiPathCreate, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(b))
	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx              context.Context
		projectID        string
		credentialsGroup string
	}
	tests := []struct {
		name    string
		args    args
		wantRes key.AccessKeyCreateResponse
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, credentialsGroup}, want, false},
		{"all ok 2", args{context.Background(), projectID, ""}, want, false},
		{"ctx is canceled", args{ctx, projectID, credentialsGroup}, want, true},
		{"project not found", args{context.Background(), "my-project", credentialsGroup}, want, true},
		{"invalid credentials group", args{context.Background(), projectID, "invalid"}, want, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := s.Create(tt.args.ctx, tt.args.projectID, expires, tt.args.credentialsGroup)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageObjectKeyService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("StorageObjectKeyService.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestStorageObjectKeyService_Delete(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	s := objectstorage.New(c).AccessKeys

	projectID := "5dae0612-f5b1-4615-b7ca-b18796aa7e78"
	keyID := "abcdefg"

	mux.HandleFunc(fmt.Sprintf(apiPathDelete, projectID, keyID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	})

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx              context.Context
		projectID        string
		KeyID            string
		credentialsGroup string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"all ok", args{context.Background(), projectID, keyID, credentialsGroup}, false},
		{"all ok 2", args{context.Background(), projectID, keyID, ""}, false},
		{"ctx is canceled", args{ctx, projectID, keyID, credentialsGroup}, true},
		{"project not found", args{context.Background(), "my-project", keyID, credentialsGroup}, true},
		{"key not found", args{context.Background(), projectID, "some-key", credentialsGroup}, true},
		{"invalid credentials group", args{context.Background(), projectID, keyID, "invalid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Delete(tt.args.ctx, tt.args.projectID, tt.args.KeyID, tt.args.credentialsGroup); (err != nil) != tt.wantErr {
				t.Errorf("StorageObjectKeyService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
