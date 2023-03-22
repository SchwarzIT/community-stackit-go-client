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

	res, err := c.Kubernetes.ProviderOptions.GetProviderOptions(ctx)
	if err = validate.Response(res, err, "JSON200.AvailabilityZones"); err != nil {
		panic(err)
	}

	if len(*res.JSON200.AvailabilityZones) == 0 {
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
