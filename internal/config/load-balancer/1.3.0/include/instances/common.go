// this file is only used to prevent wait.go
// from showing errors

package instances

import instances "github.com/SchwarzIT/community-stackit-go-client/pkg/services/load-balancer/1beta.0.0/instances"

type CreateResponse struct {
	instances.ClientWithResponsesInterface
}

type DeleteResponse struct {
	instances.ClientWithResponsesInterface
}
