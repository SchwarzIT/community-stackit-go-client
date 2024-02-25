// this file is used for validating network data and properties

package iaas

import (
	"fmt"
	"regexp"
)

// ValidateNetworkName validates a given network name
func ValidateNetworkName(name string) error {
	if len(name) > 63 || name == "" {
		return fmt.Errorf("name bust be non-empty and < 64 characters")
	}

	exp := `^[A-Za-z0-9]+((-|_|\s|\.)[A-Za-z0-9]+)*$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid cluster name. valid name is of: %s", exp)
	}
	return nil
}
