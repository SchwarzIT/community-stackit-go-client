package projects

import (
	"fmt"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated/projects"
)

// ValidateRole validates a role
func ValidateRole(role projects.ProjectMemberRole) error {
	switch role {
	case projects.PROJECT_ADMIN:
	case projects.PROJECT_OWNER:
	case projects.PROJECT_AUDITOR:
	case projects.PROJECT_MEMBER:
	default:
		return fmt.Errorf("invalid role %s ", string(role))
	}
	return nil
}
