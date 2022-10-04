package retry

import (
	"errors"
	"net/http"
	"reflect"
	"runtime"
	"testing"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestRetry_SetIsRetryable(t *testing.T) {
	r := New()
	r.IsRetryableFns = []func(err error) bool{IsRetryableNoOp}

	type args struct {
		f []func(err error) bool
	}
	tests := []struct {
		name string
		args args
		want *Retry
	}{
		{"ok", args{[]func(err error) bool{IsRetryableNoOp}}, r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got := c.SetIsRetryable(tt.args.f...)
			if len(got.IsRetryableFns) != len(r.IsRetryableFns) {
				t.Error("wrong lengths")
				return
			}
			for k, v := range got.IsRetryableFns {
				if GetFunctionName(r.IsRetryableFns[k]) != GetFunctionName(v) {
					t.Errorf("%s != %s", GetFunctionName(r.IsRetryableFns[k]), GetFunctionName(v))
					return
				}
			}
		})
	}
}

func TestIsRetryableNoOp(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"always true", args{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRetryableNoOp(tt.args.err); got != tt.want {
				t.Errorf("IsRetryableNoOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRetryableDefault(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"bad request - no retry", args{errors.New(http.StatusText(http.StatusBadRequest))}, false},
		{"bad request - no host - retry", args{errors.New(http.StatusText(http.StatusBadRequest) + ERR_NO_SUCH_HOST)}, true},
		{"bad request - unauthorized", args{errors.New(http.StatusText(http.StatusUnauthorized))}, false},
		{"other error", args{errors.New(http.StatusText(http.StatusConflict))}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRetryableDefault(tt.args.err); got != tt.want {
				t.Errorf("IsRetryableDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
