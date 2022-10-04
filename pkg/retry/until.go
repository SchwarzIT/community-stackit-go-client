package retry

import "net/http"

const (
	// Until configuration constants
	CONFIG_UNTIL = "UntilFns"
)

// UntilFns is the config struct
type UntilFns struct {
	fnList []func(*http.Response) bool
}

// Until sets functions that determin if the response given is the expected response
// this functionality is useful, for example, when waiting for an API status to be ready
func (c *Retry) Until(f ...func(*http.Response) bool) *Retry {
	return c.withConfig(&UntilFns{
		fnList: f,
	})
}

var _ = Config(&UntilFns{})

func (c *UntilFns) String() string {
	return CONFIG_UNTIL
}

func (c *UntilFns) Value() interface{} {
	return c.fnList
}
