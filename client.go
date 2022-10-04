// package client groups together all the services that the STACKIT client supports
// Active services are found under `ProductiveServices` whereas new services or services
// that still need to be further developed or tested, can be put under `Incubator`
// All services must be initialized in the `init` method, and the client must be configured
// during initialization

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/costs"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/kubernetes"
	objectstorage "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/object-storage"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/projects"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/roles"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/users"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/membership"
	resourceManager "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v2/resource-manager"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/retry"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// Client service for managing interactions with STACKIT API
type Client struct {
	client *http.Client
	config *Config
	retry  *retry.Retry

	// Productive services - services that are ready to be used in production
	ProductiveServices

	// Incubator - services under development or currently being tested
	// not ready for production usage
	Incubator IncubatorServices
}

// New returns a new client
func New(ctx context.Context, cfg *Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	c := &Client{config: cfg}
	return c.init(ctx), nil
}

// WithRetry sets retry.Retry in a shallow copy of the given client
// and returns the new copy
func (c *Client) WithRetry(r *retry.Retry) *Client {
	nc := *c
	nc.retry = r
	return &nc
}

// Service management

// ProductiveServices is the struct representing all productive services
type ProductiveServices struct {
	Argus         *argus.ArgusService
	Costs         *costs.CostsService
	Kubernetes    *kubernetes.KubernetesService
	ObjectStorage *objectstorage.ObjectStorageService
	Projects      *projects.ProjectService
	Roles         *roles.RolesService
	Users         *users.UsersService
}

// IncubatorServices is the struct representing all services that are under development
type IncubatorServices struct {
	Membership      *membership.MembershipService
	ResourceManager *resourceManager.ResourceManagerService
}

// init initializes the client and its services and returns the client
func (c *Client) init(ctx context.Context) *Client {
	c.setHttpClient(ctx)

	// init productive services
	c.Argus = argus.New(c)
	c.Costs = costs.New(c)
	c.Kubernetes = kubernetes.New(c)
	c.ObjectStorage = objectstorage.New(c)
	c.Projects = projects.New(c)
	c.Roles = roles.New(c)
	c.Users = users.New(c)

	// init incubator services
	c.Incubator = IncubatorServices{
		Membership:      membership.New(c),
		ResourceManager: resourceManager.New(c),
	}
	return c
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.client
}

func (c *Client) GetConfig() *Config {
	return c.config
}

func (c *Client) SetToken(token string) {
	c.config.Token = token
}

// setHttpClient creates the client's oauth client
func (c *Client) setHttpClient(ctx context.Context) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.config.Token},
	)
	hcl := oauth2.NewClient(ctx, ts)
	hcl.Timeout = time.Second * 10
	c.client = hcl
}

// Request creates a new API request
func (c *Client) Request(ctx context.Context, method, path string, body []byte) (*http.Request, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}
	pathSplit := strings.Split(path, "?")
	rel := &url.URL{Path: pathSplit[0]}
	if len(pathSplit) == 2 {
		rel.RawQuery = pathSplit[1]
	}
	u := c.config.BaseUrl.ResolveReference(rel)
	payload := strings.NewReader(string(body))
	req, err := http.NewRequestWithContext(ctx, method, u.String(), payload)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Do is a wrapper for do() which also manages retry if set
func (c *Client) Do(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error) {
	if c.retry == nil {
		return c.do(req, v, errorHandlers...)
	}
	return c.doWithRetry(req, v, errorHandlers...)
}

func (c *Client) doWithRetry(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error) {
	var lastErr error
	var res *http.Response
	var retry bool

	overall, cancel := context.WithTimeout(req.Context(), c.retry.Timeout)
	defer cancel()

	// set overall ctx in request
	req = req.WithContext(overall)

	// clone max retries
	maxRetries := c.retry.MaxRetries
	retries := 0
	if maxRetries != nil {
		retries = *maxRetries
	}

	for {
		res, retries, retry, lastErr = c.doTick(req, v, errorHandlers, maxRetries != nil, retries)
		if lastErr == nil || !retry {
			return res, lastErr
		}

		tick, cancelTick := context.WithTimeout(context.Background(), c.retry.Throttle)
		defer cancelTick()

		select {
		case <-tick.Done():
			// continue
		case <-overall.Done():
			return nil, errors.Wrap(lastErr, "retry context timed out")
		}
	}
}

func (c *Client) doTick(req *http.Request, v interface{}, errorHandlers []func(*http.Response) error, withMaxRetries bool, retries int) (res *http.Response, retriesLeft int, retry bool, err error) {
	retry = true

	res, err = c.do(req, v, errorHandlers...)
	if err == nil {
		retry = false
		return
	}

	// check if error is retryable
	for _, f := range c.retry.IsRetryableFns {
		if !f(err) {
			retry = false
			return
		}
	}

	retriesLeft = retries
	if withMaxRetries {
		if retries <= 0 {
			err = errors.Wrap(err, "reached max retries")
			retry = false
			return
		}
		retriesLeft = retriesLeft - 1
	}

	return
}

// Do performs the request and decodes the response if given interface != nil
func (c *Client) do(req *http.Request, v interface{}, errorHandlers ...func(*http.Response) error) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle errors in the response
	if len(errorHandlers) == 0 {
		errorHandlers = append(errorHandlers, validate.DefaultResponseErrorHandler)
	}
	for _, fn := range errorHandlers {
		if err := fn(resp); err != nil {
			return resp, err
		}
	}

	// parse response JSON
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

// OrganizationID returns the organization ID defined in the configuration
func (c *Client) OrganizationID() string {
	return c.config.OrganizationID
}

// MockServer mocks STACKIT api server
// and returns a client pointing to it, mux, teardown function and an error indicator
func MockServer() (c *Client, mux *http.ServeMux, teardown func(), err error) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	teardown = server.Close

	u, _ := url.Parse(server.URL)

	c, err = New(context.Background(), &Config{
		BaseUrl:          u,
		Token:            "token",
		ServiceAccountID: "sa-id",
		OrganizationID:   consts.SCHWARZ_ORGANIZATION_ID,
	})

	return
}
