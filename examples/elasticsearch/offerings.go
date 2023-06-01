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

	res, err := c.ElasticSearch.Offerings.List(ctx, "my-project-id")
	if err = validate.Response(res, err, "JSON200"); err != nil {
		panic(err)
	}

	fmt.Println("Received the following offerings:")
	for _, o := range res.JSON200.Offerings {
		fmt.Printf("- %s\n", o.Name)
	}
}
