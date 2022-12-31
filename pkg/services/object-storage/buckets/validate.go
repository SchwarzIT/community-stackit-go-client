package buckets

import (
	"fmt"
	"regexp"
)

// ValidateBucketName validates a given bucket name
func ValidateBucketName(name string) error {
	exp := `^[a-z0-9-]{3,63}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid bucket name. valid name is of: %s", exp)
	}
	return nil
}
