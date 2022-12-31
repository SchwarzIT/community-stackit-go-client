// package objectstorage groups together STACKIT Object Storage (S3) functionalities
// Such as creating buckets, credentials, and enabling Object Storage in projects

package objectstorage

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	keys "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/access-keys"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/buckets"
	credentialsgroup "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/credentials-group"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/projects"
)

// New returns a new handler for the service
func New(c common.Client) *ObjectStorageService {
	return &ObjectStorageService{
		Projects:         projects.New(c),
		Buckets:          buckets.New(c),
		AccessKeys:       keys.New(c),
		CredentialsGroup: credentialsgroup.New(c),
	}
}

// ObjectStorageService is the service that handles
// Object Storage related services
type ObjectStorageService struct {
	Projects         *projects.ObjectStorageProjectsService
	Buckets          *buckets.ObjectStorageBucketsService
	AccessKeys       *keys.ObjectStorageAccessKeysService
	CredentialsGroup *credentialsgroup.ObjectStorageCredentialsGroupService
}
