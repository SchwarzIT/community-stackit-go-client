package retry

import (
	"net/http"
	"strings"
)

const (
	// timeout configuration constants
	CONFIG_IS_RETRYABLE = "IsRetryable"

	// requesrt errors
	ERR_NO_SUCH_HOST = "dial tcp: lookup"
)

// IsRetryable is the config struct
type IsRetryable struct {
	fnList []func(err error) bool
}

// SetIsRetryable sets functions that determin if an error can be retried or not
func (c *Retry) SetIsRetryable(f ...func(err error) bool) *Retry {
	return c.withConfig(&IsRetryable{
		fnList: f,
	})
}

var _ = Config(&IsRetryable{})

func (c *IsRetryable) String() string {
	return CONFIG_IS_RETRYABLE
}

func (c *IsRetryable) Value() interface{} {
	return c.fnList
}

// IsRetryableNoOp always retries
func IsRetryableNoOp(err error) bool {
	return true
}

// IsRetryableDefault
func IsRetryableDefault(err error) bool {
	if strings.Contains(err.Error(), http.StatusText(http.StatusBadRequest)) {
		return strings.Contains(err.Error(), ERR_NO_SUCH_HOST)
	}

	if strings.Contains(err.Error(), http.StatusText(http.StatusUnauthorized)) {
		return false
	}

	return true
}
