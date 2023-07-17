package clients

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenFlow_processConfig(t *testing.T) {
	// test env variable loading
	a := os.Getenv(ServiceAccountEmail)
	b := os.Getenv(ServiceAccountToken)

	os.Setenv(ServiceAccountEmail, "test 1")
	os.Setenv(ServiceAccountToken, "test 2")

	rc := NewRetryConfig()
	tf := &TokenFlow{}
	tf.processConfig()
	tf.config.ClientRetry = rc

	want := TokenFlowConfig{
		ServiceAccountEmail: "test 1",
		ServiceAccountToken: "test 2",
		ClientRetry:         rc,
	}
	assert.EqualValues(t, want, *tf.config)

	// revert
	os.Setenv(ServiceAccountEmail, a)
	os.Setenv(ServiceAccountToken, b)

	// Test manual configuration
	type args struct {
		cfg []TokenFlowConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{"test manual 1", args{[]TokenFlowConfig{
			{ServiceAccountEmail: "test 1", ClientRetry: rc, Traceparent: false},
			{ServiceAccountToken: "test 2", ClientRetry: rc, Traceparent: false},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TokenFlow{}
			c.config = &TokenFlowConfig{
				ClientRetry: rc,
			}
			c.processConfig(tt.args.cfg...)
			assert.Equal(t, want, c.GetConfig())
		})
	}
}

func TestTokenFlow_Init(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg []TokenFlowConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{context.Background(), []TokenFlowConfig{
			{
				ServiceAccountEmail: "abc",
				ServiceAccountToken: "efg",
			},
		}}, false},
		{"error 1", args{context.Background(), []TokenFlowConfig{
			{
				ServiceAccountEmail: "",
				ServiceAccountToken: "",
			},
		}}, true},
		{"error 2", args{context.Background(), []TokenFlowConfig{
			{
				ServiceAccountEmail: "",
				ServiceAccountToken: "efg",
			},
		}}, true},
	}
	a := os.Getenv(ServiceAccountEmail)
	b := os.Getenv(ServiceAccountToken)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TokenFlow{}

			os.Setenv(ServiceAccountEmail, "")
			os.Setenv(ServiceAccountToken, "")
			if err := c.Init(tt.args.ctx, tt.args.cfg...); (err != nil) != tt.wantErr {
				t.Errorf("TokenFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			os.Setenv(ServiceAccountEmail, a)
			os.Setenv(ServiceAccountToken, b)
			if c.config == nil {
				t.Error("config is nil")
			}
			assert.EqualValues(t, c.config.ServiceAccountEmail, c.GetServiceAccountEmail())
		})
	}
}

func TestTokenFlow_Do(t *testing.T) {
	type fields struct {
		client *http.Client
		config *TokenFlowConfig
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"fail", fields{nil, nil}, args{}, 0, true},
		{"success", fields{&http.Client{}, &TokenFlowConfig{ClientRetry: &RetryConfig{}}}, args{}, http.StatusOK, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TokenFlow{
				client: tt.fields.client,
				config: tt.fields.config,
			}
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"status":"ok"}`)
			})
			server := httptest.NewServer(handler)
			defer server.Close()
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Error(err)
				return
			}
			req := &http.Request{
				URL: u,
			}
			got, err := c.Do(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenFlow.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.StatusCode != tt.want {
				t.Errorf("TokenFlow.Do() = %v, want %v", got.StatusCode, tt.want)
			}
		})
	}
}

func TestTokenClone(t *testing.T) {
	c := &TokenFlow{
		client: &http.Client{},
		config: &TokenFlowConfig{},
	}

	clone := c.Clone().(*TokenFlow)

	if !reflect.DeepEqual(c, clone) {
		t.Errorf("Clone() = %v, want %v", clone, c)
	}
}
