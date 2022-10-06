package wait

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type WaitFn func() (done bool, err error)

type Wait struct {
	fn       WaitFn
	throttle time.Duration
}

// New creates a new Wait instance
func New(f WaitFn) *Wait {
	return &Wait{
		fn:       f,
		throttle: 5 * time.Second,
	}
}

// SetThrottle sets the duration between func triggerring
func (w *Wait) SetThrottle(d time.Duration) *Wait {
	w.throttle = d
	return w
}

// Do starts the wait until there's an error or wait is done
func (w *Wait) Run(timeout time.Duration) (done bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		done, err = w.fn()
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
			return false, errors.New("wait.Do() timed out")
		}
	}
}
