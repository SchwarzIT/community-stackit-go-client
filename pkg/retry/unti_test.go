package retry

import (
	"net/http"
	"testing"
)

func TestRetry_Wait(t *testing.T) {
	r := New()

	r.untilFns = []UntilFn{noOp}

	type args struct {
		f []UntilFn
	}
	tests := []struct {
		name string
		args args
		want *Retry
	}{
		{"ok", args{[]UntilFn{noOp}}, r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got := c.SetUntil(tt.args.f...)
			if len(got.untilFns) != len(r.untilFns) {
				t.Error("wrong lengths")
				return
			}
			for k, v := range got.untilFns {
				if GetFunctionName(r.untilFns[k]) != GetFunctionName(v) {
					t.Errorf("%s != %s", GetFunctionName(r.untilFns[k]), GetFunctionName(v))
					return
				}
			}
		})
	}
}

func noOp(_ *http.Response) (bool, error) {
	return true, nil
}
