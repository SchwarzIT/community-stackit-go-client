// package instances is used to manange MongoDB Flex instances

package instances

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPathList   = consts.API_PATH_MONGO_DB_FLEX_INSTANCES
	apiPathCreate = consts.API_PATH_MONGO_DB_FLEX_INSTANCES
	apiPathGet    = consts.API_PATH_MONGO_DB_FLEX_INSTANCE
	apiPathUpdate = consts.API_PATH_MONGO_DB_FLEX_INSTANCE
)

// New returns a new handler for the service
func New(c common.Client) *MongoDBInstancesService {
	return &MongoDBInstancesService{
		Client: c,
	}
}

// MongoDBInstancesService is the service that manages MongoDB Flex instances
type MongoDBInstancesService common.Service

// ListResponse represents a list of instances returned from the server
type ListResponse struct {
	Count int                `json:"count,omitempty"`
	Items []ListResponseItem `json:"items,omitempty"`
}

// ListResponseItem is an item in the response item list
type ListResponseItem struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ProjectID string `json:"projectId,omitempty"`
}

// Flavor is a signle falvor struct
type Flavor struct {
	ID          string   `json:"id,omitempty"`
	Description string   `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	CPU         int      `json:"cpu,omitempty"`
	Memory      int      `json:"memory,omitempty"`
}

// GetResponse is the server response for Get call
type GetResponse struct {
	Item Instance `json:"item,omitempty"`
}

// UpdateResponse is the server response for an Update call
type UpdateResponse struct {
	Item Instance `json:"item,omitempty"`
}

// Instance is a struct representing an instance item
type Instance struct {
	ACL            ACL     `json:"acl,omitempty"`
	BackupSchedule string  `json:"backupSchedule,omitempty"`
	Flavor         Flavor  `json:"flavor,omitempty"`
	ID             string  `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	ProjectID      string  `json:"projectId,omitempty"`
	Replicas       int     `json:"replicas,omitempty"`
	Status         string  `json:"status,omitempty"`
	Storage        Storage `json:"storage,omitempty"`
	Users          []User  `json:"users,omitempty"`
	Version        string  `json:"version,omitempty"`
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

// CreateRequest holds data for requesting new instance
type CreateRequest struct {
	ACL            ACL               `json:"acl"`
	BackupSchedule string            `json:"backupSchedule"`
	FlavorID       string            `json:"flavorId"`
	Labels         map[string]string `json:"labels"`
	Name           string            `json:"name"`
	Options        map[string]string `json:"options"`
	Replicas       int               `json:"replicas"`
	Storage        Storage           `json:"storage"`
	Version        string            `json:"version"`
}

// ACL is the access list
type ACL struct {
	Items []string `json:"items"`
}

// Storage represents the instance storage configuration
type Storage struct {
	Class string `json:"class"`
	Size  int    `json:"size"`
}

// CreateResponse is the server response when creating a new Instance
type CreateResponse struct {
	ID string `json:"id,omitempty"`
}

// UpdateRequest holds data for updating instance
type UpdateRequest struct {
	ACL            ACL               `json:"acl"`
	BackupSchedule string            `json:"backupSchedule"`
	FlavorID       string            `json:"flavorId"`
	Labels         map[string]string `json:"labels"`
	Options        map[string]string `json:"options"`
}

// List returns a list of MongoDB Flex instances in project
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/instance/paths/~1projects~1{projectId}~1instances/get
func (svc *MongoDBInstancesService) List(ctx context.Context, projectID string) (res ListResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathList, projectID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Get returns the instance information by project and instance IDs
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/instance/paths/~1projects~1{projectId}~1instances~1{instanceId}/get
func (svc *MongoDBInstancesService) Get(ctx context.Context, projectID, instanceID string) (res GetResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathGet, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

// Create creates a new MongoDB instance and returns the server response (CreateResponse) and a wait handler
// which upon call to `Wait()` will wait until the instance is successfully created
// Wait() returns the full instance details (Instance) and error if it occurred
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/instance/paths/~1projects~1{projectId}~1instances/post
func (svc *MongoDBInstancesService) Create(ctx context.Context, projectID, instanceName, flavorID string,
	storage Storage, version string, replicas int,
	backupSchedule string, labels, options map[string]string, acl ACL,
) (res CreateResponse, w *wait.Handler, err error) {

	// build body
	data, _ := svc.buildCreateRequest(instanceName, flavorID, storage, version, replicas, backupSchedule, labels, options, acl)

	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPathCreate, projectID), data)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, &res)

	// create Wait service
	w = wait.New(svc.waitForCreateOrUpdate(ctx, projectID, res.ID))
	w.SetTimeout(1 * time.Hour)
	return res, w, err
}

