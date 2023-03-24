package clients

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	// API configuration options:
	Environment = "STACKIT_ENV"

	// Auth flow options:
	// Service Account Token Flow
	ServiceAccountEmail = "STACKIT_SERVICE_ACCOUNT_EMAIL"
	ServiceAccountToken = "STACKIT_SERVICE_ACCOUNT_TOKEN"

	// Key Access Flow (1)
	ServiceAccountKey = "STACKIT_SERVICE_ACCOUNT_KEY"
	PrivateKey        = "STACKIT_PRIVATE_KEY"

	// Key Access Flow (2) using file paths
	ServiceAccountKeyPath = "STACKIT_SERVICE_ACCOUNT_KEY_PATH"
	PrivateKeyPath        = "STACKIT_PRIVATE_KEY_PATH"
)

const (
	// Known error messages
	ClientTimeoutErr           = "Client.Timeout exceeded while awaiting headers"
	ClientContextDeadlineErr   = "context deadline exceeded"
	ClientConnectionRefusedErr = "connection refused"
	ClientEOFError             = "unexpected EOF"
	ClientGWTimeoutFError      = "Gateway Timeout"
)

// do performs the request
func do(client *http.Client, req *http.Request, maxRetries int, retryWait, retryTimeout time.Duration) (resp *http.Response, err error) {
	if err := wait.PollImmediate(retryWait, retryTimeout, wait.ConditionFunc(
		func() (bool, error) {
			resp, err = client.Do(req)
			if err != nil {
				if maxRetries > 0 {
					if oneOfSubstr(err, ClientTimeoutErr, ClientContextDeadlineErr, ClientConnectionRefusedErr, ClientGWTimeoutFError) ||
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

func oneOfSubstr(err error, msgs ...string) bool {
	for _, m := range msgs {
		if strings.Contains(err.Error(), m) {
			return true
		}
	}
	return false
}
