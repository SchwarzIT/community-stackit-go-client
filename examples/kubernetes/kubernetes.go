package main

import (
	"context"
	"fmt"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewClientWithKeyAuth(ctx)

	res, err := c.Kubernetes.ProviderOptions.List(ctx)
	if err = validate.Response(res, err, "JSON200.AvailabilityZones"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("STACKIT Kubernetes Engine (SKE) availability zones:")
	for _, zone := range *res.JSON200.AvailabilityZones {
		if zone.Name == nil {
			continue
		}
		fmt.Printf("- %s\n", *zone.Name)
	}
}
