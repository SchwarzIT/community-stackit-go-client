// this file is only used to prevent wait.go
// from showing errors

package cluster

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.1/cluster"

type CreateOrUpdateResponse struct {
	cluster.ClientWithResponsesInterface
}

type DeleteResponse struct {
	cluster.ClientWithResponsesInterface
}
