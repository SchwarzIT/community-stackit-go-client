# Community Go Client for STACKIT

[![Go Report Card](https://goreportcard.com/badge/github.com/SchwarzIT/community-stackit-go-client)](https://goreportcard.com/report/github.com/SchwarzIT/community-stackit-go-client) [![UnitTests](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/go.yml/badge.svg)](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/go.yml) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens) [![GoDoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/SchwarzIT/community-stackit-go-client) [![License](https://img.shields.io/badge/License-Apache_2.0-lightgray.svg)](https://opensource.org/licenses/Apache-2.0)

<br />

This repo's goal is to create a go-based http client for consuming STACKIT APIs

The client is community-supported and not an official STACKIT release, it is maintained by internal Schwarz IT teams integrating with STACKIT


## Usage example

To get started, a Service Account[^1] and a Customer Account[^2] must be in place

If you're not sure how to get this information, please contact [STACKIT support](https://support.stackit.cloud)

```
package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
)

func main() {
	c, err := client.New(context.Background(), &client.Config{
		ServiceAccountID: os.Getenv("STACKIT_SERVICE_ACCOUNT_ID"), 
		Token:            os.Getenv("STACKIT_SERVICE_ACCOUNT_TOKEN"),
		OrganizationID:   os.Getenv("STACKIT_CUSTOMER_ACCOUNT_ID"), 
	})
	if err != nil {
		panic(err)
	}

	projectID := "1234-56789-101112"
	bucketName := "example"

	err = c.ObjectStorage.Buckets.Create(context.TODO(), projectID, bucketName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("bucket '%s' created successfully", bucketName)
}

```

Another usage example can be found in [`terraform-provider-stackit`](https://github.com/SchwarzIT/terraform-provider-stackit) which is built using the community client

[^1]: In order to use the client, a Service Account and Token must be created using [Service Account API](https://api.stackit.schwarz/service-account/openapi.v1.html#operation/post-projects-projectId-service-accounts-v2)<br />
After creation, assign roles to the Service Account using [Membership API](https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/post-organizations-organizationId-projects-projectId-roles-roleName-service-accounts)<br />
If your Service Account needs to operate outside the scope of your project, you may need to contact STACKIT to assign further permissions

<br />

[^2]: The Customer Account ID is also referred to as Organization ID
