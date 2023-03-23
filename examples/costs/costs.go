package main

import (
	"context"
	"fmt"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v2.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	c := stackit.NewClient(ctx)

	params := &costs.GetProjectCostsParams{}
	res, err := c.Costs.GetProjectCosts(
		ctx,
		uuid.MustParse("Customer Account ID"), // update to relevat Customer Account ID
		uuid.MustParse("Project ID"),          // update to relevant Project ID
		params,
	)
	if err != nil {
		panic(err)
	}

	// check for empty response
	if err := validate.Field(res, "JSON200"); err != nil {
		panic(err)
	}

	v := res.JSON200
	fmt.Printf("Costs for project %s (ID: %s)\n", v.ProjectName, v.ProjectID)
	if v.Services != nil {
		for _, s := range *v.Services {
			fmt.Printf("- service: %s (SKU: %s): Quantity: %d, charge: %f, discount: %f\n", s.ServiceName, s.Sku, s.TotalQuantity, s.TotalCharge, s.TotalDiscount)
		}
	}
}
