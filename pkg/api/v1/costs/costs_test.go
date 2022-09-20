package costs_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/costs"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

var customerAccountRespBody string = `
[
    {
        "projectId": "ba04b091-ec32-4423-81e6-976f1b8a8363",
        "projectName": "sit.dev",
        "totalCharge": 201030.23,
        "totalDiscount": 150.00,
        "services":
        [
            {
                "totalQuantity": 321,
                "totalCharge": 201030.23,
                "totalDiscount": 150.00,
                "unitLabel": "Giga Hours",
                "sku": "ST-0010801",
                "serviceName": "SAP_DB",
                "serviceCategoryName": "Compute Engine",
                "reportData":
                [
                    {
                        "timePeriod":
                        {
                            "start": "2021-11-29",
                            "end": "2021-11-30"
                        },
                        "charge": 2010.30,
                        "discount": 2010.30,
                        "quantity": 123
                    }
                ]
            }
        ]
    }
]
`

func TestCostsService_GetCustomerAccountCosts(t *testing.T) {
	t.Run("successful response", func(t *testing.T) {
		client, mux, teardown, _ := client.MockServer()
		defer teardown()

		mux.HandleFunc("/costs-service/v1/costs/07a1ed91-2efb-42c2-9d00-e84ae71bce0d", func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("content-type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(customerAccountRespBody))
		})

		costsSvc := costs.New(client)

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		resp, err := costsSvc.GetCustomerAccountCosts(context.Background(), from, to, consts.COSTS_GRANULARITY_DAILY, consts.COSTS_DEPTH_AUTO)

		if err != nil {
			t.Errorf("wanted no error, got %v", err)
		}

		json, _ := json.Marshal(resp)
		want := string(json)
		got := customerAccountRespBody

		if strings.Compare(want, got) == 0 {
			t.Errorf("wanted \n%v\n, got \n%v", want, got)
		}

	})

	t.Run("not found response", func(t *testing.T) {
		client, _, teardown, _ := client.MockServer()
		defer teardown()

		costsSvc := costs.New(client)

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		_, err := costsSvc.GetCustomerAccountCosts(context.Background(), from, to, consts.COSTS_GRANULARITY_DAILY, consts.COSTS_DEPTH_AUTO)

		if err == nil {
			t.Errorf("wanted error, got nil")
		}
	})

	t.Run("server error response", func(t *testing.T) {
		client, mux, teardown, _ := client.MockServer()
		defer teardown()

		mux.HandleFunc("/costs-service/v1/costs/07a1ed91-2efb-42c2-9d00-e84ae71bce0d", func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("content-type", "application/json")
			res.WriteHeader(http.StatusInternalServerError)
		})

		costsSvc := costs.New(client)

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		_, err := costsSvc.GetCustomerAccountCosts(context.Background(), from, to, consts.COSTS_GRANULARITY_DAILY, consts.COSTS_DEPTH_AUTO)

		if err == nil {
			t.Errorf("wanted error, got nil")
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		client, mux, teardown, _ := client.MockServer()
		defer teardown()

		mux.HandleFunc("/costs-service/v1/costs/07a1ed91-2efb-42c2-9d00-e84ae71bce0d", func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("content-type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(customerAccountRespBody))
		})

		ctx, cancel := context.WithCancel(context.TODO())
		cancel()

		costsSvc := costs.New(client)

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		_, err := costsSvc.GetCustomerAccountCosts(ctx, from, to, consts.COSTS_GRANULARITY_DAILY, consts.COSTS_DEPTH_AUTO)

		if err == nil {
			t.Error("wanted error, got nil")
		}
	})
}

var projectResponseBody string = `
	{
		"projectId": "ba04b091-ec32-4423-81e6-976f1b8a8363",
		"projectName": "sit.dev",
		"totalCharge": 201030.23,
		"totalDiscount": 150.00,
		"services": [
			{
				"totalQuantity": 321,
				"totalCharge": 201030.23,
				"totalDiscount": 150.00,
				"unitLabel": "Giga Hours",
				"sku": "ST-0010801",
				"serviceName": "SAP_DB",
				"serviceCategoryName": "Compute Engine",
				"reportData": [
					{
						"timePeriod": {
							"start": "2021-11-29",
							"end": "2021-11-30"
						},
						"charge": 2010.30,
						"discount": 2010.30,
						"quantity": 123
					}
				]
			}
		]
	}
`

