package wait

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type WaitFn func() (res interface{}, done bool, err error)

type Handler struct {
	fn           WaitFn
	throttle     time.Duration
	timeout      time.Duration
	initialDelay time.Duration
}

// New creates a new Handler instance with sensible defaults
func New(f WaitFn) *Handler {
	return &Handler{
		fn:       f,
		throttle: 5 * time.Second,
		timeout:  30 * time.Minute,
	}
}

// SetThrottle sets the duration between each poll. Returns *Handler for chaining.
func (w *Handler) SetThrottle(d time.Duration) (*Handler, error) {
	if d == 0 {
		return nil, errors.New("throttle duration cannot be 0")
	}
	w.throttle = d
	return w, nil
}

// SetTimeout sets the maximum duration before Wait times out. Returns *Handler for chaining.
func (w *Handler) SetTimeout(d time.Duration) *Handler {
	w.timeout = d
	return w
}

// SetInitialDelay sets an optional delay before the first poll is executed.
// This is useful when an operation is known to need time before its status changes.
func (w *Handler) SetInitialDelay(d time.Duration) *Handler {
	w.initialDelay = d
	return w
}

// Wait starts polling until the WaitFn returns done, an error occurs, or the timeout is reached.
func (w *Handler) Wait() (interface{}, error) {
	return w.WaitWithContext(context.Background())
}

// WaitWithContext starts polling until the WaitFn returns done, an error occurs,
// the provided context is cancelled, or the timeout is reached.
func (w *Handler) WaitWithContext(ctx context.Context) (res interface{}, err error) {
	var done bool

	ctx, cancel := context.WithTimeout(ctx, w.timeout)
	defer cancel()

	if w.initialDelay > 0 {
		select {
		case <-time.After(w.initialDelay):
			// proceed to polling
		case <-ctx.Done():
			return nil, errors.New("wait timed out during initial delay")
		}
	}

	ticker := time.NewTicker(w.throttle)
	defer ticker.Stop()

	for {
		res, done, err = w.fn()
		if err != nil {
			return res, errors.Wrap(err, "wait function returned an error")
		}
		if done {
			return res, nil
		}

		select {
		case <-ticker.C:
			// continue to next poll
		case <-ctx.Done():
			return res, errors.New("wait timed out")
		}
	}
}
