// package options is used to retrieve various options used for configuring DSA

package options_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	client "github.com/SchwarzIT/community-stackit-go-client"
	dataservices "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/data-services/options"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

// constants
const (
	apiPathOfferings = consts.API_PATH_DSA_OFFERINGS
	broker           = "example"
)

func TestMongoDBOptionsService_GetOfferings(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	dsa := dataservices.New(c, broker)

	projectID := "abc"

	mux.HandleFunc(fmt.Sprintf(apiPathOfferings, broker, projectID), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Error("wrong method")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"offerings": [
			  {
				"name": "string",
				"version": "11",
				"latest": true,
				"documentationUrl": "string",
				"description": "string",
				"quotaCount": 0,
				"imageUrl": "string",
				"schema": {
				  "create": {
					"parameters": {}
				  },
				  "update": {
					"parameters": {}
				  }
				},
				"plans": [
				  {
					"id": "string",
					"name": "string",
					"description": "string",
					"free": true
				  }
				]
			  }
			]
		  }`)
	})

	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes options.OfferingsResponse
		wantErr bool
	}{
		{"ok", args{context.Background(), projectID}, options.OfferingsResponse{
			Offerings: []options.Offer{
				{
					Name:             "string",
					Version:          "11",
					Latest:           true,
					DocumentationURL: "string",
					Description:      "string",
					QuotaCount:       0,
					ImageURL:         "string",
					Schema: options.Schema{
						Create: options.ActionSetup{
							Parameters: map[string]string{},
						},
						Update: options.ActionSetup{
							Parameters: map[string]string{},
						},
					},
					Plans: []options.Plan{
						{
							Name:        "string",
							ID:          "string",
							Description: "string",
							Free:        true,
						},
					},
				},
			},
		}, false},
		{"nil ctx", args{nil, projectID}, options.OfferingsResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := dsa.Options.GetOfferings(tt.args.ctx, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSAOptionsService.GetOfferings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) && !tt.wantErr {
				t.Errorf("DSAOptionsService.GetOfferings() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
