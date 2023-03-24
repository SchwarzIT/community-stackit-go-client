package main

import (
	"context"
	"fmt"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewStaticTokenClient(ctx)

	projectID := "123-456-789"
	bucketName := "bucket"

	bucket := c.ObjectStorage.Bucket
	res, err := bucket.Create(ctx, projectID, bucketName)
	if agg := validate.Response(res, err); agg != nil {
		panic(err)
	}

	process := bucket.WaitForeCreate(ctx, projectID, bucketName)
	if _, err := process.WaitWithContext(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("bucket '%s' created successfully", bucketName)
}
