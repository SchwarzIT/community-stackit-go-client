package resourcemanagement

import (
	"fmt"

	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
)

// ValidateRole validates a role
func ValidateRole(role resourcemanagement.ProjectMemberRole) error {
	switch role {
	case resourcemanagement.PROJECT_ADMIN:
	case resourcemanagement.PROJECT_OWNER:
	case resourcemanagement.PROJECT_AUDITOR:
	case resourcemanagement.PROJECT_MEMBER:
	default:
		return fmt.Errorf("invalid role %s ", string(role))
	}
	return nil
}
