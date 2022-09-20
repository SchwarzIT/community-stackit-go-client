// this file is used for functions that validate data in the Argus metrics package

package metrics

import (
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
)

// Validate validates the retention data
func (r Config) Validate() error {
	ret1hr, err := validate.Duration(r.MetricsRetentionTime1h)
	if err != nil {
		return errors.Wrap(err, "metrics retention time 1h")
	}

	ret5m, err := validate.Duration(r.MetricsRetentionTime5m)
	if err != nil {
		return errors.Wrap(err, "metrics retention time 5m")
	}

	retRaw, err := validate.Duration(r.MetricsRetentionTimeRaw)
	if err != nil {
		return errors.Wrap(err, "metrics retention time raw")
	}

	tm, _ := validate.Duration("13months")
	if retRaw > tm {
		return errors.New("retention time of longtime storage of raw sampled data must not be bigger than 13 months")
	}

	if ret5m > retRaw {
		return errors.New("retention time of longtime storage of 5m sampled data must not be bigger than metricsRetentionTimeRaw")
	}

	if ret1hr > ret5m {
		return errors.New("retention time of longtime storage of 1h sampled data must not be bigger than metricsRetentionTime5m")
	}

	return nil
}
