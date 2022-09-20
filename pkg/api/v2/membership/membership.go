// package membership is used to group together membership related services
// such as members, roles and permission services

package membership

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/members"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/permissions"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership/roles"
)

// New returns a new handler for the service
func New(c common.Client) *MembershipService {
	return &MembershipService{
		Members:     *members.New(c),
		Roles:       *roles.New(c),
		Permissions: *permissions.New(c),
	}
}

// MembershipService is the service that handles
// membership related services, such as members and roles of a resource
type MembershipService struct {
	Members     members.MembersService
	Roles       roles.RolesService
	Permissions permissions.PermissionsService
}
