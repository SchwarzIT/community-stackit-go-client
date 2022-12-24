// this file is only used to prevent wait.go
// from showing errors

package instance

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/instance"

type PostInstancesResponse struct {
	instance.ClientWithResponsesInterface
}

type PutInstanceResponse struct {
	instance.ClientWithResponsesInterface
}

type DeleteInstanceResponse struct {
	instance.ClientWithResponsesInterface
}
