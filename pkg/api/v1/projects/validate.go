// This file is used for validation functions that validate project management data

package projects

import (
	"errors"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// ValidateProjectCreationRoles validates that the given users and roles are correctly defined
// to fit project creation requirements
// Reference: https://api.stackit.schwarz/resource-management/openapi.v1.html#operation/post-organizations-organizationId-projects
func ValidateProjectCreationRoles(roles []ProjectRole) error {
	if len(roles) == 0 {
		return errors.New("no role definition found. at least one role needs to be defined")
	}
	foundSA := false
	foundUser := false
	for _, role := range roles {
		if err := validate.Role(role.Name); err != nil {
			return err
		}
		if len(role.Users) > 1 || len(role.ServiceAccounts) > 1 {
			return errors.New("during project creation, up to 1 service account and/or user can be defined")
		}
		if len(role.Users) == 1 {
			if foundUser {
				return errors.New("up to 1 user can be defined during project creation")
			}
			foundUser = true
		}
		if len(role.ServiceAccounts) == 1 {
			if foundSA {
				return errors.New("up to 1 service account can be defined during project creation")
			}
			foundSA = true
		}
	}
	return nil
}
