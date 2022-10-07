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

// SetThrottle sets the duration between func triggerring
func (w *Handler) SetThrottle(d time.Duration) *Handler {
	w.throttle = d
	return w
}

// SetTimeout sets the duration for wait timeout
func (w *Handler) SetTimeout(d time.Duration) *Handler {
	w.timeout = d
	return w
}

// Do starts the wait until there's an error or wait is done
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
			return res, errors.New("wait.Do() timed out")
		}
	}
}
