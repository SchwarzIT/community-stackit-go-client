package retry

import "time"

const (
	// timeout configuration constants
	CONFIG_THROTTLE         = "Throttle"
	CONFIG_THROTTLE_DEFAULT = 5 * time.Second
)

type Throttle struct {
	Duration time.Duration
}

// Helper function
func (c *Retry) SetThrottle(d time.Duration) *Retry {
	return c.withConfig(&Throttle{
		Duration: d,
	})
}

var _ = Config(&Throttle{})

func (c *Throttle) String() string {
	return CONFIG_THROTTLE
}

func (c *Throttle) Value() interface{} {
	return c.Duration
}
