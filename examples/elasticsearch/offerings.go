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

	res, err := c.ElasticSearch.Offerings.Get(ctx, "{project-id}")
	if err != nil {
		panic(err)
	}

	if err := validate.Field(res, "JSON200"); err != nil {
		panic(err)
	}

	fmt.Println("Received the following offerings:")
	for _, o := range res.JSON200.Offerings {
		fmt.Printf("- %s\n", o.Name)
	}
}
