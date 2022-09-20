// this file is used for validation functions

package projects

import (
	"fmt"
	"strconv"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
)

// ValidateList validates the filters & container information required for listing projects
func (svc *ProjectsService) ValidateList(containerParentID string, containerIDs []string, offset, limit, creationTime string) error {
	if len(containerIDs) == 0 && containerParentID == "" {
		return errors.New("At least one of 'containerParentID' or 'containerIDs' needs to be provided")
	}
	if offset != "" {
		if v, err := strconv.Atoi(offset); err != nil {
			return errors.Wrap(err, "bad offset value")
		} else if v < 0 {
			return errors.New("offset must be >= 0")
		}
	}
	if limit != "" {
		if v, err := strconv.Atoi(limit); err != nil {
			return errors.Wrap(err, "bad limit value")
		} else if v < 0 || v > 100 {
			return errors.New("limit must be between 0..100")
		}
	}
	if creationTime != "" {
		if err := validate.ISO8601(creationTime); err != nil {
			return errors.Wrap(err, "invalid creation time start value")
		}
	}
	return nil
}

// ValidateCreateData validates the data required for creating a project
func (svc *ProjectsService) ValidateCreateData(name string, labels map[string]string, members []ProjectMember) error {
	if err := validateOwnerExists(members); err != nil {
		return err
	}
	if err := validate.ProjectName(name); err != nil {
		return err
	}
	if err := validateMandatoryLabels(labels); err != nil {
		return err
	}
	return nil
}

// ValidateUpdateData validates the data required for updating a project
func (svc *ProjectsService) ValidateUpdateData(containerID, containerParentID, name string, labels map[string]string) error {
	if err := validate.ProjectName(name); err != nil {
		return err
	}
	// if err := validate.OrganizationID(parentID); err != nil {
	// 	return err
	// }
	// if err := validate.ProjectID(projectID); err != nil {
	// 	return err
	// }
	if err := validateMandatoryLabels(labels); err != nil {
		return err
	}
	return nil
}

// validateMembersInProjectCreateRequest validates that at least one project owner exists in the members slice
// Reference: https://api.stackit.schwarz/resource-management/openapi.v2.html#operation/post-projects
func validateOwnerExists(members []ProjectMember) error {
	for _, m := range members {
		if m.Role == consts.ROLE_PROJECT_OWNER {
			return nil
		}
	}
	return errors.New("at least one member with project.owner role is required")
}

func validateMandatoryLabels(labels map[string]string) error {
	if len(labels) > 100 {
		return errors.New("only up to 100 labels are allowed")
	}
	m := MandatoryLabels{BillingReference: "", Scope: ""}
	for label := range m.ToMap() {
		found := false
		for k, v := range labels {
			if label == k {
				found = true
			}
			if k == "billingReference" {
				if err := validate.BillingRef(v); err != nil {
					return err
				}
			}
		}
		if !found {
			return fmt.Errorf("couldn't find mandatory label '%s' in labels map", label)
		}
	}
	return nil
}
