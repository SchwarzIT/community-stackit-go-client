package retry

import (
	"reflect"
	"testing"
	"time"
)

func TestRetry_SetThrottle(t *testing.T) {
	r := New()
	r.throttle = time.Minute
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want *Retry
	}{
		{"ok", args{time.Minute}, r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			if got := c.SetThrottle(tt.args.d); !reflect.DeepEqual(got.throttle, tt.want.throttle) {
				t.Errorf("Retry.SetThrottle() = %v, want %v", got.throttle, tt.want.throttle)
			}
		})
	}
}
