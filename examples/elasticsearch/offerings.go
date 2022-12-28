package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
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

	res, err := c.Services.ElasticSearch.Offerings.GetWithResponse(ctx, "example")
	if err != nil {
		panic(fmt.Sprintf("preparing request failed: %s", err))
	}
	if res.HasError != nil {
		panic(fmt.Sprintf("request failed: %s", res.HasError))
	}
	if res.JSON200 == nil {
		panic("received an empty response from API")
	}
	fmt.Println("Received the following offerings:")
	for _, o := range res.JSON200.Offerings {
		fmt.Printf("- %s\n", o.Name)
	}

}
