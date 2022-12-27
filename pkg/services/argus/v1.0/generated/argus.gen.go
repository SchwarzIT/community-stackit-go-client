// Package argus provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/do87/oapi-codegen version v0.5.0 DO NOT EDIT.
package argus

import (
	"net/url"
	"strings"

	common "github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/acl"
	alertconfig "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/alert-config"
	alertgroups "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/alert-groups"
	alertrecords "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/alert-records"
	alertrules "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/alert-rules"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/backup"
	certcheck "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/cert-check"
	grafanaconfigs "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/grafana-configs"
	httpcheck "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/http-check"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/logs"
	metricsstorageretention "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/metrics-storage-retention"
	networkcheck "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/network-check"
	pingcheck "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/ping-check"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/plans"
	scrapeconfig "github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/scrape-config"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/argus/v1.0/generated/traces"
)

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// list of connected client services
	Instances               *instances.Client
	Acl                     *acl.Client
	AlertConfig             *alertconfig.Client
	AlertGroups             *alertgroups.Client
	AlertRules              *alertrules.Client
	AlertRecords            *alertrecords.Client
	Backup                  *backup.Client
	CertCheck               *certcheck.Client
	GrafanaConfigs          *grafanaconfigs.Client
	HttpCheck               *httpcheck.Client
	Logs                    *logs.Client
	MetricsStorageRetention *metricsstorageretention.Client
	NetworkCheck            *networkcheck.Client
	PingCheck               *pingcheck.Client
	ScrapeConfig            *scrapeconfig.Client
	Traces                  *traces.Client
	Plans                   *plans.Client

	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client common.Client
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a factory client
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}

	client.Instances = instances.NewClient(server, client.Client)
	client.Acl = acl.NewClient(server, client.Client)
	client.AlertConfig = alertconfig.NewClient(server, client.Client)
	client.AlertGroups = alertgroups.NewClient(server, client.Client)
	client.AlertRules = alertrules.NewClient(server, client.Client)
	client.AlertRecords = alertrecords.NewClient(server, client.Client)
	client.Backup = backup.NewClient(server, client.Client)
	client.CertCheck = certcheck.NewClient(server, client.Client)
	client.GrafanaConfigs = grafanaconfigs.NewClient(server, client.Client)
	client.HttpCheck = httpcheck.NewClient(server, client.Client)
	client.Logs = logs.NewClient(server, client.Client)
	client.MetricsStorageRetention = metricsstorageretention.NewClient(server, client.Client)
	client.NetworkCheck = networkcheck.NewClient(server, client.Client)
	client.PingCheck = pingcheck.NewClient(server, client.Client)
	client.ScrapeConfig = scrapeconfig.NewClient(server, client.Client)
	client.Traces = traces.NewClient(server, client.Client)
	client.Plans = plans.NewClient(server, client.Client)

	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer common.Client) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	Client *Client

	// list of connected client services
	Instances               *instances.ClientWithResponses
	Acl                     *acl.ClientWithResponses
	AlertConfig             *alertconfig.ClientWithResponses
	AlertGroups             *alertgroups.ClientWithResponses
	AlertRules              *alertrules.ClientWithResponses
	AlertRecords            *alertrecords.ClientWithResponses
	Backup                  *backup.ClientWithResponses
	CertCheck               *certcheck.ClientWithResponses
	GrafanaConfigs          *grafanaconfigs.ClientWithResponses
	HttpCheck               *httpcheck.ClientWithResponses
	Logs                    *logs.ClientWithResponses
	MetricsStorageRetention *metricsstorageretention.ClientWithResponses
	NetworkCheck            *networkcheck.ClientWithResponses
	PingCheck               *pingcheck.ClientWithResponses
	ScrapeConfig            *scrapeconfig.ClientWithResponses
	Traces                  *traces.ClientWithResponses
	Plans                   *plans.ClientWithResponses
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}

	cwr := &ClientWithResponses{Client: client}
	cwr.Instances = instances.NewClientWithResponses(server, client.Client)
	cwr.Acl = acl.NewClientWithResponses(server, client.Client)
	cwr.AlertConfig = alertconfig.NewClientWithResponses(server, client.Client)
	cwr.AlertGroups = alertgroups.NewClientWithResponses(server, client.Client)
	cwr.AlertRules = alertrules.NewClientWithResponses(server, client.Client)
	cwr.AlertRecords = alertrecords.NewClientWithResponses(server, client.Client)
	cwr.Backup = backup.NewClientWithResponses(server, client.Client)
	cwr.CertCheck = certcheck.NewClientWithResponses(server, client.Client)
	cwr.GrafanaConfigs = grafanaconfigs.NewClientWithResponses(server, client.Client)
	cwr.HttpCheck = httpcheck.NewClientWithResponses(server, client.Client)
	cwr.Logs = logs.NewClientWithResponses(server, client.Client)
	cwr.MetricsStorageRetention = metricsstorageretention.NewClientWithResponses(server, client.Client)
	cwr.NetworkCheck = networkcheck.NewClientWithResponses(server, client.Client)
	cwr.PingCheck = pingcheck.NewClientWithResponses(server, client.Client)
	cwr.ScrapeConfig = scrapeconfig.NewClientWithResponses(server, client.Client)
	cwr.Traces = traces.NewClientWithResponses(server, client.Client)
	cwr.Plans = plans.NewClientWithResponses(server, client.Client)

	return cwr, nil
}
