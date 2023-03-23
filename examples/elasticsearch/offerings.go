package main

import (
	"context"
	"fmt"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	project, ctx := "", context.Background()
	c, err := stackit.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter Project ID:")
	fmt.Scanln(&project)
	res, err := c.ElasticSearch.Offerings.Get(ctx, project)
	if err = validate.Response(res, err, "JSON200"); err != nil {
		panic(err)
	}

	fmt.Println("Received the following offerings:")
	for _, o := range res.JSON200.Offerings {
		fmt.Printf("- %s\n", o.Name)
	}
}
