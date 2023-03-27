// this file is only used to prevent wait.go
// from showing errors

package resourcemanagement

import resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"

type CreateResponse struct {
	resourcemanagement.ClientWithResponsesInterface
}

type DeleteResponse struct {
	resourcemanagement.ClientWithResponsesInterface
}
