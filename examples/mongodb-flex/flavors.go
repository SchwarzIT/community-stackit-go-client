package main

import (
	"context"
	"fmt"
	"os"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/mongodb-flex/v1.0/flavors"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewClientWithKeyAuth(ctx)

	fmt.Println("looking for the smallest mongodb-flex flavor..")
	res, err := c.MongoDBFlex.Flavors.List(ctx, os.Getenv("STACKIT_PROJECT_ID"))
	if err = validate.Response(res, err, "JSON200.Flavors"); err != nil {
		panic(err)
	}

	var smallestFlavor *flavors.InfraFlavor
	for _, flavor := range *res.JSON200.Flavors {
		flavor := flavor // bypass re-used var
		if flavor.ID == nil || flavor.Memory == nil || flavor.CPU == nil {
			continue
		}
		if smallestFlavor == nil ||
			(*flavor.Memory <= *smallestFlavor.Memory && *flavor.CPU <= *smallestFlavor.CPU) {
			smallestFlavor = &flavor
		}
	}

	if smallestFlavor != nil {
		fmt.Printf("Smallest flavor ID: %s\n", *smallestFlavor.ID)
	}
}
