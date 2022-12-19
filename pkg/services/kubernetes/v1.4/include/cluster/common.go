// this file is only used to prevent wait.go
// from showing errors

package cluster

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.4/generated/cluster"

type CreateOrUpdateClusterResponse struct {
	cluster.ClientWithResponsesInterface
}

type DeleteClusterResponse struct {
	cluster.ClientWithResponsesInterface
}
