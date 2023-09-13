// this file is only used to prevent wait.go
// from showing errors

package project

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/load-balancer/1beta.0.0/project"

type EnableProjectResponse struct {
	project.ClientWithResponsesInterface
}

type DisableProjectResponse struct {
	project.ClientWithResponsesInterface
}
