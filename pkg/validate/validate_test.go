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
	abc := "abc"
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
		{"nil Error, nil fields", args{requestError: nil, resp: struct {
			Error   error
			JSON200 struct{ ABC *string }
		}{JSON200: struct{ ABC *string }{ABC: &abc}}, checkNullFields: []string{"JSON200.ABC"}}, ""},
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

func TestErrorIsOneOf(t *testing.T) {
	type args struct {
		err  error
		msgs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil err", args{nil, nil}, false},
		{"no match 1", args{errors.New("abcd"), []string{}}, false},
		{"no match 2", args{errors.New("abcd"), []string{"efd"}}, false},
		{"no match 3", args{errors.New("abcd"), []string{"efd", "hij"}}, false},
		{"match", args{errors.New("abcd"), []string{"cd"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate.ErrorIsOneOf(tt.args.err, tt.args.msgs...); got != tt.want {
				t.Errorf("ErrorHasAnySubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetworkName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"no ok", args{"Invalid Name!"}, true},
		{"ok [1]", args{"My-Example_String.123"}, false},
		{"ok [2]", args{"Hello_World"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.NetworkName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("NetworkName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNameServer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok [1]", args{"256.100.50.25"}, true},
		{"not ok [2]", args{"GHT:::1200:::0000"}, true},
		{"ok [1]", args{"192.168.1.1"}, false},
		{"ok [2]", args{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, false},
		{"ok [2]", args{"\"1.1.1.1\""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.NameServer(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("NameServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrefixLength(t *testing.T) {
	type args struct {
		prefix int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok [1]", args{21}, true},
		{"not ok [2]", args{30}, true},
		{"ok [1]", args{22}, false},
		{"ok [2]", args{29}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.PrefixLengthV4(tt.args.prefix); (err != nil) != tt.wantErr {
				t.Errorf("NameServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrefix(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok", args{"999.999.999.999/32"}, true},
		{"ok [1]", args{"192.168.1.1/24"}, false},
		{"ok [2]", args{"fe80::1ff:fe23:4567:890a/64"}, false},
		{"ok [1]", args{"\"192.168.1.1/24\""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.Prefix(tt.args.prefix); (err != nil) != tt.wantErr {
				t.Errorf("Prefix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPublicIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok", args{"999.999.999.999"}, true},
		{"ok [1]", args{"192.168.1.1"}, false},
		{"ok [2]", args{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.PublicIP(tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("PublicIP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetworkID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not ok", args{"14b0cd73-3342-44a3-a10e-q1bc20af4497"}, true},
		{"not ok", args{"14b0cd73-3342-44a3-a10e-q1bc20af44979"}, true},
		{"ok [1]", args{"14b0cd73-3342-44a3-a10e-f1bc20af4497"}, false},
		{"ok [2]", args{"14b0cd52-6677-44a3-a10e-f1bc21af4445"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.NetworkID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("NetworkID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
