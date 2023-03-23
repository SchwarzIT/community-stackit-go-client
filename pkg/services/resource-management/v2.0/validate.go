package resourcemanagement

import (
	"fmt"
)

// ValidateRole validates a role
func ValidateRole(role ProjectMemberRole) error {
	switch role {
	case PROJECT_ADMIN:
	case PROJECT_OWNER:
	case PROJECT_AUDITOR:
	case PROJECT_MEMBER:
	default:
		return fmt.Errorf("invalid role %s ", string(role))
	}
	return nil
}
