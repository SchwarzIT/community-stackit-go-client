package wait

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type WaitFn func() (res interface{}, done bool, err error)

type Wait struct {
	fn       WaitFn
	throttle time.Duration
	timeout  time.Duration
}

// New creates a new Wait instance
func New(f WaitFn) *Wait {
	return &Wait{
		fn:       f,
		throttle: 5 * time.Second,
		timeout:  30 * time.Minute,
	}
}

// SetThrottle sets the duration between func triggerring
func (w *Wait) SetThrottle(d time.Duration) *Wait {
	w.throttle = d
	return w
}

// SetTimeout sets the duration for wait timeout
func (w *Wait) SetTimeout(d time.Duration) *Wait {
	w.timeout = d
	return w
}

// Do starts the wait until there's an error or wait is done
func (w *Wait) Run() (res interface{}, err error) {
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
