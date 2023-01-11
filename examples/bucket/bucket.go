package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c, err := client.New(ctx, client.Config{
		ServiceAccountEmail: os.Getenv("STACKIT_SERVICE_ACCOUNT_EMAIL"),
		ServiceAccountToken: os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN"),
	})
	if err != nil {
		panic(err)
	}

	projectID := "123-456-789"
	bucketName := "example"

	res, err := c.ObjectStorage.Bucket.CreateWithResponse(ctx, projectID, bucketName)
	if agg := validate.Response(res, err); agg != nil {
		panic(err)
	}

	process := res.WaitHandler(ctx, c.ObjectStorage.Bucket, projectID, bucketName)
	if _, err := process.WaitWithContext(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("bucket '%s' created successfully", bucketName)
}
