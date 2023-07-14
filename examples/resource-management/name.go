package main

import (
	"context"
	"fmt"
	"os"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewClientWithKeyAuth(ctx)

	fmt.Println("looking for project name..")

	res, err := c.ResourceManagement.Get(ctx, os.Getenv("STACKIT_PROJECT_ID"), &resourcemanagement.GetParams{})
	if err = validate.Response(res, err, "JSON200"); err != nil {
		panic(err)
	}

	fmt.Println(res.JSON200.Name)
}
