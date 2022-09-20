package credentialsgroup

import (
	"fmt"
	"regexp"
)

// ValidateDisplayName validates a given credentialsGroup name
func ValidateDisplayName(name string) error {
	minLen, maxLen := 1, 32
	if len(name) < minLen || len(name) > maxLen {
		return fmt.Errorf("invalid display name. Length must be between %v and %v", minLen, maxLen)
	}
	return nil
}

// ValidateCredentialsGroupID validates a given  credentials group ID
func ValidateCredentialsGroupID(id string) error {
	exp := `^[a-z0-9-]{36}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(id) {
		return fmt.Errorf("invalid credentialsGroup id. valid id is of: %s", exp)
	}
	return nil
}
