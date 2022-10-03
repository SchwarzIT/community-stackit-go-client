package retry

import (
	"time"
)

type Retry struct {
	Timeout        time.Duration          // max duration
	MaxRetries     *int                   // max retries, when nil, there's no retry limit
	Interval       time.Duration          // wait duration between calls
	IsRetryableFns []func(err error) bool // functions to determine if error is retriable
}

type Config interface {
	String() string
	Value() interface{}
}

func New(cfg ...Config) *Retry {
	c := &Retry{
		Timeout:        CONFIG_TIMEOUT_DEFAULT,
		MaxRetries:     nil,
		Interval:       CONFIG_INTERVAL_DEFAULT,
		IsRetryableFns: []func(err error) bool{IsRetryableDefault},
	}
	return c.WithConfig(cfg...)
}

func (c *Retry) WithConfig(cfgs ...Config) *Retry {
	for _, cfg := range cfgs {
		switch cfg.String() {
		case CONFIG_TIMEOUT:
			c.Timeout = cfg.Value().(time.Duration)
		case CONFIG_INTERVAL:
			c.Interval = cfg.Value().(time.Duration)
		case CONFIG_MAX_RETRIES:
			c.MaxRetries = cfg.Value().(*int)
		case CONFIG_IS_RETRYABLE:
			c.IsRetryableFns = cfg.Value().([]func(err error) bool)
		}
	}
	return c
}
