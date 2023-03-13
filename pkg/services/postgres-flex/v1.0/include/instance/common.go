// this file is only used to prevent wait.go
// from showing errors

package instance

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/instance"

type CreateResponse struct {
	instance.ClientWithResponsesInterface
}

type UpdateResponse struct {
	instance.ClientWithResponsesInterface
}

type PatchUpdateResponse struct {
	instance.ClientWithResponsesInterface
}

type DeleteResponse struct {
	instance.ClientWithResponsesInterface
}
