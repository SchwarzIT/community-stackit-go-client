package main

import (
	"context"
	"fmt"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c, err := stackit.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	projectID := "123-456-789"
	bucketName := "bucket"

	res, err := c.ObjectStorage.Bucket.Create(ctx, projectID, bucketName)
	if agg := validate.Response(res, err); agg != nil {
		panic(err)
	}

	process := res.WaitHandler(ctx, c.ObjectStorage.Bucket, projectID, bucketName)
	if _, err := process.WaitWithContext(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("bucket '%s' created successfully", bucketName)
}
