package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/versions"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewClientWithKeyAuth(ctx)

	fmt.Println("looking for the postgres-flex versions..")
	versions, err := c.PostgresFlex.Versions.List(ctx, os.Getenv("STACKIT_PROJECT_ID"), &versions.ListParams{})
	if err = validate.Response(versions, err, "JSON200.Versions"); err != nil {
		panic(err)
	}

	opts := []string{}
	for _, v := range *versions.JSON200.Versions {
		opts = append(opts, fmt.Sprintf("`%s`", v))
	}
	fmt.Println(strings.Join(opts, ", "))

	fmt.Println("looking for the postgres-flex flavors..")
	res, err := c.PostgresFlex.Flavors.List(ctx, os.Getenv("STACKIT_PROJECT_ID"))
	if err = validate.Response(res, err, "JSON200.Flavors"); err != nil {
		panic(err)
	}

	opts = []string{}
	for _, flavor := range *res.JSON200.Flavors {
		if flavor.ID == nil || flavor.Memory == nil || flavor.Cpu == nil {
			continue
		}
		opts = append(opts, fmt.Sprintf("`%s` (%d CPU, %d Memory)", *flavor.ID, *flavor.Cpu, *flavor.Memory))
	}
	fmt.Println(strings.Join(opts, ", "))
}
