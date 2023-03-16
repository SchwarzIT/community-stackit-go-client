package validate_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
)

func TestUUID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"5dae0612-f5b1-4615-b7ca-b18796aa7e78"}, false},
		{"not ok", args{"bad-uuid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.UUID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectID(t *testing.T) {
	type args struct {
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"5dae0612-f5b1-4615-b7ca-b18796aa7e78"}, false},
		{"not ok", args{"bad-uuid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.ProjectID(tt.args.projectID); (err != nil) != tt.wantErr {
				t.Errorf("ProjectID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok", args{"project name!"}, true},
		{"ok [1]", args{"project name"}, false},
		{"ok [2]", args{"project-name"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.ProjectName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ProjectName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBillingRef(t *testing.T) {
	type args struct {
		billingRef string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok", args{"invalid!"}, true},
		{"ok", args{"T-123456B"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.BillingRef(tt.args.billingRef); (err != nil) != tt.wantErr {
				t.Errorf("BillingRef() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultResponseErrorHandler(t *testing.T) {
	r := io.NopCloser(strings.NewReader("ABC"))
	resp := &http.Response{StatusCode: 400, Body: r, ContentLength: 3, Request: &http.Request{URL: &url.URL{}}}
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{&http.Response{StatusCode: 202}}, false},
		{"not ok", args{resp}, true},
		{"not ok 2", args{resp}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.DefaultResponseErrorHandler(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("DefaultResponseErrorHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSemVer(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok 1", args{"1.2.3"}, false},
		{"ok 2", args{"1.2"}, false},
		{"not ok 1", args{"ab1.2.3"}, true},
		{"not ok 2", args{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.SemVer(tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("SemVer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRFC3339(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"2020-09-04T00:00:00Z"}, false},
		{"not ok 1", args{"2020-09-04 00:00:00"}, true},
		{"not ok 2", args{"2020/09/04 00:00"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.RFC3339(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("RFC3339() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestISO8601(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"2020-09-04T00:00:00.605Z"}, false},
		{"not ok 1", args{"2020-09-04 00:00:00"}, true},
		{"not ok 2", args{"2020/09/04 00:00"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.ISO8601(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("ISO8601() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"1s"}, false},
		{"ok 2", args{"1m"}, false},
		{"ok 3", args{"60s"}, false},
		{"not ok 1", args{"abcd"}, true},
		{"not ok 2", args{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := validate.Duration(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Duration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponse(t *testing.T) {
	type args struct {
		resp            interface{}
		requestError    error
		checkNullFields []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"error from request", args{requestError: errors.New("a request error")}, "a request error"},
		{"nil resp", args{requestError: nil, resp: nil}, "response interface is nil"},
		{"no Error", args{requestError: nil, resp: struct{}{}}, "No such field: Error in obj"},
		{"not struct", args{requestError: nil, resp: 1}, "Cannot use GetField on a non-struct interface"},
		{"nil Error", args{requestError: nil, resp: struct{ Error error }{}}, ""},
		{"defined Error", args{requestError: nil, resp: struct{ Error error }{Error: errors.New("an error")}}, "an error"},
		{"nil Error, nil JSON200", args{requestError: nil, resp: struct {
			Error   error
			JSON200 interface{}
		}{JSON200: nil}, checkNullFields: []string{"JSON200"}}, "field JSON200 in response is nil"},
		{"nil Error, notfound JSON200", args{requestError: nil, resp: struct {
			Error error
		}{}, checkNullFields: []string{"JSON200"}}, "No such field: JSON200 in obj"},
		{"nil Error, JSON200.ABC is nil", args{requestError: nil, resp: struct {
			Error   error
			JSON200 struct{ ABC *string }
		}{JSON200: struct{ ABC *string }{ABC: nil}}, checkNullFields: []string{"JSON200.ABC"}}, "field JSON200.ABC in response is nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.Response(tt.args.resp, tt.args.requestError, tt.args.checkNullFields...); (err != nil) && (err.Error() != tt.want) {
				t.Errorf("Response() error = %v, want %s", err, tt.want)
			}
		})
	}
}

type sample struct{}

func (sample) StatusCode() int {
	return http.StatusAccepted
}

func TestStatusEquals(t *testing.T) {

	var a *sample = &sample{}
	var b *sample = nil

	if validate.StatusEquals(b, http.StatusAccepted) {
		t.Error("expected false for b, got true")
	}

	if !validate.StatusEquals(a, http.StatusAccepted) {
		t.Error("expected true for a, got false")
	}
	if validate.StatusEquals(a, http.StatusBadGateway) {
		t.Error("expected false for a, got true")
	}
}
