// package kubernetes groups together STACKIT Kubernetes Engine (SKE) related functionalities
// such as cluster management, options retrieval and enabling SKE in projects

package kubernetes

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes/clusters"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes/options"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes/projects"
)

// New returns a new handler for the service
func New(c common.Client) *KubernetesService {
	return &KubernetesService{
		Clusters: clusters.New(c),
		Options:  options.New(c),
		Projects: projects.New(c),
	}
}

// KubernetesService is the service that handles
// SKE related services, such as clusters, credentials & operations
type KubernetesService struct {
	Clusters *clusters.KubernetesClusterService
	Options  *options.KubernetesOptionsService
	Projects *projects.KubernetesProjectsService
}
