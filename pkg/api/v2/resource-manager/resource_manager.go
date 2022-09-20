// package resourcemanager groups together services that implement STACKIT's Resource Manager v2 API

package resourcemanager

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-manager/organizations"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-manager/projects"
)

// New returns a new handler for the service
func New(c common.Client) *ResourceManagerService {
	return &ResourceManagerService{
		Organizations: organizations.New(c),
		Projects:      projects.New(c),
	}
}

// ResourceManagerService is the service that handles
// project, organization and folder related services
type ResourceManagerService struct {
	Organizations *organizations.OrganizationsService
	Projects      *projects.ProjectsService
}
