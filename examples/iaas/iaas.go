package main

import (
	"context"
	"fmt"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/google/uuid"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	iaas "github.com/SchwarzIT/community-stackit-go-client/pkg/services/iaas-api/v1alpha"
)

func main() {
	ctx := context.Background()
	c := stackit.MustNewClientWithKeyAuth(ctx)

	l := 25

	servers := []iaas.V1IP{
		"8.8.8.8",
	}

	req := iaas.V1CreateNetworkJSONBody{
		// Name The name for a General Object. Matches Names and also UUIDs.
		Name: "SNA Network",

		// Nameservers List of DNS Servers/Nameservers.
		Nameservers:    &servers, //
		PrefixLengthV4: &l,
	}
	projectID := uuid.New()

	res, err := c.IAAS.V1CreateNetwork(ctx, projectID, iaas.V1CreateNetworkJSONRequestBody(req))
	if err = validate.Response(res, err, "JSON200.AvailabilityZones"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.JSON202.RequestID)
}
