package retry

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Retry
	}{
		{"ok", New()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			if got.MaxRetries != tt.want.MaxRetries ||
				got.Throttle != tt.want.Throttle ||
				got.Timeout != tt.want.Timeout ||
				len(got.IsRetryableFns) != len(tt.want.IsRetryableFns) {
				t.Error("one or more values don't match")
			}
		})
	}
}
