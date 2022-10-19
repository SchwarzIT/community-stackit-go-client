// package users is used to manange MongoDB Flex users

package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPathList   = consts.API_PATH_MONGO_DB_FLEX_USERS
	apiPathCreate = consts.API_PATH_MONGO_DB_FLEX_USER
	apiPathGet    = consts.API_PATH_MONGO_DB_FLEX_USER
)

// New returns a new handler for the service
func New(c common.Client) *MongoDBUsersService {
	return &MongoDBUsersService{
		Client: c,
	}
}

// MongoDBUsersService is the service that manages MongoDB Flex instances
type MongoDBUsersService common.Service

// ListResponse represents a list of users returned from the server
type ListResponse struct {
	Count int            `json:"count,omitempty"`
	Items []UserListItem `json:"items,omitempty"`
}

// UserListItem is an item in the Items list of ListResponse
type UserListItem struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// GetResponse is a struct representing the server's get response
type GetResponse struct {
	Item UserGetItem `json:"item,omitempty"`
}

// UserGetItem is an item in GetResponse
type UserGetItem struct {
	Database string   `json:"database,omitempty"`
	Hostname string   `json:"hostname,omitempty"`
	ID       string   `json:"id,omitempty"`
	Port     int      `json:"port,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Username string   `json:"username,omitempty"`
}

// User represents a user with access to the database
type User struct {
	Database string   `json:"database,omitempty"`
	Hostname string   `json:"hostname,omitempty"`
	ID       string   `json:"id,omitempty"`
	Password string   `json:"password,omitempty"`
	Port     int      `json:"port,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	URI      string   `json:"uri,omitempty"`
	Username string   `json:"username,omitempty"`
}

// CreateRequest holds data for requesting new user
type CreateRequest struct {
	Database string   `json:"database,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Username string   `json:"username,omitempty"`
}

// CreateResponse is the server response when creating a new user
type CreateResponse struct {
	Item User `json:"item,omitempty"`
}

// List returns a list of MongoDB Flex users
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/user/paths/~1projects~1{projectId}~1instances~1{instanceId}~1users/get
func (svc *MongoDBUsersService) List(ctx context.Context, projectID, instanceID string) (res ListResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathList, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Get returns the user information by project, instance ID and user ID
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/user/paths/~1projects~1{projectId}~1instances~1{instanceId}~1users~1{userId}/get
func (svc *MongoDBUsersService) Get(ctx context.Context, projectID, instanceID, userID string) (res GetResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathGet, projectID, instanceID, userID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Create adds a new MongoDB user and returns the server response (CreateResponse) and error if occurred
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/user/paths/~1projects~1{projectId}~1instances~1{instanceId}~1users~1{userId}/post
func (svc *MongoDBUsersService) Create(ctx context.Context, projectID, instanceID, userID, username, database string, roles []string) (res CreateResponse, err error) {

	// build body
	data, _ := svc.buildCreateRequest(username, database, roles)

	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPathCreate, projectID, instanceID, userID), data)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, &res)
	return
}

func (svc *MongoDBUsersService) buildCreateRequest(username, database string, roles []string) ([]byte, error) {
	return json.Marshal(CreateRequest{
		Username: username,
		Database: database,
		Roles:    roles,
	})
}
