package retry

import "time"

const (
	// timeout configuration constants
	CONFIG_THROTTLE         = "Throttle"
	CONFIG_THROTTLE_DEFAULT = 5 * time.Second
)

// throttle is the config struct
type throttle struct {
	Duration time.Duration
}

// SetThrottle sets the duration to wait between calls
func (c *Retry) SetThrottle(d time.Duration) *Retry {
	return c.withConfig(&throttle{
		Duration: d,
	})
}

var _ = Config(&throttle{})

// String return the name of the config
func (c *throttle) String() string {
	return CONFIG_THROTTLE
}

// Value returns the defined value
func (c *throttle) Value() interface{} {
	return c.Duration
}
