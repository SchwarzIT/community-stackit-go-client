package retry

const (
	// name of the configuration
	CONFIG_MAX_RETRIES = "MaxRetries"
)

type MaxRetries struct {
	Retries *int
}

// Helper function
func (c *Retry) SetMaxRetries(r int) *Retry {
	return c.WithConfig(&MaxRetries{
		Retries: &r,
	})
}

var _ = Config(&Timeout{})

func (c *MaxRetries) String() string {
	return CONFIG_MAX_RETRIES
}

func (c *MaxRetries) Value() interface{} {
	return c.Retries
}
