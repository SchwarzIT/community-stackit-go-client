package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
)

func main() {
	ctx := context.Background()
	c, err := client.New(ctx, &client.Config{
		ServiceAccountID: os.Getenv("STACKIT_SERVICE_ACCOUNT_ID"),
		Token:            os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN"),
		OrganizationID:   os.Getenv("STACKIT_ORGANIZATION_ID"),
	})
	if err != nil {
		panic(err)
	}

	projectID := "1234"
	bucketName := "example"

	process, err := c.ObjectStorage.Buckets.Create(ctx, projectID, bucketName)
	if err != nil {
		panic(err)
	}

	// wait for bucket to be created
	if _, err := process.Wait(); err != nil {
		panic(err)
	}

	fmt.Printf("bucket '%s' created successfully", bucketName)
}
