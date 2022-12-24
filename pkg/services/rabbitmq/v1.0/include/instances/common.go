// this file is only used to prevent wait.go
// from showing errors

package instances

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/rabbitmq/v1.0/generated/instances"

type ProvisionResponse struct {
	instances.ClientWithResponsesInterface
}

type UpdateResponse struct {
	instances.ClientWithResponsesInterface
}

type DeprovisionResponse struct {
	instances.ClientWithResponsesInterface
}
