// package users is used for migrating Schwarz IT KG users to STACKIT
// this package is intended to be used by the authClient
// IMPORTANT: this package and the authClient will soon be removed

package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// Public types

// constants
const (
	apiPath = consts.API_PATH_SHADOW_USERS
)

// New returns a new handler for the service
func New(c common.Client) *UsersService {
	return &UsersService{
		Client: c,
	}
}

// UsersService is the service that handles
// CRUD functionality for STACKIT shadow users
type UsersService common.Service

// User struct holds important info about a shadow user
type User struct {
	UUID           string
	Email          string
	Origin         string
	OrganizationID string // Also called Customer Account
}

// Requests

type usersGetOrCreateUUIDReqBody struct {
	Email           string `json:"email"`
	Origin          string `json:"origin"`
	CustomerAccount string `json:"customer-account"`
}

// Responses

// ShadowUsersResBody is the response struct from the shadow api
type ShadowUsersResBody struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Origin   string `json:"origin"`
}

// Implementation

// Get returns the user's their UUID (this is a shadow user UUID)
// Reference https://api.stackit.schwarz/appcloud-shadow-user-creator/openapi.v1.html#section/Authentication
func (svc *UsersService) Get(ctx context.Context, email, origin string) (User, error) {
	// validate origin
	if err := validate.UserOrigin(origin); err != nil {
		return User{}, err
	}

	body, err := svc.usersBuildGetOrCreateUUIDBody(email, origin, svc.Client.OrganizationID())
	if err != nil {
		return User{}, err
	}

	req, err := svc.Client.Request(ctx, http.MethodPut, apiPath, body)
	if err != nil {
		return User{}, err
	}

	resBody := &ShadowUsersResBody{}
	if _, err = svc.Client.Do(req, resBody); err != nil {
		return User{}, err
	}

	return User{
		UUID:           resBody.UUID,
		Origin:         resBody.Origin,
		Email:          resBody.Username,
		OrganizationID: svc.Client.OrganizationID(),
	}, nil
}

func (svc *UsersService) usersBuildGetOrCreateUUIDBody(email, origin, organizationID string) ([]byte, error) {
	return json.Marshal(usersGetOrCreateUUIDReqBody{
		Email:           email,
		Origin:          origin,
		CustomerAccount: organizationID,
	})
}
