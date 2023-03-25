package clients

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	// API configuration options:
	Environment = "STACKIT_ENV"
)

const (
	// Known error messages
	ClientTimeoutErr           = "Client.Timeout exceeded while awaiting headers"
	ClientContextDeadlineErr   = "context deadline exceeded"
	ClientConnectionRefusedErr = "connection refused"
	ClientEOFError             = "unexpected EOF"
	ClientGWTimeoutFError      = "Gateway Timeout"
)

const (
	DefaultClientTimeout         = time.Second * 10
	DefaultRetryMaxRetries       = 3
	DefaultRetryWaitBetweenCalls = 1 * time.Second
	DefaultRetryTimeout          = 2 * time.Minute
)

type RetryConfig struct {
	MaxRetries       int
	WaitBetweenCalls time.Duration
	Timeout          time.Duration
}

func NewRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:       DefaultRetryMaxRetries,
		WaitBetweenCalls: DefaultRetryWaitBetweenCalls,
		Timeout:          DefaultRetryTimeout,
	}
}

// do performs the request
func do(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error) {
	maxRetries := cfg.MaxRetries
	if err := wait.PollImmediate(cfg.WaitBetweenCalls, cfg.Timeout, wait.ConditionFunc(
		func() (bool, error) {
			resp, err = client.Do(req)
			if err != nil {
				if maxRetries > 0 {
					if validate.ErrorIsOneOf(err, ClientTimeoutErr, ClientContextDeadlineErr, ClientConnectionRefusedErr, ClientGWTimeoutFError) ||
						(req.Method == http.MethodGet && strings.Contains(err.Error(), ClientEOFError)) {

						// reduce retries counter and retry
						maxRetries = maxRetries - 1
						return false, nil
					}
				}
				return false, err
			}
			if maxRetries > 0 && resp != nil {
				if resp.StatusCode == http.StatusBadGateway ||
					resp.StatusCode == http.StatusGatewayTimeout ||
					resp.StatusCode == http.StatusInternalServerError {

					maxRetries = maxRetries - 1
					return false, nil
				}
			}
			return true, nil
		}),
	); err != nil {
		return resp, errors.Wrap(err, fmt.Sprintf("url: %s\nmethod: %s", req.URL.String(), req.Method))
	}

	return resp, err
}
