// package instances is used to manange MongoDB Flex instances

package instances

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/wait"
)

// constants
const (
	apiPathList           = consts.API_PATH_MONGO_DB_FLEX_INSTANCES
	apiPathCreate         = consts.API_PATH_MONGO_DB_FLEX_INSTANCES
	apiPathWithInstanceID = consts.API_PATH_MONGO_DB_FLEX_INSTANCE
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
	Count int `json:"count,omitempty"`
	Items []struct {
		ID        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		ProjectID string `json:"projectId,omitempty"`
	} `json:"items,omitempty"`
}

// Flavor is a signle falvor struct
type Flavor struct {
	ID          string   `json:"id,omitempty"`
	Description string   `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	CPU         int      `json:"cpu,omitempty"`
	Memory      int      `json:"memory,omitempty"`
}

// InstanceResponse is the server response for Get call
type InstanceResponse struct {
	Item struct {
		ACL struct {
			Items []string `json:"items,omitempty"`
		} `json:"acl,omitempty"`
		BackupSchedule string  `json:"backupSchedule,omitempty"`
		Flavor         Flavor  `json:"flavor,omitempty"`
		ID             string  `json:"id,omitempty"`
		Name           string  `json:"name,omitempty"`
		ProjectID      string  `json:"projectId,omitempty"`
		Replicas       int     `json:"replicas,omitempty"`
		Status         string  `json:"status,omitempty"`
		Storage        Storage `json:"storage,omitempty"`
		Users          []struct {
			Database string   `json:"database,omitempty"`
			Hostname string   `json:"hostname,omitempty"`
			ID       string   `json:"id,omitempty"`
			Password string   `json:"password,omitempty"`
			Port     int      `json:"port,omitempty"`
			Roles    []string `json:"roles,omitempty"`
			URI      string   `json:"uri,omitempty"`
			Username string   `json:"username,omitempty"`
		} `json:"users,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"item,omitempty"`
}

// CreateRequest holds data for requesting new instance
type CreateRequest struct {
	ACL struct {
		Items []string `json:"items"`
	} `json:"acl"`
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
func (svc *MongoDBInstancesService) Get(ctx context.Context, projectID, instanceID string) (res InstanceResponse, err error) {
	req, err := svc.Client.Request(ctx, http.MethodGet, fmt.Sprintf(apiPathWithInstanceID, projectID, instanceID), nil)
	if err != nil {
		return
	}
	_, err = svc.Client.Do(req, &res)
	return
}

func (svc *MongoDBInstancesService) buildRequest(instanceName, flavorID string, storage Storage, version string, replicas int, backupSchedule string, labels, options map[string]string, acl ACL) ([]byte, error) {
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

// Create creates a new MongoDB instance and returns the server response (CreateResponse) and a wait handler
// which upon call to `Wait()` will wait until the instance is successfully created
// Wait() returns the full instance details (Instance) and error if it occurred
// See also https://api.stackit.schwarz/mongo-flex-service/openapi.html#tag/instance/paths/~1projects~1{projectId}~1instances/post
func (svc *MongoDBInstancesService) Create(ctx context.Context, projectID, instanceName, flavorID string,
	storage Storage, version string, replicas int,
	backupSchedule string, labels, options map[string]string, acl ACL,
) (res CreateResponse, w *wait.Handler, err error) {

	// build body
	data, _ := svc.buildRequest(instanceName, flavorID, storage, version, replicas, backupSchedule, labels, options, acl)

	// prepare request
	req, err := svc.Client.Request(ctx, http.MethodPost, fmt.Sprintf(apiPathCreate, projectID), data)
	if err != nil {
		return
	}

	// do request
	_, err = svc.Client.Do(req, &res)

	// create Wait service
	w = wait.New(svc.waitForCreation(ctx, projectID, res.ID))
	w.SetTimeout(1 * time.Hour)
	return res, w, err
}

func (svc *MongoDBInstancesService) waitForCreation(ctx context.Context, projectID, instanceID string) wait.WaitFn {
	return func() (res interface{}, done bool, err error) {
		s, err := svc.Get(ctx, projectID, instanceID)
		if err != nil {
			return nil, false, err
		}
		if s.Item.Status == consts.MONGODB_STATUS_READY {
			return s.Item, true, nil
		}
		return s.Item, false, nil
	}
}
