package clients

import (
	"reflect"
	"testing"
)

func TestNewRetryConfig(t *testing.T) {
	got := NewRetryConfig()
	if got == nil {
		t.Error("NewRetryConfig returned nil")
	}
	want := RetryConfig{
		MaxRetries:       DefaultRetryMaxRetries,
		WaitBetweenCalls: DefaultRetryWaitBetweenCalls,
		Timeout:          DefaultRetryTimeout,
	}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("%+v != %+v", *got, want)
	}
}
