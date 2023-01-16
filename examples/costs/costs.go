package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
	costs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/costs/v1.0/generated"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/google/uuid"
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

	params := &costs.GetProjectCostsParams{}
	res, err := c.Costs.GetProjectCostsWithResponse(
		ctx,
		uuid.MustParse("Customer Account ID"), // update to relevat Customer Account ID
		uuid.MustParse("Project ID"),          // update to relevant Project ID
		params,
	)

	// check for errors or empty response
	if agg := validate.Response(res, err, "JSON200"); agg != nil {
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
