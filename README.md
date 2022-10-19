# Community Go Client for STACKIT

[![Go Report Card](https://goreportcard.com/badge/github.com/SchwarzIT/community-stackit-go-client)](https://goreportcard.com/report/github.com/SchwarzIT/community-stackit-go-client) [![Unit Tests](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml/badge.svg)](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml) [![Coverage Status](https://coveralls.io/repos/github/SchwarzIT/community-stackit-go-client/badge.svg?branch=main)](https://coveralls.io/github/SchwarzIT/community-stackit-go-client?branch=main) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens) [![GoDoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/SchwarzIT/community-stackit-go-client) [![License](https://img.shields.io/badge/License-Apache_2.0-lightgray.svg)](https://opensource.org/licenses/Apache-2.0)

<br />

The client is community-supported and not an official STACKIT release, it is maintained by internal Schwarz IT teams integrating with STACKIT

## Install

To install the latest stable release, run:

```
go get github.com/SchwarzIT/community-stackit-go-client@latest
```


## Usage example

To get started, a Service Account[^1] and a Customer Account[^2] must be in place

If you're not sure how to get this information, please contact [STACKIT support](https://support.stackit.cloud)

```Go
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
		OrganizationID:   os.Getenv("STACKIT_ORGANIZATION_ID"),
	})
	if err != nil {
		panic(err)
	}

	projectID := "1234-56789-101112"
	bucketName := "example"

	process, err := c.ObjectStorage.Buckets.Create(context.TODO(), projectID, bucketName)
	if err != nil {
		panic(err)
	}

	// wait for bucket to be created
	if _, err := process.Wait(); err != nil {
		panic(err)
	}

	fmt.Printf("bucket '%s' created successfully", bucketName)
}

```

Further usage examples can be found in [`terraform-provider-stackit`](https://github.com/SchwarzIT/terraform-provider-stackit) 


## Auto retry

The client can automatically retry failed calls using `pkg/retry`

To enable retry in the client, use `SetRetry` in the following way:

```Go
c, _ := client.New(context.Background(), &client.Config{})
c.SetRetry(retry.New())
```


[^1]: In order to use the client, a Service Account and Token must be created using [Service Account API](https://api.stackit.schwarz/service-account/openapi.v1.html#operation/post-projects-projectId-service-accounts-v2)<br />
After creation, assign roles to the Service Account using [Membership API](https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/post-organizations-organizationId-projects-projectId-roles-roleName-service-accounts)<br />
If your Service Account needs to operate outside the scope of your project, you may need to contact STACKIT to assign further permissions

<br />

[^2]: The Customer Account ID is also referred to as Organization ID
