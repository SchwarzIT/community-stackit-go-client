package validate

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/google/uuid"
	parse "github.com/karrick/tparse/v2"
	"github.com/pkg/errors"
)

func WrapError(err error) error {
	return errors.Wrap(err, "client validation error (Bad Request)")
}

// UUID validates a given UUID
func UUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}
	return nil
}

// OrganizationID validates a given organization ID
func OrganizationID(orgID string) error {
	if err := UUID(orgID); err != nil {
		return errors.Wrap(err, "invalid UUID for organization")
	}
	return nil
}

// ProjectID validates a given project ID
func ProjectID(projectID string) error {
	if err := UUID(projectID); err != nil {
		return errors.Wrap(err, "invalid UUID for project")
	}
	return nil
}

// ProjectName validates a given project name
func ProjectName(name string) error {
	exp := `^[a-zA-Z][ a-zA-Z0-9_-]{1,39}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid project name. valid name is of: %s", exp)
	}
	return nil
}

// BillingRef validates a given billing reference
func BillingRef(billingRef string) error {
	exp := `^[a-zA-Z][a-zA-Z0-9_-]{1,29}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(billingRef) {
		return fmt.Errorf("invalid billing reference. valid reference is of: %s", exp)
	}
	return nil
}

// SemVer validates a given version
func SemVer(version string) error {
	if version == "" {
		return errors.New("version is empty")
	}
	exp := `^\d+\.\d+(?:\.\d+)?$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(version) {
		return fmt.Errorf("invalid version. valid version is of: %s", exp)
	}
	return nil
}

// Role validates a role
func Role(role string) error {
	switch role {
	case consts.ROLE_PROJECT_ADMIN:
	case consts.ROLE_PROJECT_OWNER:
	case consts.ROLE_PROJECT_AUDITOR:
	case consts.ROLE_PROJECT_MEMBER:
	default:
		return fmt.Errorf("invalid role %s ", role)
	}
	return nil
}

// ResourceType validates a resource type
func ResourceType(r string) error {
	switch r {
	case consts.RESOURCE_TYPE_ORG:
	case consts.RESOURCE_TYPE_PROJECT:
	default:
		return fmt.Errorf("invalid resource type %s ", r)
	}
	return nil
}

// UserOrigin validates a given user origin
// @TODO: remove after APIv1 is deprecated
func UserOrigin(origin string) error {
	if origin != consts.SCHWARZ_AUTH_ORIGIN {
		return fmt.Errorf("unsupported user origin: %s", origin)
	}
	return nil
}

// DefaultResponseErrorHandler is the default error handler used to check
// if a giving STACKIT API response returned an error
func DefaultResponseErrorHandler(resp *http.Response) error {
	if resp.StatusCode < 400 {
		return nil
	}

	var b []byte
	if resp.Body != nil {
		b, _ = io.ReadAll(resp.Body)
	}
	return fmt.Errorf(
		"call error:\nHTTP status code: %d\nHTTP status message: %s\nServer response: %s\n%s",
		resp.StatusCode,
		http.StatusText(resp.StatusCode),
		string(b),
		resp.Request.URL.String(),
	)
}

// ISO8601 Validates that given time is formatted as ISO 8601
func ISO8601(t string) error {
	isoFmt := "2006-01-02T15:04:05.999Z"
	_, err := time.Parse(isoFmt, t)
	if err != nil {
		return errors.Wrap(err, "couldn't parse given time as ISO8601")
	}
	return nil
}

// RFC3339 Validates that given time is formatted as RFC3339
func RFC3339(t string) error {
	_, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return errors.Wrap(err, "couldn't parse given time as RFC3339")
	}
	return nil
}

// Duration validates that a given string can be parsed as duration
// i.e. 5m, 60s, 1h
func Duration(s string) (time.Duration, error) {
	if s == "" {
		return 0, errors.New("can't parse empty string as duration")
	}
	return parse.AbsoluteDuration(time.Now(), s)
}
