package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

func TestNewAuth(t *testing.T) {
	cfg := &AuthConfig{ClientID: consts.SCHWARZ_ORGANIZATION_ID, ClientSecret: "secret"}
	type args struct {
		ctx context.Context
		cfg *AuthConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *AuthClient
		wantErr bool
	}{
		{"no client ID", args{context.Background(), &AuthConfig{}}, &AuthClient{}, true},
		{"no client secret", args{context.Background(), &AuthConfig{ClientID: consts.SCHWARZ_ORGANIZATION_ID}}, &AuthClient{}, true},
		{"all ok", args{context.Background(), cfg}, &AuthClient{Config: cfg}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuth(tt.args.ctx, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got.Config, tt.want.Config) {
				t.Errorf("NewAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthClient_Request(t *testing.T) {
	c, err := NewAuth(context.Background(), &AuthConfig{ClientID: consts.SCHWARZ_ORGANIZATION_ID, ClientSecret: "secret"})
	if err != nil {
		t.Error(err)
	}

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
		{"bad context", args{nil, "something", "my-path", ""}, &http.Request{}, true},
		{"bad method", args{context.Background(), "something", "my-path", ""}, &http.Request{}, true},
		{"all ok", args{context.Background(), http.MethodGet, "my-path", ""}, &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/my-path"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Request(tt.args.ctx, tt.args.method, tt.args.path, tt.args.body)
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

func TestAuthClient_Do(t *testing.T) {
	type Test struct {
		Payload string `json:"payload,omitempty"`
	}

	c, mux, teardown, err := MockAuthServer()
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

	req, err := c.Request(context.Background(), http.MethodPost, "/echo", string(body))
	if err != nil {
		t.Errorf("new request: %v", err)
		return
	}

	var got Test
	_, err = c.Do(req, &got)
	if err != nil {
		t.Errorf("do request: %v", err)
	}

	want := p
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}

	req.URL.Path = "/blah"
	if _, err = c.Do(req, &got); err == nil {
		t.Error("expected error over path")
	}

	ctx2, cancel := context.WithTimeout(context.TODO(), 0)
	defer cancel()
	req = req.WithContext(ctx2)
	if _, err = c.Do(req, &got); err == nil {
		t.Error("expected error over context timeout")
	}

}

func TestAuthClient_GetToken(t *testing.T) {
	c, mux, teardown, err := MockAuthServer()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}
	defer teardown()

	mux.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		r.ParseForm()
		cid := r.Form.Get("client_id")
		if cid != "id" {
			t.Error("bad client id provided")
		}
		csc := r.Form.Get("client_secret")
		if csc != "secret" {
			t.Error("bad client secret provided")
		}
		cgt := r.Form.Get("grant_type")
		if cgt != "client_credentials" {
			t.Error("bad grant type provided")
		}

		b, err := json.Marshal(AuthAPIClientCredentialFlowRes{
			AccessToken: "my-access-token",
		})
		if err != nil {
			log.Fatalf("json response marshal: %v", err)
		}
		fmt.Fprint(w, string(b))
	})

	got, err := c.GetToken(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := AuthAPIClientCredentialFlowRes{
		AccessToken: "my-access-token",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}

	ctx2, cancel := context.WithTimeout(context.TODO(), 0)
	defer cancel()
	if _, err = c.GetToken(ctx2); err == nil {
		t.Error("expected error over context timeout")
	}
}
