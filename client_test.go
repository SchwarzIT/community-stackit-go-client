package stackit

import (
	"context"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

func TestNew(t *testing.T) {

	eml := os.Getenv(email)
	tok := os.Getenv(token)
	env := os.Getenv(apienv)

	if tok == "" && eml == "" {
		os.Setenv(email, "abc")
		os.Setenv(token, "efg")
		os.Setenv(apienv, "hij")
		tc := NewClient(context.Background())
		if tc.GetConfig().ServiceAccountEmail != "abc" || tc.GetConfig().ServiceAccountToken != "efg" || tc.GetConfig().Environment != "hij" {
			t.Errorf("NewClient config doesn't match")
		}
	} else {
		tc := NewClient(context.Background())
		if tc.GetConfig().ServiceAccountEmail != eml || tc.GetConfig().ServiceAccountToken != tok {
			t.Errorf("NewClient config doesn't match env")
		}
	}
	os.Setenv(email, "")
	os.Setenv(token, "")
	os.Setenv(apienv, "")

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
			got, err := NewClientWithConfig(tt.args.ctx, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (!reflect.DeepEqual(got.config.ServiceAccountEmail, tt.want.config.ServiceAccountEmail) || !reflect.DeepEqual(got.config.ServiceAccountToken, tt.want.config.ServiceAccountToken)) {
				t.Errorf("NewAuth() = %v, want %v", got, tt.want)
			}
		})
	}
	os.Setenv(email, eml)
	os.Setenv(token, tok)
	os.Setenv(apienv, env)
}

func TestClient_DoWithRetry(t *testing.T) {
	c, mux, teardown, err := MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	if c.GetEnvironment() != common.ENV_PROD {
		t.Error("default environment isn't set to prod")
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.GetConfig().BaseUrl.String()+"/ep", nil)
	if err != nil {
		t.Error(err)
	}

	if _, err := c.Do(req); err != nil {
		t.Error(err)
	}

	ctx3, td3 := context.WithTimeout(ctx, 6*basetime)
	defer td3()

	c.GetHTTPClient().Timeout = basetime * 1
	mux.HandleFunc("/ep2", func(w http.ResponseWriter, r *http.Request) {

		if ctx3.Err() == nil {
			time.Sleep(1 * basetime)
		}
	})

	req, err = http.NewRequestWithContext(ctx, http.MethodGet, c.config.BaseUrl.String()+"/ep2", nil)
	if err != nil {
		t.Error(err)
	}

	if _, err := c.Do(req); err != nil {
		t.Error(err)
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

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "23423!$!err2", nil)
	if err != nil {
		t.Error(err)
	}
	if _, err := c.Do(req); err == nil {
		t.Error("expected do request to return error")
	}
}

func TestClient_GetEnvironment(t *testing.T) {
	type fields struct {
		ctx         context.Context
		client      *http.Client
		config      Config
		RetryTimout time.Duration
		RetryWait   time.Duration
		services    services
	}
	tests := []struct {
		name   string
		fields fields
		want   common.Environment
	}{
		{"ok 1", fields{config: Config{Environment: ""}}, common.ENV_PROD},
		{"ok 2", fields{config: Config{Environment: "dev"}}, common.ENV_DEV},
		{"ok 3", fields{config: Config{Environment: "qa"}}, common.ENV_QA},
		{"ok 4", fields{config: Config{Environment: "prod"}}, common.ENV_PROD},
		{"ok 5", fields{config: Config{Environment: "test"}}, common.ENV_PROD},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				ctx:         tt.fields.ctx,
				client:      tt.fields.client,
				config:      tt.fields.config,
				RetryTimout: tt.fields.RetryTimout,
				RetryWait:   tt.fields.RetryWait,
				services:    tt.fields.services,
			}
			if got := c.GetEnvironment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}
