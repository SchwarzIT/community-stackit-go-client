package retry_test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/retry"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *retry.Retry
	}{
		{"ok", retry.New()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := retry.New()
			if got.MaxRetries != tt.want.MaxRetries ||
				got.Throttle != tt.want.Throttle ||
				got.Timeout != tt.want.Timeout ||
				len(got.IsRetryableFns) != len(tt.want.IsRetryableFns) {
				t.Error("one or more values don't match")
			}
		})
	}
}

const (
	response_before_2s = `{"status":false}`
	response_after_2s  = `{"status":true}`
)

func TestClient_DoWithRetryThrottle(t *testing.T) {
	c, mux, teardown, err := client.MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	mux.HandleFunc("/2s", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if ctx.Err() != nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, response_after_2s)
		} else {
			w.WriteHeader(http.StatusLocked)
			fmt.Fprint(w, response_before_2s)
		}
	})

	c = c.WithRetry(retry.New().SetThrottle(1 * time.Second))

	req, _ := c.Request(context.Background(), http.MethodGet, "/2s", nil)

	var got struct {
		Status bool `json:"status"`
	}

	if _, err := c.Do(req, &got); err != nil {
		t.Errorf("do request: %v", err)
	}

	if !got.Status {
		t.Errorf("received status = %v", got.Status)
	}
}

func TestClient_DoWithRetryNonRetryableError(t *testing.T) {
	c, mux, teardown, err := client.MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	mux.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	})

	c = c.WithRetry(retry.New())
	req, _ := c.Request(context.Background(), http.MethodGet, "/err", nil)
	if _, err := c.Do(req, nil); err == nil {
		t.Error("expected do request to return error but got nil instead")
	}
}

func TestClient_DoWithRetryMaxRetries(t *testing.T) {
	c, mux, teardown, err := client.MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	mux.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusLocked)
	})

	c = c.WithRetry(retry.New().SetThrottle(1 * time.Second).SetMaxRetries(2))
	req, _ := c.Request(context.Background(), http.MethodGet, "/err", nil)
	if _, err := c.Do(req, nil); !strings.Contains(err.Error(), "reached max retries") {
		t.Errorf("expected do request to return max retries error but got '%v' instead", err)
	}
}

func TestClient_DoWithRetryTimeout(t *testing.T) {
	c, mux, teardown, err := client.MockServer()
	defer teardown()
	if err != nil {
		t.Errorf("error from mock.AuthServer: %s", err.Error())
	}

	mux.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusLocked)
	})

	c = c.WithRetry(retry.New().SetTimeout(5 * time.Second))
	req, _ := c.Request(context.Background(), http.MethodGet, "/err", nil)
	if _, err := c.Do(req, nil); !strings.Contains(err.Error(), "retry context timed out") && !strings.Contains(err.Error(), http.StatusText(http.StatusLocked)) {
		t.Errorf("expected do request to return retry context timed out error with locked error status but got '%v' instead", err)
	}

	c = c.WithRetry(retry.New().SetTimeout(0 * time.Second))
	req, _ = c.Request(context.Background(), http.MethodGet, "/err", nil)
	if _, err := c.Do(req, nil); !strings.Contains(err.Error(), "retry context timed out") {
		t.Errorf("expected do request to return retry context timed out error but got '%v' instead", err)
	}
}
