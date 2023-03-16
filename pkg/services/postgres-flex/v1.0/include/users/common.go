// this file is only used to prevent wait.go
// from showing errors

package users

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated/users"

type DeleteUserResponse struct {
	users.ClientWithResponsesInterface
}
