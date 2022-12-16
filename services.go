package client

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.4/generated"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
)

type services struct {
	Kubernetes   *kubernetes.ClientWithResponses
	PostgresFlex *postgresflex.ClientWithResponses
}

func (c *Client) initServices() *Client {
	c.Services.Kubernetes, _ = kubernetes.NewClientWithResponses(consts.DEFAULT_BASE_URL, kubernetes.WithHTTPClient(c))
	c.Services.PostgresFlex, _ = postgresflex.NewClientWithResponses(consts.DEFAULT_BASE_URL, postgresflex.WithHTTPClient(c))
	return c
}
