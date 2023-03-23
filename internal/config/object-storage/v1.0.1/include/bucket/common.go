// this file is only used to prevent wait.go
// from showing errors

package bucket

import "github.com/SchwarzIT/community-stackit-go-client/pkg/services/object-storage/v1.0.1/bucket"

type CreateResponse struct {
	bucket.ClientWithResponsesInterface
}

type DeleteResponse struct {
	bucket.ClientWithResponsesInterface
}
