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

type IsRetryable struct {
	fnList []func(err error) bool
}

// Helper function
func (c *Retry) SetIsRetryable(override bool, f ...func(err error) bool) *Retry {
	return c.WithConfig(&IsRetryable{
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
