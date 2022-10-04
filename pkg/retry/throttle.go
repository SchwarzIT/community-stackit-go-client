package retry

import "time"

const (
	// timeout configuration constants
	CONFIG_THROTTLE         = "Throttle"
	CONFIG_THROTTLE_DEFAULT = 5 * time.Second
)

// Throttle is the config struct
type Throttle struct {
	Duration time.Duration
}

// SetThrottle sets the duration to wait between calls
func (c *Retry) SetThrottle(d time.Duration) *Retry {
	return c.withConfig(&Throttle{
		Duration: d,
	})
}

var _ = Config(&Throttle{})

// String return the name of the config
func (c *Throttle) String() string {
	return CONFIG_THROTTLE
}

// Value returns the defined value
func (c *Throttle) Value() interface{} {
	return c.Duration
}
