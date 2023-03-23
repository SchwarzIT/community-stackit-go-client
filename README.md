# Community Go Client for STACKIT

[![Go Report Card](https://goreportcard.com/badge/github.com/SchwarzIT/community-stackit-go-client)](https://goreportcard.com/report/github.com/SchwarzIT/community-stackit-go-client) [![Unit Tests](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml/badge.svg)](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml) [![Coverage Status](https://coveralls.io/repos/github/SchwarzIT/community-stackit-go-client/badge.svg?branch=main)](https://coveralls.io/github/SchwarzIT/community-stackit-go-client?branch=main) [![GoDoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/SchwarzIT/community-stackit-go-client) [![License](https://img.shields.io/badge/License-Apache_2.0-lightgray.svg)](https://opensource.org/licenses/Apache-2.0)

&nbsp;

The client is community-supported and not an official STACKIT release, it is maintained by internal Schwarz IT teams integrating with STACKIT

## Install

To install the latest stable release, run:

```bash
go get github.com/SchwarzIT/community-stackit-go-client@latest
```

## Usage Example

In order to use the client, a STACKIT Service Account [must be created](https://api.stackit.schwarz/service-account/openapi.v1.html#operation/post-projects-projectId-service-accounts-v2) and have relevant roles [assigned to it](https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/post-organizations-organizationId-projects-projectId-roles-roleName-service-accounts).<br />
For further assistance, please contact [STACKIT support](https://support.stackit.cloud)

```Go
package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	project, ctx := "", context.Background()
	c, err := client.New(ctx, client.Config{
		ServiceAccountEmail: os.Getenv("STACKIT_SERVICE_ACCOUNT_EMAIL"),
		ServiceAccountToken: os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN"),
	})
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
```

### Further Examples

1. Under [`/examples`](https://github.com/SchwarzIT/community-stackit-go-client/tree/main/examples) directory
2. In our [`terraform-provider-stackit`](https://github.com/SchwarzIT/terraform-provider-stackit)


&nbsp;

## Working with API environments

In order to modify the API environment, set the `Environment` field to one of `dev`, `qa` or `prod`

- The `Environment` field is optional
- By default `prod` is being used

```Go
c, err := client.New(ctx, client.Config{
	// ...
	Environment: "qa"
})
```
