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

	res, err := c.PostgresFlex.Versions.List(ctx, os.Getenv("STACKIT_PROJECT_ID"), &versions.ListParams{})
	if err = validate.Response(res, err, "JSON200.Versions"); err != nil {
		panic(err)
	}

	fmt.Printf(
		"Postgres flex enterprise supports the following versions: \n- %s\n",
		strings.Join(*res.JSON200.Versions, "\n- "),
	)
}
