package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
)

func main() {
	c, err := client.New(context.Background(), &client.Config{
		ServiceAccountID: os.Getenv("STACKIT_SERVICE_ACCOUNT_ID"),
		Token:            os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN"),
		OrganizationID:   os.Getenv("STACKIT_CUSTOMER_ACCOUNT_ID"),
	})
	if err != nil {
		panic(err)
	}

	projectID := "1234-56789-101112"
	bucketName := "example"

	w, err := c.ObjectStorage.Buckets.Create(context.TODO(), projectID, bucketName)
	if _, err := w.Wait(); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("bucket '%s' created successfully", bucketName)
}
