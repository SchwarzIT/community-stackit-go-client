// this file is only used to prevent wait.go
// from showing errors

package project

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated/project"

type CreateProjectResponse struct {
	project.ClientWithResponsesInterface
}

type DeleteProjectResponse struct {
	project.ClientWithResponsesInterface
}
