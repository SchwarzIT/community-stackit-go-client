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
)

const (
	DefaultClientTimeout         = time.Minute
	DefaultRetryMaxRetries       = 3
	DefaultRetryWaitBetweenCalls = 30 * time.Second
	DefaultRetryTimeout          = 2 * time.Minute
)

type RetryConfig struct {
	MaxRetries       int           // Max retries
	WaitBetweenCalls time.Duration // Time to wait between requests
	RetryTimeout     time.Duration // Max time to re-try
	ClientTimeout    time.Duration // HTTP Client timeout
}

func NewRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:       DefaultRetryMaxRetries,
		WaitBetweenCalls: DefaultRetryWaitBetweenCalls,
		RetryTimeout:     DefaultRetryTimeout,
		ClientTimeout:    DefaultClientTimeout,
	}
}

// do performs the request
func do(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error) {
	if cfg == nil {
		cfg = NewRetryConfig()
	}
	if client == nil {
		client = &http.Client{}
	}
	client.Timeout = cfg.ClientTimeout
	maxRetries := cfg.MaxRetries
	err = wait.PollImmediate(cfg.WaitBetweenCalls, cfg.RetryTimeout, wait.ConditionFunc(
		func() (bool, error) {
			resp, err = client.Do(req)
			if err != nil {
				if maxRetries > 0 {
					if validate.ErrorIsOneOf(err, ClientTimeoutErr, ClientContextDeadlineErr, ClientConnectionRefusedErr) ||
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
	)
	if err != nil {
		return resp, errors.Wrap(err, fmt.Sprintf("url: %s\nmethod: %s", req.URL.String(), req.Method))
	}

	return resp, err
}
