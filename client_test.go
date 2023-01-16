package client

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func TestNew(t *testing.T) {
	cfg := Config{
		ServiceAccountToken: "token",
		ServiceAccountEmail: "sa-id",
	}
	type args struct {
		ctx context.Context
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{"no token", args{context.Background(), Config{}}, &Client{}, true},
		{"no sa id", args{context.Background(), Config{ServiceAccountToken: "token"}}, &Client{}, true},
		{"all ok", args{context.Background(), cfg}, &Client{config: cfg}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ctx, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (!reflect.DeepEqual(got.config.ServiceAccountEmail, tt.want.config.ServiceAccountEmail) || !reflect.DeepEqual(got.config.ServiceAccountToken, tt.want.config.ServiceAccountToken)) {
				t.Errorf("NewAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Request(t *testing.T) {
	cfg := Config{
		ServiceAccountToken: "token",
		ServiceAccountEmail: "sa-id",
	}
	c, err := New(context.Background(), cfg)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	type args struct {
		ctx    context.Context
		method string
		path   string
		body   string
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{"canceled context", args{ctx, "something", "my-path", ""}, &http.Request{}, true},
		{"bad context", args{nil, "something", "my-path", ""}, &http.Request{}, true},
		{"bad method", args{context.Background(), "something", "my-path", ""}, &http.Request{}, true},
		{"all ok", args{context.Background(), http.MethodGet, "my-path", ""}, &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/my-path"}}, false},
		{"all ok with url params", args{context.Background(), http.MethodGet, "my-path?a=b", ""}, &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/my-path"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Request(tt.args.ctx, tt.args.method, tt.args.path, []byte(tt.args.body))
			if !tt.wantErr {
				if err != nil {
					t.Error(err)
				}
				if tt.want.Method != got.Method {
					t.Error("wrong method")
				}
				if tt.want.URL.Path != got.URL.Path {
					t.Error("wrong url path", tt.want.URL.Path, got.URL.Path)
				}
			}
		})
	}
}

func TestClient_GetHTTPClient(t *testing.T) {
	type fields struct {
		client   *http.Client
		config   Config
		Services services
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Client
	}{
		{"all ok", fields{client: &http.Client{}}, &http.Client{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:   tt.fields.client,
				config:   tt.fields.config,
				services: tt.fields.Services,
			}
			if got := c.GetHTTPClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetHTTPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetConfig(t *testing.T) {
	type fields struct {
		client   *http.Client
		config   Config
		Services services
	}
	tests := []struct {
		name   string
		fields fields
		want   Config
	}{
		{"all ok", fields{config: Config{}}, Config{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:   tt.fields.client,
				config:   tt.fields.config,
				services: tt.fields.Services,
			}
			if got := c.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetToken(t *testing.T) {
	type fields struct {
		client   *http.Client
		config   Config
		Services services
	}
	type args struct {
		token string
	}
	c := Config{ServiceAccountToken: "abc"}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"all ok", fields{config: c}, args{token: c.ServiceAccountToken}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:   tt.fields.client,
				config:   tt.fields.config,
				services: tt.fields.Services,
			}
			c.SetToken(tt.args.token)
		})
	}
}

func TestClient_GeneralTests(t *testing.T) {
	c, _, teardown, err := MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	basetime := 200 * time.Millisecond
	c.RetryTimout = 5 * basetime
	c.RetryWait = basetime

	c.SetBaseURL("ht#@://aa")
	req, _ := c.Request(context.Background(), http.MethodGet, "err2", nil)
	if _, err := c.Do(req); err == nil {
		t.Error("expected do request to return error")
	}

	c.SetBaseURL(common.DEFAULT_BASE_URL)
	if c.GetBaseURL() != common.DEFAULT_BASE_URL {
		t.Error("bad url")
	}
	cfg := c.GetConfig()
	if cfg.BaseUrl.String() != common.DEFAULT_BASE_URL {
		t.Errorf("expected base URL to be %s, got %s instead", common.DEFAULT_BASE_URL, cfg.BaseUrl.String())
	}
}

func TestClient_DoWithRetry(t *testing.T) {
	c, mux, teardown, err := MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	basetime := 200 * time.Millisecond
	ctx := context.Background()
	ctx1, td1 := context.WithTimeout(ctx, 2*basetime)
	defer td1()

	ctx2, td2 := context.WithTimeout(ctx, 3*basetime)
	defer td2()

	c.RetryTimout = 5 * basetime
	c.RetryWait = basetime

	mux.HandleFunc("/ep", func(w http.ResponseWriter, r *http.Request) {
		if ctx1.Err() == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if ctx2.Err() == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("{}"))
	})

	req, _ := c.Request(context.Background(), http.MethodGet, "/ep", nil)
	if _, err := c.Do(req); err != nil {
		t.Error(err)
	}

	ctx3, td3 := context.WithTimeout(ctx, 6*basetime)
	defer td3()

	c.client.Timeout = basetime * 1
	mux.HandleFunc("/ep2", func(w http.ResponseWriter, r *http.Request) {

		if ctx3.Err() == nil {
			time.Sleep(1 * basetime)
		}
	})

	req, _ = c.Request(context.Background(), http.MethodGet, "/ep2", nil)

	if _, err := c.Do(req); err != nil {
		t.Error(err)
	}
}