func TestCostsService_GetProjectCosts(t *testing.T) {
	t.Run("wrong projectId argument", func(t *testing.T) {
		client, _, teardown, _ := client.MockServer()
		defer teardown()

		costsSvc := costs.New(client)
		_, err := costsSvc.GetProjectCosts(
			context.Background(),
			"invalid-id",
			time.Now().AddDate(0, 0, -30),
			time.Now(),
			consts.COSTS_GRANULARITY_DAILY,
			consts.COSTS_DEPTH_PROJECT,
		)

		if err == nil {
			t.Error("wanted error, got nil")
		}
	})

	t.Run("successful response", func(t *testing.T) {
		client, mux, teardown, _ := client.MockServer()
		defer teardown()

		mux.HandleFunc("/costs-service/v1/costs/07a1ed91-2efb-42c2-9d00-e84ae71bce0d/projects/ba04b091-ec32-4423-81e6-976f1b8a8363", func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("content-type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(projectResponseBody))
		})

		costsSvc := costs.New(client)

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		resp, err := costsSvc.GetProjectCosts(
			context.Background(),
			"ba04b091-ec32-4423-81e6-976f1b8a8363",
			from,
			to,
			consts.COSTS_GRANULARITY_DAILY,
			consts.COSTS_DEPTH_PROJECT,
		)

		if err != nil {
			t.Errorf("wanted no error, got %v", err)
		}

		json, _ := json.Marshal(resp)
		want := string(json)
		got := projectResponseBody

		if strings.Compare(want, got) == 0 {
			t.Errorf("wanted \n%v\n, got \n%v", want, got)
		}

	})

	t.Run("canceled context", func(t *testing.T) {
		client, mux, teardown, _ := client.MockServer()
		defer teardown()

		mux.HandleFunc("/costs-service/v1/costs/07a1ed91-2efb-42c2-9d00-e84ae71bce0d/projects/ba04b091-ec32-4423-81e6-976f1b8a8363", func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("content-type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(projectResponseBody))
		})

		costsSvc := costs.New(client)

		ctx, cancel := context.WithCancel(context.TODO())
		cancel()

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		_, err := costsSvc.GetProjectCosts(
			ctx,
			"ba04b091-ec32-4423-81e6-976f1b8a8363",
			from,
			to,
			consts.COSTS_GRANULARITY_DAILY,
			consts.COSTS_DEPTH_SERVICE,
		)

		if err == nil {
			t.Error("wanted error, got nil")
		}

	})

	t.Run("not found response", func(t *testing.T) {
		client, _, teardown, _ := client.MockServer()
		defer teardown()

		costsSvc := costs.New(client)

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		_, err := costsSvc.GetProjectCosts(
			context.Background(),
			"ba04b091-ec32-4423-81e6-976f1b8a8363",
			from,
			to,
			consts.COSTS_GRANULARITY_DAILY,
			consts.COSTS_DEPTH_PROJECT,
		)

		if err == nil {
			t.Errorf("wanted error, got nil")
		}
	})

	t.Run("server error response", func(t *testing.T) {
		client, mux, teardown, _ := client.MockServer()
		defer teardown()

		mux.HandleFunc("/costs-service/v1/costs/07a1ed91-2efb-42c2-9d00-e84ae71bce0d/projects/ba04b091-ec32-4423-81e6-976f1b8a8363", func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("content-type", "application/json")
			res.WriteHeader(http.StatusInternalServerError)
		})

		costsSvc := costs.New(client)

		from := time.Date(int(2022), time.May, int(22), int(0), int(0), int(0), int(0), time.UTC)
		to := time.Date(int(2022), time.June, int(21), int(0), int(0), int(0), int(0), time.UTC)
		_, err := costsSvc.GetProjectCosts(
			context.Background(),
			"ba04b091-ec32-4423-81e6-976f1b8a8363",
			from,
			to,
			consts.COSTS_GRANULARITY_DAILY,
			consts.COSTS_DEPTH_PROJECT,
		)

		if err == nil {
			t.Errorf("wanted error, got nil")
		}
	})
}
