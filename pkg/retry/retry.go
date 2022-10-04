package retry

import (
	"time"
)

type Retry struct {
	Timeout        time.Duration          // max duration
	MaxRetries     *int                   // max retries, when nil, there's no retry limit
	Throttle       time.Duration          // wait duration between calls
	IsRetryableFns []func(err error) bool // functions to determine if error is retriable
}

type Config interface {
	String() string
	Value() interface{}
}

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
		}
	}
	return c
}
