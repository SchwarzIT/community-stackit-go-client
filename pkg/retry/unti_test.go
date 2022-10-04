package retry

import (
	"net/http"
	"testing"
)

func TestRetry_Until(t *testing.T) {
	r := New()
	r.UntilFns = []func(*http.Response) bool{noOp}

	type args struct {
		f []func(*http.Response) bool
	}
	tests := []struct {
		name string
		args args
		want *Retry
	}{
		{"ok", args{[]func(*http.Response) bool{noOp}}, r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got := c.Until(tt.args.f...)
			if len(got.UntilFns) != len(r.UntilFns) {
				t.Error("wrong lengths")
				return
			}
			for k, v := range got.UntilFns {
				if GetFunctionName(r.UntilFns[k]) != GetFunctionName(v) {
					t.Errorf("%s != %s", GetFunctionName(r.UntilFns[k]), GetFunctionName(v))
					return
				}
			}
		})
	}
}

func noOp(_ *http.Response) bool {
	return true
}
