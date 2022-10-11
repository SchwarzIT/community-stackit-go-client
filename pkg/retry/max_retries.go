package retry

const (
	// name of the configuration
	CONFIG_MAX_RETRIES = "MaxRetries"
)

// maxRetries is the config struct
type maxRetries struct {
	Retries *int
}

// SetMaxRetries sets the maximum retries
func (c *Retry) SetMaxRetries(r int) *Retry {
	return c.withConfig(&maxRetries{
		Retries: &r,
	})
}

var _ = Config(&maxRetries{})

// String returns the config name
func (c *maxRetries) String() string {
	return CONFIG_MAX_RETRIES
}

// Value return the value
func (c *maxRetries) Value() interface{} {
	return c.Retries
}
