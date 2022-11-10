package retry

import (
	"net/http"
	"strings"
)

const (
	// timeout configuration constants
	CONFIG_IS_RETRYABLE = "IsRetryable"

	// requesrt errors
	ERR_NO_SUCH_HOST   = "dial tcp: lookup"
	ERR_CLIENT_TIMEOUT = "Client.Timeout"
)

// IsRetryable is the config struct
type isRetryable struct {
	fnList []func(err error) bool
}

// SetIsRetryable sets functions that determin if an error can be retried or not
func (c *Retry) SetIsRetryable(f ...func(err error) bool) *Retry {
	return c.withConfig(&isRetryable{
		fnList: f,
	})
}

var _ = Config(&isRetryable{})

func (c *isRetryable) String() string {
	return CONFIG_IS_RETRYABLE
}

func (c *isRetryable) Value() interface{} {
	return c.fnList
}

// IsRetryableNoOp always retries
func IsRetryableNoOp(err error) bool {
	return true
}

// IsRetryableDefault
func IsRetryableDefault(err error) bool {
	if strings.Contains(err.Error(), http.StatusText(http.StatusBadRequest)) {
		return strings.Contains(err.Error(), ERR_NO_SUCH_HOST) || strings.Contains(err.Error(), ERR_CLIENT_TIMEOUT)
	}

	if strings.Contains(err.Error(), http.StatusText(http.StatusUnauthorized)) ||
		strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
		return false
	}

	return true
}
