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

	res, err := c.Services.Kubernetes.ProviderOptions.GetProviderOptionsWithResponse(ctx)
	if err != nil {
		panic(fmt.Sprintf("preparing request failed: %s", err))
	}
	if res.HasError != nil {
		panic(fmt.Sprintf("request failed: %s", res.HasError))
	}

	if res.JSON200.AvailabilityZones == nil || len(*res.JSON200.AvailabilityZones) == 0 {
		fmt.Println("No Kubernetes availability zones found")
		return
	}

	fmt.Println("We found the following Kubernetes availability zones:")
	for _, zone := range *res.JSON200.AvailabilityZones {
		if zone.Name == nil {
			continue
		}
		fmt.Printf("- %s\n", *zone.Name)
	}

}
