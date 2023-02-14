// this file is only used to prevent wait.go
// from showing errors

package instances

import instances "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0.1/generated/instances"

type CreateResponse struct {
	instances.ClientWithResponsesInterface
}

type UpdateResponse struct {
	instances.ClientWithResponsesInterface
}

type DeleteResponse struct {
	instances.ClientWithResponsesInterface
}
