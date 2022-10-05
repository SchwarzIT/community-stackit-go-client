package retry

import (
	"reflect"
	"testing"
)

func TestRetry_SetMaxRetries(t *testing.T) {
	r := New()
	five := 5
	r.maxRetries = &five
	type args struct {
		r int
	}
	tests := []struct {
		name string
		args args
		want *Retry
	}{
		{"ok", args{5}, r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			if got := c.SetMaxRetries(tt.args.r); !reflect.DeepEqual(*got.maxRetries, *tt.want.maxRetries) {
				t.Errorf("Retry.SetMaxRetries() = %v, want %v", *got.maxRetries, *tt.want.maxRetries)
			}
		})
	}
}
