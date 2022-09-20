// This file is used to validate data in the instances package

package instances

import (
	"errors"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// Validate
func Validate(projectID, instanceName, planID string) error {
	if err := validate.ProjectID(projectID); err != nil {
		return err
	}
	if err := ValidateInstanceName(instanceName); err != nil {
		return err
	}
	if err := ValidatePlanID(planID); err != nil {
		return err
	}
	return nil
}

// ValidateInstanceName validates argus instance name
func ValidateInstanceName(instanceName string) error {
	if len(instanceName) < 1 || len(instanceName) > 200 {
		return errors.New("instance name must be between 1..200 chars")
	}
	return nil
}

// ValidatePlanID validates argus instance plan ID
func ValidatePlanID(planID string) error {
	if planID == "" {
		return errors.New("instance plan ID is mandatory")
	}
	return nil
}
