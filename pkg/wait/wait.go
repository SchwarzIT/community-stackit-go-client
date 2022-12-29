package wait

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type WaitFn func() (res interface{}, done bool, err error)

type Handler struct {
	fn       WaitFn
	throttle time.Duration
	timeout  time.Duration
}

// New creates a new Wait instance
func New(f WaitFn) *Handler {
	return &Handler{
		fn:       f,
		throttle: 5 * time.Second,
		timeout:  30 * time.Minute,
	}
}

// SetThrottle sets the duration between func triggering
func (w *Handler) SetThrottle(d time.Duration) (*Handler, error) {
	if d == 0 {
		return w, errors.New("Throttle can't be 0")
	}
	w.throttle = d
	return w, nil
}

// SetTimeout sets the duration for wait timeout
func (w *Handler) SetTimeout(d time.Duration) *Handler {
	w.timeout = d
	return w
}

// Wait starts the wait until there's an error or wait is done
func (w *Handler) Wait() (res interface{}, err error) {
	var done bool
	ctx, cancel := context.WithTimeout(context.Background(), w.timeout)
	defer cancel()
	for {
		res, done, err = w.fn()
		if err != nil || done {
			return
		}

		// context timeout was chosen in order to support throttle = 0
		tick, cancelTick := context.WithTimeout(context.Background(), w.throttle)
		defer cancelTick()

		select {
		case <-tick.Done():
			// continue
		case <-ctx.Done():
			return res, errors.New("Wait() has timed out")
		}
	}
}

// WaitWithContext starts the wait until there's an error or wait is done
func (w Handler) WaitWithContext(ctx context.Context) (res interface{}, err error) {
	var done bool

	ctx, cancel := context.WithTimeout(ctx, w.timeout)
	defer cancel()

	ticker := time.NewTicker(w.throttle)
	defer ticker.Stop()

	for {
		res, done, err = w.fn()
		if err != nil {
			return res, errors.Wrap(err, "defined wait function returned an error")
		}
		if done {
			return res, nil
		}

		select {
		case <-ticker.C:
			// continue
		case <-ctx.Done():
			return res, errors.New("Wait() has timed out")
		}
	}
}
