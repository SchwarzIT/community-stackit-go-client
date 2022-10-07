package wait

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {
	simple := func() (res interface{}, done bool, err error) { return nil, true, nil }
	type args struct {
		f WaitFn
	}
	tests := []struct {
		name string
		args args
		want *Wait
	}{
		{"ok", args{simple}, &Wait{fn: simple, throttle: 5 * time.Second}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.f); !reflect.DeepEqual(got.throttle, tt.want.throttle) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWait_SetThrottle(t *testing.T) {
	simple := func() (res interface{}, done bool, err error) { return nil, true, nil }
	f := &Wait{
		fn:       simple,
		throttle: 10 * time.Second,
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want *Wait
	}{
		{"ok", args{10 * time.Second}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(simple)
			if got := w.SetThrottle(tt.args.d); !reflect.DeepEqual(got.throttle, tt.want.throttle) {
				t.Errorf("Wait.SetThrottle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWait_Run(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	type fields struct {
		fn       WaitFn
		throttle time.Duration
		timeout  time.Duration
	}
	tests := []struct {
		name     string
		fields   fields
		wantDone bool
		wantErr  bool
	}{
		{"ok", fields{throttle: 1 * time.Second, timeout: 1 * time.Hour, fn: func() (res interface{}, done bool, err error) {
			return nil, true, nil
		}}, true, false},

		{"ok 2", fields{throttle: 1 * time.Second, timeout: 1 * time.Hour, fn: func() (res interface{}, done bool, err error) {
			if ctx.Err() == nil {
				return nil, false, nil
			}
			return nil, true, nil
		}}, true, false},

		{"err", fields{throttle: 1 * time.Second, timeout: 1 * time.Hour, fn: func() (res interface{}, done bool, err error) {
			return nil, true, errors.New("something happened")
		}}, true, true},

		{"timeout", fields{throttle: 1 * time.Second, timeout: 1 * time.Second, fn: func() (res interface{}, done bool, err error) {
			return nil, false, nil
		}}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wait{
				fn:       tt.fields.fn,
				throttle: tt.fields.throttle,
				timeout:  tt.fields.timeout,
			}
			_, err := w.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Wait.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWait_SetTimeout(t *testing.T) {
	f := &Wait{
		throttle: 5 * time.Second,
		timeout:  5 * time.Hour,
	}

	type fields struct {
		throttle time.Duration
		timeout  time.Duration
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Wait
	}{
		{"ok", fields{timeout: 1 * time.Hour, throttle: 5 * time.Second}, args{d: 5 * time.Hour}, f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wait{
				throttle: tt.fields.throttle,
				timeout:  tt.fields.timeout,
			}
			if got := w.SetTimeout(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wait.SetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
