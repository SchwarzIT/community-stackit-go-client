package retry

import "time"

const (
	// timeout configuration constants
	CONFIG_TIMEOUT         = "Timeout"
	CONFIG_TIMEOUT_DEFAULT = 60 * time.Minute
)

type Timeout struct {
	Duration time.Duration
}

// Helper function
func (c *Retry) SetTimeout(d time.Duration) *Retry {
	return c.WithConfig(&Timeout{
		Duration: d,
	})
}

var _ = Config(&Timeout{})

func (c *Timeout) String() string {
	return CONFIG_TIMEOUT
}

func (c *Timeout) Value() interface{} {
	return c.Duration
}
