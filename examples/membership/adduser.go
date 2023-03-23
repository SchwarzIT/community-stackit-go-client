package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/SchwarzIT/community-stackit-go-client"
	membership "github.com/SchwarzIT/community-stackit-go-client/pkg/services/membership/v2.0"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated/projects"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
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
	member := membership.Member{
		Subject: "user@host.name",
		Role:    string(projects.PROJECT_OWNER),
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
		if agg := validate.Response(res, err); agg != nil {
			fmt.Println(agg)
			continue
		}
		fmt.Printf("added to project %s\n", p)
	}
}
