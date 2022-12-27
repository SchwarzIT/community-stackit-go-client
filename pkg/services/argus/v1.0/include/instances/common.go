package instances

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/instances"

type InstanceCreateResponse struct {
	instances.ClientWithResponsesInterface
}

type InstanceUpdateResponse struct {
	instances.ClientWithResponsesInterface
}

type InstanceDeleteResponse struct {
	instances.ClientWithResponsesInterface
}
