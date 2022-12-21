package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
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

func TestClient_Do(t *testing.T) {
	type Test struct {
		Payload string `json:"payload,omitempty"`
	}

	c, mux, teardown, err := MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprint(w, string(b))
	})

	p := Test{Payload: "my-payload"}
	body, err := json.Marshal(p)
	if err != nil {
		t.Errorf("marshal err: %v", err)
		return
	}

	req, err := c.Request(context.Background(), http.MethodPost, "/echo", body)
	if err != nil {
		t.Errorf("new request: %v", err)
		return
	}

	var got Test
	_, err = c.LegacyDo(req, &got)
	if err != nil {
		t.Errorf("do request: %v", err)
	}

	want := p
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}

	req.URL.Path = "/blah"
	if _, err = c.LegacyDo(req, &got); err == nil {
		t.Error("expected error over path")
	}

	ctx2, cancel := context.WithTimeout(context.TODO(), 0)
	defer cancel()
	req = req.WithContext(ctx2)
	if _, err = c.LegacyDo(req, &got); err == nil {
		t.Error("expected error over context timeout")
	}

}

func TestClient_GetHTTPClient(t *testing.T) {
	type fields struct {
		client             *http.Client
		config             Config
		ProductiveServices ProductiveServices
		Incubator          IncubatorServices
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
				client:             tt.fields.client,
				config:             tt.fields.config,
				ProductiveServices: tt.fields.ProductiveServices,
				Incubator:          tt.fields.Incubator,
			}
			if got := c.GetHTTPClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetHTTPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetConfig(t *testing.T) {
	type fields struct {
		client             *http.Client
		config             Config
		ProductiveServices ProductiveServices
		Incubator          IncubatorServices
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
				client:             tt.fields.client,
				config:             tt.fields.config,
				ProductiveServices: tt.fields.ProductiveServices,
				Incubator:          tt.fields.Incubator,
			}
			if got := c.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetToken(t *testing.T) {
	type fields struct {
		client             *http.Client
		config             Config
		ProductiveServices ProductiveServices
		Incubator          IncubatorServices
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
				client:             tt.fields.client,
				config:             tt.fields.config,
				ProductiveServices: tt.fields.ProductiveServices,
				Incubator:          tt.fields.Incubator,
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

	c.SetBaseURL(consts.DEFAULT_BASE_URL)
	if c.GetBaseURL() != consts.DEFAULT_BASE_URL {
		t.Error("bad url")
	}
	cfg := c.GetConfig()
	if cfg.BaseUrl.String() != consts.DEFAULT_BASE_URL {
		t.Errorf("expected base URL to be %s, got %s instead", consts.DEFAULT_BASE_URL, cfg.BaseUrl.String())
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

	c.RetryTimout = 5 * basetime
	c.RetryWait = basetime

	mux.HandleFunc("/ep", func(w http.ResponseWriter, r *http.Request) {
		if ctx1.Err() == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
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
}
