package retry

import (
	"reflect"
	"testing"
	"time"
)

func TestRetry_SetTimeout(t *testing.T) {
	r := New()
	r.timeout = 1 * time.Minute
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want *Retry
	}{
		{"all ok", args{time.Minute}, r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			if got := c.SetTimeout(tt.args.d); !reflect.DeepEqual(got.timeout, tt.want.timeout) {
				t.Errorf("Retry.SetTimeout() = %v, want %v", got.timeout, tt.want.timeout)
			}
		})
	}
}
