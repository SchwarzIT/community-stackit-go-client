package retry

import "time"

const (
	// timeout configuration constants
	CONFIG_TIMEOUT         = "Timeout"
	CONFIG_TIMEOUT_DEFAULT = 60 * time.Minute
)

// Timeout is the config struct
type Timeout struct {
	Duration time.Duration
}

// SetTimeout sets the maximum run duration
func (c *Retry) SetTimeout(d time.Duration) *Retry {
	return c.withConfig(&Timeout{
		Duration: d,
	})
}

var _ = Config(&Timeout{})

// String return the name of the config
func (c *Timeout) String() string {
	return CONFIG_TIMEOUT
}

// Value returns the defined value
func (c *Timeout) Value() interface{} {
	return c.Duration
}
