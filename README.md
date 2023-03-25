# Community Go Client for STACKIT

[![Go Report Card](https://goreportcard.com/badge/github.com/SchwarzIT/community-stackit-go-client)](https://goreportcard.com/report/github.com/SchwarzIT/community-stackit-go-client) [![Unit Tests](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml/badge.svg)](https://github.com/SchwarzIT/community-stackit-go-client/actions/workflows/tests.yml) [![Coverage Status](https://coveralls.io/repos/github/SchwarzIT/community-stackit-go-client/badge.svg?branch=main)](https://coveralls.io/github/SchwarzIT/community-stackit-go-client?branch=main) [![GoDoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/SchwarzIT/community-stackit-go-client) [![License](https://img.shields.io/badge/License-Apache_2.0-lightgray.svg)](https://opensource.org/licenses/Apache-2.0)

This is a Go client designed to help developers interact with STACKIT APIs. It is maintained by the STACKIT community within Schwarz IT.

&nbsp;

## Installation

To install the community-stackit-go-client package, run the following command:

```bash
go get github.com/SchwarzIT/community-stackit-go-client
```

&nbsp;

## Authentication

Before you can start using the client, you will need to create a STACKIT Service Account in your project and assign it the appropriate permissions (i.e. `project.owner`).

After the service account has been created, you can authenticate to the client using the `Key` authentication flow (recommended) or with the static `Token` flow (less secure as the token is long-lived).

### Key flow

1. Create an RSA key pair:

   ```bash
   openssl req -x509 -nodes -newkey rsa:2048 -days 365 \
      -keyout private_key.pem -out public_key.pem -subj "/CN=unused"
   ```

2. Create a service account key:

   - copy the private key to `examples/service-accounts`

   - Modify `create_sa_key.go` (fill out the consts)

   - Run with:  

        ```bash
        go run create_sa_key.go
        ```

3. Set environment variables:

   ```bash
   export STACKIT_SERVICE_ACCOUNT_KEY_PATH="sa_key.json"
   export STACKIT_PRIVATE_KEY_PATH="private_key.pem"

   # optionally modify the API environment to one of:
   # `dev`, `qa` or `prod` (default)
   export STACKIT_ENV=prod
   ```

4. Configure the client

   ```go
   package main

   import (
       "context"
       stackit "github.com/SchwarzIT/community-stackit-go-client"
   )

   func main() {
       ctx := context.Background()
       c := stackit.MustNewClientWithTokenAuth(ctx)
       // ...
   }
   ```

### Token flow

1. Set the following environment variables:

    ```bash
    export STACKIT_SERVICE_ACCOUNT_EMAIL=email
    export STACKIT_SERVICE_ACCOUNT_TOKEN=token

    # optionally modify the API environment to one of:
    # `dev`, `qa` or `prod` (default)
    export STACKIT_ENV=prod
    ```

2. Configure the client

   ```go
   package main

   import (
       "context"
       stackit "github.com/SchwarzIT/community-stackit-go-client"
   )

   func main() {
       ctx := context.Background()
       c := stackit.MustNewClientWithTokenAuth(ctx)
       // ...
   }
   ```

## Usage

Create a file called `example.go`:

```go
package main

import (
    "context"
    "fmt"

    stackit "github.com/SchwarzIT/community-stackit-go-client"
    "github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
    ctx := context.Background()
    c := stackit.MustNewClientWithTokenAuth(ctx)

    res, err := c.ElasticSearch.Offerings.Get(ctx, "my-project-id")
    if err = validate.Response(res, err, "JSON200"); err != nil {
        panic(err)
    }

    fmt.Println("Received the following offerings:")
    for _, o := range res.JSON200.Offerings {
        fmt.Printf("- %s\n", o.Name)
    }
}
```

Then, you can run the example with the following command:

```bash
go run example.go
```

### Further Examples

1. Under [`/examples`](https://github.com/SchwarzIT/community-stackit-go-client/tree/main/examples) directory
2. In our [`terraform-provider-stackit`](https://github.com/SchwarzIT/terraform-provider-stackit)

&nbsp;

## Contributing

If you find a bug or have an idea for a new feature, feel free to submit an issue or pull request!

Please make sure to include tests for any new functionality you add, and to run the existing tests before submitting your changes.

&nbsp;

## License

This project is licensed under the Apache-2.0 license - see the LICENSE file for details.
