// this file is only used to prevent wait.go
// from showing errors

package project

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.1/project"

type CreateResponse struct {
	project.ClientWithResponsesInterface
}

type DeleteResponse struct {
	project.ClientWithResponsesInterface
}
