package retry

import "time"

const (
	// timeout configuration constants
	CONFIG_TIMEOUT         = "Timeout"
	CONFIG_TIMEOUT_DEFAULT = 2 * time.Minute
)

// timeout is the config struct
type timeout struct {
	Duration time.Duration
}

// SetTimeout sets the maximum run duration
func (c *Retry) SetTimeout(d time.Duration) *Retry {
	return c.withConfig(&timeout{
		Duration: d,
	})
}

var _ = Config(&timeout{})

// String return the name of the config
func (c *timeout) String() string {
	return CONFIG_TIMEOUT
}

// Value returns the defined value
func (c *timeout) Value() interface{} {
	return c.Duration
}
