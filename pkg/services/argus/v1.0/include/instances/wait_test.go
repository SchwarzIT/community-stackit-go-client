package instances

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/instances"
)

func TestInstanceCreateResponse_WaitHandler(t *testing.T) {
	ctx := context.Background()
	baseTime := time.Millisecond * 100
	type args struct {
		c InstanceReadWithResponse
	}
	tests := []struct {
		name    string
		args    args
		wantRes interface{}
		wantErr error
	}{
		{"call error", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return nil, errors.New("call error")
			},
		}, nil, errors.New("defined wait function returned an error: call error")},
		{"connection reset", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return nil, errors.New(connection_reset)
			},
		}, nil, errors.New("Wait() has timed out")},
		{"status is StatusInternalServerError", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return &instances.InstanceReadResponse{HTTPResponse: &http.Response{StatusCode: http.StatusInternalServerError}}, nil
			},
		}, nil, errors.New("Wait() has timed out")},
		{"status is StatusBadGateway", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return &instances.InstanceReadResponse{HTTPResponse: &http.Response{StatusCode: http.StatusBadGateway}}, nil
			},
		}, nil, errors.New("Wait() has timed out")},
		{"HasError != nil", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return &instances.InstanceReadResponse{
					HTTPResponse: &http.Response{StatusCode: http.StatusOK},
					HasError:     errors.New("some error"),
				}, nil
			},
		}, nil, errors.New("defined wait function returned an error: some error")},
		{"JSON200 == nil", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return &instances.InstanceReadResponse{
					HTTPResponse: &http.Response{StatusCode: http.StatusOK},
					JSON200:      nil,
				}, nil
			},
		}, nil, errors.New("defined wait function returned an error: received an empty response. JSON200 == nil")},
		{"JSON200.status == PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return &instances.InstanceReadResponse{
					HTTPResponse: &http.Response{StatusCode: http.StatusOK},
					JSON200: &instances.ProjectInstanceUI{
						Status: instances.PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED,
					},
				}, nil
			},
		}, &instances.ProjectInstanceUI{
			Status: instances.PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED,
		}, nil},
		{"JSON200.status != PROJECT_INSTANCE_UI_STATUS_CREATE_SUCCEEDED", args{
			c: func(ctx context.Context, projectID, instanceID string, reqEditors ...instances.RequestEditorFn) (*instances.InstanceReadResponse, error) {
				return &instances.InstanceReadResponse{
					HTTPResponse: &http.Response{StatusCode: http.StatusOK},
					JSON200: &instances.ProjectInstanceUI{
						Status: instances.PROJECT_INSTANCE_UI_STATUS_CREATING,
					},
				}, nil
			},
		}, &instances.ProjectInstanceUI{
			Status: instances.PROJECT_INSTANCE_UI_STATUS_CREATING,
		}, errors.New("Wait() has timed out")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InstanceCreateResponse{}
			got := r.WaitHandler(ctx, tt.args.c, "", "")
			got.SetThrottle(baseTime)
			got.SetTimeout(baseTime)
			res, err := got.WaitWithContext(ctx)
			if !reflect.DeepEqual(res, tt.wantRes) {
				t.Errorf("response = %v, want %v", got, tt.wantRes)
			}
			if fmt.Sprintf("%s", err) != fmt.Sprintf("%s", tt.wantErr) {
				t.Errorf("err = %s, want %s", err, tt.wantErr)
			}
		})
	}
}