func (svc *MongoDBInstancesService) buildCreateRequest(instanceName, flavorID string, storage Storage, version string, replicas int, backupSchedule string, labels, options map[string]string, acl ACL) ([]byte, error) {
	return json.Marshal(CreateRequest{
		Name:           instanceName,
		FlavorID:       flavorID,
		Storage:        storage,
		BackupSchedule: backupSchedule,
		Version:        version,
		Replicas:       replicas,
		Labels:         labels,
		Options:        options,
		ACL:            acl,
	})
}

func (svc *MongoDBInstancesService) waitForCreateOrUpdate(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.Item.Status == consts.MONGODB_STATUS_READY {
			return s.Item, true, nil
		}
		if s.Item.Status == consts.MONGODB_STATUS_FAILED {
			return s.Item, false, errors.New("received status FAILED from server")
		}
		return s.Item, false, nil
	}
}

// Update updates a MongoDB instance and returns the server response (UpdateResponse) and a wait handler
// which upon call to `Wait()` will wait until the instance is successfully updated
// Wait() returns the full instance details (Instance) and error if it occurred
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/instance/paths/~1projects~1{projectId}~1instances~1{instanceId}/put
func (svc *MongoDBInstancesService) Update(ctx context.Context, projectID, instanceID, flavorID string,
	backupSchedule string, labels, options map[string]string, acl ACL,
) (res UpdateResponse, w *wait.Handler, err error) {

	// build body
	data, _ := svc.buildUpdateRequest(flavorID, backupSchedule, labels, options, acl)

	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodPut, fmt.Sprintf(apiPathUpdate, projectID, instanceID), data)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, &res)

	// create Wait service
	w = wait.New(svc.waitForCreateOrUpdate(ctx, projectID, instanceID))
	w.SetTimeout(1 * time.Hour)
	return res, w, err
}

func (svc *MongoDBInstancesService) buildUpdateRequest(flavorID string, backupSchedule string, labels, options map[string]string, acl ACL) ([]byte, error) {
	return json.Marshal(UpdateRequest{
		FlavorID:       flavorID,
		BackupSchedule: backupSchedule,
		Labels:         labels,
		Options:        options,
		ACL:            acl,
	})
}

// Delete deletes a MongoDB instance and returns a wait handler and error if occured
// `Wait()` will wait until the instance is successfully deleted
// Wait() returns nil (empty response from server) and error if it occurred
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/instance/paths/~1projects~1{projectId}~1instances~1{instanceId}/put
func (svc *MongoDBInstancesService) Delete(ctx context.Context, projectID, instanceID string) (w *wait.Handler, err error) {
	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodDelete, fmt.Sprintf(apiPathUpdate, projectID, instanceID), nil)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, nil)

	// create Wait service
	w = wait.New(svc.waitForDeletion(ctx, projectID, instanceID))
	w.SetTimeout(1 * time.Hour)
	return w, err
}

func (svc *MongoDBInstancesService) waitForDeletion(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		if _, err := svc.Get(ctx, projectID, instanceID); err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				return nil, true, nil
			}
			return nil, false, err
		}
		return nil, false, nil
	}
}
