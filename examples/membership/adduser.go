package main

import (
	"context"
	"fmt"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0"
	rm "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

func main() {
	ctx := context.Background()
	c := stackit.NewClient(ctx)
	member := membership.Member{
		Subject: "user@host.name",
		Role:    string(rm.PROJECT_OWNER),
	}
	body := membership.AddMembersJSONRequestBody{
		ResourceType: membership.RESOURCE_TYPE_PROJECT,
		Members:      []membership.Member{member},
	}
	projects := []string{
		// add project or container IDs here
		"123-456-789",
	}
	for _, p := range projects {
		res, err := c.Membership.AddMembers(ctx, p, body)
		if err = validate.Response(res, err); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("added to project %s\n", p)
	}
}
