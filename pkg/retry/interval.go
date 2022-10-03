package retry

import "time"

const (
	// timeout configuration constants
	CONFIG_INTERVAL         = "Interval"
	CONFIG_INTERVAL_DEFAULT = 5 * time.Second
)

type Interval struct {
	Duration time.Duration
}

// Helper function
func (c *Retry) SetInterval(d time.Duration) *Retry {
	return c.WithConfig(&Interval{
		Duration: d,
	})
}

var _ = Config(&Interval{})

func (c *Interval) String() string {
	return CONFIG_INTERVAL
}

func (c *Interval) Value() interface{} {
	return c.Duration
}
