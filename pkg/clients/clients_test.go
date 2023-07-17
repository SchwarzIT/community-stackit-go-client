package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewRetryConfig(t *testing.T) {
	got := NewRetryConfig()
	if got == nil {
		t.Error("NewRetryConfig returned nil")
	}
	want := RetryConfig{
		MaxRetries:       DefaultRetryMaxRetries,
		WaitBetweenCalls: DefaultRetryWaitBetweenCalls,
		RetryTimeout:     DefaultRetryTimeout,
		ClientTimeout:    DefaultClientTimeout,
	}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("%+v != %+v", *got, want)
	}
}

func Test_do(t *testing.T) {
	type args struct {
		cfg            *RetryConfig
		serverStatus   int
		serverResponse string
	}
	tests := []struct {
		name     string
		args     args
		wantResp *http.Response
		wantErr  bool
		errMsg   string
	}{
		{"all ok", args{
			cfg:            &RetryConfig{0, time.Microsecond, time.Second, DefaultClientTimeout, false},
			serverStatus:   http.StatusOK,
			serverResponse: `{"status":"ok", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusOK}, false, ""},
		{"all ok nil client", args{
			cfg:            &RetryConfig{0, time.Microsecond, time.Second, DefaultClientTimeout, false},
			serverStatus:   http.StatusOK,
			serverResponse: `{"status":"ok", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusOK}, false, ""},
		{"fail 1", args{
			cfg:            &RetryConfig{1, time.Microsecond, time.Second, DefaultClientTimeout, false},
			serverStatus:   http.StatusInternalServerError,
			serverResponse: `{"status":"error 1", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusInternalServerError}, false, ""},
		{"fail 2 - timeout error", args{
			cfg:            &RetryConfig{3, time.Microsecond, time.Second, DefaultClientTimeout, false},
			serverStatus:   http.StatusOK,
			serverResponse: `{"status":"ok", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusOK}, true, "no such host"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.args.serverStatus)
				fmt.Fprintln(w, fmt.Sprintf(tt.args.serverResponse, tt.name))
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
			c := &http.Client{}
			if tt.name == "fail 2 - timeout error" && server != nil {
				t.Log("closing server")
				server.Close()
			}
			if tt.name == "all ok nil client" {
				c = nil
			}
			gotResp, err := do(c, req, tt.args.cfg)
			if server != nil {
				server.Close()
			}
			if (err != nil) != tt.wantErr && (err != nil && !strings.Contains(err.Error(), tt.errMsg)) {
				body := []byte{}
				if gotResp != nil {
					body, err = ioutil.ReadAll(gotResp.Body)
					if err != nil {
						t.Error(err)
						return
					}
				}
				t.Errorf("do() error = %v, wantErr %v, got: %s", err, tt.wantErr, string(body))
				return
			}
			if gotResp != nil && tt.wantResp.StatusCode != gotResp.StatusCode {
				t.Errorf("do() status code = %v, want %v", tt.wantResp.StatusCode, gotResp.StatusCode)
			}
		})
	}
}
