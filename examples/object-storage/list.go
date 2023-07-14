package main

import (
	"context"
	"fmt"
	"os"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewClientWithKeyAuth(ctx)

	fmt.Println("looking for the object storage buckets..")
	res, err := c.ObjectStorage.Bucket.List(ctx, os.Getenv("STACKIT_PROJECT_ID"))
	if err = validate.Response(res, err, "JSON200.Buckets"); err != nil {
		panic(err)
	}

	for _, bucket := range res.JSON200.Buckets {
		fmt.Println(bucket.Name)
	}
}
