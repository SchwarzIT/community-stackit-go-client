// package resourcemanager groups together services that implement STACKIT's Resource Manager v2 API

package resourcemanager

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/resource-management/projects"
)

// New returns a new handler for the service
func New(c common.Client) *ResourceManagementV1Service {
	return &ResourceManagementV1Service{
		Projects: projects.New(c),
	}
}

// ResourceManagementService is the service that handles
// project, organization and folder related services
type ResourceManagementV1Service struct {
	Projects *projects.ProjectService
}
