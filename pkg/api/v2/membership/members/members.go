// package members is intended for managing members in STACKIT resources
// such as projects and organizations

package members

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// constants
const (
	apiPath                 = consts.API_PATH_MEMBERSHIP_V2_MEMBERS
	apiPathWithResourceType = consts.API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_TYPE_MEMBERS
)

// New returns a new handler for the service
func New(c common.Client) *MembersService {
	return &MembersService{
		Client: c,
	}
}

// Public types

// MembersService is the service that handles
// CRUD functionality for membership
type MembersService common.Service

// Member struct represents a single member
type Member struct {
	Subject   string     `json:"subject"`
	Role      string     `json:"role"`
	Condition *Condition `json:"condition,omitempty"`
}

// Condition struct for member
type Condition struct {
	ExpiresAt string `json:"expiresAt"`
}

// ResourceMembers struct represents member in a resource
type ResourceMembers struct {
	ResourceID   string   `json:"resourceId,omitempty"`
	ResourceType string   `json:"resourceType,omitempty"`
	Members      []Member `json:"members"`
}

// Implementation

// Get returns the members belonging to a resource
// Reference: https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/get-members
func (svc *MembersService) Get(ctx context.Context, resourceID, resourceType string) (members ResourceMembers, err error) {
	if err = validate.ResourceType(resourceType); err != nil {
		err = validate.WrapError(err)
		return
	}
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathWithResourceType, resourceType, resourceID), nil)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &members)
	return
}

// add members

// AddMembersRequest structure representing the request body for adding or updating members in a resource
type AddMembersRequest struct {
	ResourceType string   `json:"resourceType"`
	Members      []Member `json:"members"`
}

func (svc *MembersService) buildRequest(resourceType string, members []Member) ([]byte, error) {
	return json.Marshal(AddMembersRequest{
		ResourceType: resourceType,
		Members:      members,
	})
}

// Add members to a resource
// https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/patch-members
func (svc *MembersService) Add(ctx context.Context, resourceID, resourceType string, members []Member) (res ResourceMembers, err error) {
	if err = validate.ResourceType(resourceType); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildRequest(resourceType, members)

	params := url.Values{}
	params.Add("resourceType", resourceType)
	req, err := svc.Client.Request(ctx, http.MethodPatch, fmt.Sprintf(apiPath, resourceID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Remove removes members from a resource
// https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/post-members-remove
func (svc *MembersService) Remove(ctx context.Context, resourceID, resourceType string, members []Member) (res ResourceMembers, err error) {
	if err = validate.ResourceType(resourceType); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildRequest(resourceType, members)
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(consts.API_PATH_MEMBERSHIP_V2_REMOVE, resourceID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Replace overrides resource members
// https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/put-members
func (svc *MembersService) Replace(ctx context.Context, resourceID, resourceType string, members []Member) (res ResourceMembers, err error) {
	if err = validate.ResourceType(resourceType); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildRequest(resourceType, members)
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPath, resourceID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}

// Validate validates if members are allowed to be added to a resource
// If one or more members are not allowed, an error is returned
// https://api.stackit.schwarz/membership-service/openapi.v2.html#operation/post-members-validate
func (svc *MembersService) Validate(ctx context.Context, resourceID, resourceType string, members []Member) (res ResourceMembers, err error) {
	if err = validate.ResourceType(resourceType); err != nil {
		err = validate.WrapError(err)
		return
	}

	body, _ := svc.buildRequest(resourceType, members)
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(consts.API_PATH_MEMBERSHIP_V2_VALIDATE, resourceID), body)
	if err != nil {
		return
	}

	_, err = svc.Client.LegacyDo(req, &res)
	return
}
