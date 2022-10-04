package retry

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Retry struct {
	Timeout        time.Duration               // max duration
	MaxRetries     *int                        // max retries, when nil, there's no retry limit
	Throttle       time.Duration               // wait duration between calls
	IsRetryableFns []func(err error) bool      // functions to determine if error is retriable
	UntilFns       []func(*http.Response) bool // functions to the determine if the given response is the expected response
}

type Config interface {
	String() string
	Value() interface{}
}

// New returns a new instance of Retry with default values
func New() *Retry {
	c := &Retry{
		Timeout:        CONFIG_TIMEOUT_DEFAULT,
		MaxRetries:     nil,
		Throttle:       CONFIG_THROTTLE_DEFAULT,
		IsRetryableFns: []func(err error) bool{IsRetryableDefault},
	}
	return c
}

func (c *Retry) withConfig(cfgs ...Config) *Retry {
	for _, cfg := range cfgs {
		switch cfg.String() {
		case CONFIG_TIMEOUT:
			c.Timeout = cfg.Value().(time.Duration)
		case CONFIG_THROTTLE:
			c.Throttle = cfg.Value().(time.Duration)
		case CONFIG_MAX_RETRIES:
			c.MaxRetries = cfg.Value().(*int)
		case CONFIG_IS_RETRYABLE:
			c.IsRetryableFns = cfg.Value().([]func(err error) bool)
		case CONFIG_UNTIL:
			c.UntilFns = cfg.Value().([]func(*http.Response) bool)
		}
	}
	return c
}

// Do runs a given do function with a given request -> do(req)
func (r *Retry) Do(req *http.Request, do func(*http.Request) (*http.Response, error)) (*http.Response, error) {
	var lastErr error
	var res *http.Response
	var retry bool

	overall, cancel := context.WithTimeout(req.Context(), r.Timeout)
	defer cancel()

	// set overall ctx in request
	newReq := req.WithContext(overall)

	// clone max retries
	maxRetries := -1
	if r.MaxRetries != nil {
		maxRetries = *r.MaxRetries
	}

	for {
		res, maxRetries, retry, lastErr = r.doIteration(newReq, do, maxRetries)
		if lastErr == nil && r.isFulfilled(res) {
			return res, nil
		}

		if !retry {
			return res, lastErr
		}

		// context timeout was chosen in order to support throttle = 0
		tick, cancelTick := context.WithTimeout(context.Background(), r.Throttle)
		defer cancelTick()

		select {
		case <-tick.Done():
			// continue
		case <-overall.Done():
			return nil, errors.Wrap(lastErr, "retry context timed out")
		}
	}
}

// doIteration runs the do function with the given request
func (r *Retry) doIteration(req *http.Request, do func(*http.Request) (*http.Response, error), retries int) (res *http.Response, retriesLeft int, retry bool, err error) {
	retriesLeft = retries
	retry = true

	res, err = do(req)
	if err == nil {
		return
	}

	// check if error is retryable
	for _, f := range r.IsRetryableFns {
		if !f(err) {
			retry = false
			return
		}
	}

	if retries != -1 {
		if retries == 0 {
			err = errors.Wrap(err, "reached max retries")
			retry = false
			return
		}
		retriesLeft = retries - 1
	}

	return
}

// isFulfilled check if Until functions are all fulfilled
func (r *Retry) isFulfilled(res *http.Response) bool {
	for _, f := range r.UntilFns {
		if !f(res) {
			return false
		}
	}
	return true
}
