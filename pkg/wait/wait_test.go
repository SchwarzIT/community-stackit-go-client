package wait

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
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
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantDone bool
		wantErr  bool
	}{
		{"ok", fields{throttle: 1 * time.Second, fn: func() (res interface{}, done bool, err error) {
			return nil, true, nil
		}}, args{1 * time.Hour}, true, false},

		{"ok 2", fields{throttle: 1 * time.Second, fn: func() (res interface{}, done bool, err error) {
			if ctx.Err() == nil {
				return nil, false, nil
			}
			return nil, true, nil
		}}, args{1 * time.Hour}, true, false},

		{"err", fields{throttle: 1 * time.Second, fn: func() (res interface{}, done bool, err error) {
			return nil, true, errors.New("something happened")
		}}, args{1 * time.Hour}, true, true},

		{"timeout", fields{throttle: 1 * time.Second, fn: func() (res interface{}, done bool, err error) {
			return nil, false, nil
		}}, args{1 * time.Second}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wait{
				fn:       tt.fields.fn,
				throttle: tt.fields.throttle,
			}
			_, gotDone, err := w.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Wait.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDone != tt.wantDone {
				t.Errorf("Wait.Run() = %v, want %v", gotDone, tt.wantDone)
			}
		})
	}
}
