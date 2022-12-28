# Community Go Client for STACKIT

[![Go Report Card](https://goreportcard.com/badge/github.com/SchwarzIT/community-stackit-go-client)](https://goreportcard.com/report/github.com/SchwarzIT/community-stackit-go-client) [![Unit Tests](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml/badge.svg)](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml) [![Coverage Status](https://coveralls.io/repos/github/SchwarzIT/community-stackit-go-client/badge.svg?branch=main)](https://coveralls.io/github/SchwarzIT/community-stackit-go-client?branch=main) [![GoDoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/SchwarzIT/community-stackit-go-client) [![License](https://img.shields.io/badge/License-Apache_2.0-lightgray.svg)](https://opensource.org/licenses/Apache-2.0)

<br />

The client is community-supported and not an official STACKIT release, it is maintained by internal Schwarz IT teams integrating with STACKIT

## Install

To install the latest stable release, run:

```
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

	res, err := c.Services.Kubernetes.ProviderOptions.GetProviderOptionsWithResponse(ctx)
	if err != nil {
		panic(fmt.Sprintf("preparing request failed: %s", err))
	}
	if res.HasError != nil {
		panic(fmt.Sprintf("request failed: %s", res.HasError))
	}
	if res.JSON200 == nil {
		panic("received an empty response from API")
	}
	if res.JSON200.AvailabilityZones == nil || len(*res.JSON200.AvailabilityZones) == 0 {
		fmt.Println("No Kubernetes availability zones found")
		return
	}

	fmt.Println("We found the following Kubernetes availability zones:")
	for _, zone := range *res.JSON200.AvailabilityZones {
		if zone.Name == nil {
			continue
		}
		fmt.Printf("- %s\n", *zone.Name)
	}

}


```

Further usage examples can be found in [`terraform-provider-stackit`](https://github.com/SchwarzIT/terraform-provider-stackit) 
