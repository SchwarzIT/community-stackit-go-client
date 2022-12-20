package client

import (
	kubernetes "github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/generated"
	postgresflex "github.com/SchwarzIT/community-stackit-go-client/pkg/services/postgres-flex/v1.0/generated"
)

type services struct {
	Kubernetes   *kubernetes.ClientWithResponses
	PostgresFlex *postgresflex.ClientWithResponses
}

func (c *Client) initServices() error {
	var err error

	c.Services.Kubernetes, err = kubernetes.NewService(c)
	if err != nil {
		return err
	}

	c.Services.PostgresFlex, err = postgresflex.NewService(c)
	if err != nil {
		return err
	}

	return nil
}
