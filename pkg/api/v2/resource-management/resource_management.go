// package resourcemanager groups together services that implement STACKIT's Resource Manager v2 API

package resourcemanager

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-management/organizations"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-management/projects"
)

// New returns a new handler for the service
func New(c common.Client) *ResourceManagementService {
	return &ResourceManagementService{
		Organizations: organizations.New(c),
		Projects:      projects.New(c),
	}
}

// ResourceManagementService is the service that handles
// project, organization and folder related services
type ResourceManagementService struct {
	Organizations *organizations.OrganizationsService
	Projects      *projects.ProjectsService
}
