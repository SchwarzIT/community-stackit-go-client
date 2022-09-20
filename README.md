# Community Go Client for STACKIT

<br />

🚀&nbsp; This repo's goal is to create a go-based http client for consuming STACKIT APIs

🦫&nbsp; The client is written in Go v1.18

☘️&nbsp; The client is community-supported and not an official STACKIT release, it is maintained by internal Schwarz IT teams integrating with STACKIT

<br />

## Usage example

In order to use the client, a Service Account needs to be created first. At the moment, this can be done strictly [from the API](https://api.stackit.schwarz/service-account/openapi.v1.html#operation/post-projects-projectId-service-accounts-v2).

The customer account ID of your company or team must also be known in advanced.

- If you're not sure how to get this information, please contact [STACKIT support](https://support.stackit.cloud)
- To use the Service Account it must be assigned relevant roles using the [Membership API](https://api.stackit.schwarz/membership-service/openapi.v1.html#operation/post-organizations-organizationId-projects-projectId-roles-roleName-service-accounts)
- If your Service Account needs to operate outside the scope of your project, you may need to contact STACKIT to assign further permissions

```
c, err := client.New(context.Background(), &config.Config{
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

```
