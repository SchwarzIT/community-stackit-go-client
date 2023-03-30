package main

import (
	"context"
	"fmt"
	"os"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
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

	cpu := 0
	mem := 0
	id := ""
	for _, flavor := range *res.JSON200.Flavors {
		if flavor.ID == nil || flavor.CPU == nil || flavor.Memory == nil {
			continue
		}
		if id == "" {
			id = *flavor.ID
			cpu = *flavor.CPU
			mem = *flavor.Memory
		}
		if cpu >= *flavor.CPU && mem >= *flavor.Memory {
			id = *flavor.ID
			cpu = *flavor.CPU
			mem = *flavor.Memory
		}
	}
	if id == "" {
		fmt.Println("couldn't find smallest flavor")
		return
	}
	fmt.Printf("found flavor id %s with %d cpu and %d memory", id, cpu, mem)
}
