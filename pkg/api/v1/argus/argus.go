package argus

import (
	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/credentials"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/grafana"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/instances"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/jobs"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/metrics"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/plans"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/argus/traces"
)

// New returns a new handler for the service
func New(c common.Client) *ArgusService {
	return &ArgusService{
		Credentials: &credentials.CredentialsService{Client: c},
		Grafana:     &grafana.GrafanaService{Client: c},
		Instances:   &instances.InstancesService{Client: c},
		Jobs:        &jobs.JobsService{Client: c},
		Plans:       plans.New(c),
		Metrics:     &metrics.MetricsService{Client: c},
		Traces:      &traces.TracesService{Client: c},
	}
}

// ArgusService is the service that handles
// Monitoring & Logging related services managed by Stackit's Argus
type ArgusService struct {
	Credentials *credentials.CredentialsService
	Grafana     *grafana.GrafanaService
	Instances   *instances.InstancesService
	Jobs        *jobs.JobsService
	Plans       *plans.PlansService
	Metrics     *metrics.MetricsService
	Traces      *traces.TracesService
}
