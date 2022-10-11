package retry

import "net/http"

const (
	// Wait configuration constants
	CONFIG_UNTIL = "Until"
)

type UntilFn func(*http.Response) (bool, error)

// until is the config struct
type until struct {
	fnList []UntilFn
}

// SetUntil sets functions that determin if the response given is the expected response
// this functionality is useful, for example, when waiting for an API status to be ready
func (c *Retry) SetUntil(f ...UntilFn) *Retry {
	return c.withConfig(&until{
		fnList: f,
	})
}

var _ = Config(&until{})

func (c *until) String() string {
	return CONFIG_UNTIL
}

func (c *until) Value() interface{} {
	return c.fnList
}
