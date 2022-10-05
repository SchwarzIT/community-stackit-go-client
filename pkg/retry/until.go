package retry

import "net/http"

const (
	// Until configuration constants
	CONFIG_UNTIL = "UntilFns"
)

// untilFns is the config struct
type untilFns struct {
	fnList []func(*http.Response) bool
}

// Until sets functions that determin if the response given is the expected response
// this functionality is useful, for example, when waiting for an API status to be ready
func (c *Retry) Until(f ...func(*http.Response) bool) *Retry {
	return c.withConfig(&untilFns{
		fnList: f,
	})
}

var _ = Config(&untilFns{})

func (c *untilFns) String() string {
	return CONFIG_UNTIL
}

func (c *untilFns) Value() interface{} {
	return c.fnList
}
