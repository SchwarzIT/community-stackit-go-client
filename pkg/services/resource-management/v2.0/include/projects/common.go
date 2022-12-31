// this file is only used to prevent wait.go
// from showing errors

package projects

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated/projects"

type CreateResponse struct {
	projects.ClientWithResponsesInterface
}

type DeleteResponse struct {
	projects.ClientWithResponsesInterface
}
