package retry

const (
	// name of the configuration
	CONFIG_MAX_RETRIES = "MaxRetries"
)

// MaxRetries is the config struct
type MaxRetries struct {
	Retries *int
}

// SetMaxRetries sets the maximum retries
func (c *Retry) SetMaxRetries(r int) *Retry {
	return c.withConfig(&MaxRetries{
		Retries: &r,
	})
}

var _ = Config(&Timeout{})

// String returns the config name
func (c *MaxRetries) String() string {
	return CONFIG_MAX_RETRIES
}

// Value return the value
func (c *MaxRetries) Value() interface{} {
	return c.Retries
}
