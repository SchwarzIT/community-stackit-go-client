// This file is used to validate data in the jobs package

package jobs

import (
	"fmt"
	"regexp"
	"time"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
	"github.com/pkg/errors"
)

// Validate validates a job
func (job Job) Validate() error {
	if err := ValidateJobName(job.JobName); err != nil {
		return err
	}
	if err := ValidateScheme(job.Scheme); err != nil {
		return err
	}
	if err := ValidateDurations(job.ScrapeInterval, job.ScrapeTimeout); err != nil {
		return err
	}
	if err := ValidateStaticConfigs(job.StaticConfigs); err != nil {
		return err
	}
	if err := ValidateMetricsPath(job.MetricsPath); err != nil {
		return err
	}
	return nil
}

// ValidateJobName validates the job name
func ValidateJobName(name string) error {
	if name == "" {
		return errors.New("job name is required")
	}
	exp := `^[a-zA-Z0-9-]{1,200}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid job name.\n- valid name is of: '%s'\n- '%s' was given", exp, name)
	}
	return nil
}

// ValidateScheme validates the protocol
func ValidateScheme(protocol string) error {
	if protocol == "" {
		return errors.New("job scheme is a required field")
	}
	if protocol != "http" && protocol != "https" {
		return errors.New("protocol must be one of 'http' or 'https'")
	}
	return nil
}

// ValidateDurations validates the scraping interval and timeout interval
func ValidateDurations(intervalDuration, timeoutDuration string) error {
	if intervalDuration == "" {
		return errors.New("job scrape interval is a required field")
	}
	d, err := validate.Duration(intervalDuration)
	if err != nil {
		return errors.Wrap(err, "scrape interval")
	}
	if d < (60 * time.Second) {
		return errors.New("scrape interval must be >= 60s")
	}
	if timeoutDuration == "" {
		return errors.New("job scrape timeout is a required field")
	}
	t, err := validate.Duration(timeoutDuration)
	if err != nil {
		return errors.Wrap(err, "scrape timeout")
	}
	if t > d {
		return errors.New("scrape timeout must be smaller than scrape interval")
	}
	return nil
}

// ValidateStaticConfigs validates provided static configs
func ValidateStaticConfigs(sc []StaticConfig) error {
	if len(sc) == 0 {
		return errors.New("static configs must be specified")
	}
	for _, v := range sc {
		if v.Targets == nil || len(v.Targets) == 0 {
			return errors.New("targets must be specified for every item in static configs")
		}
	}
	return nil
}

// ValidateMetricsPath validates the metrics path
func ValidateMetricsPath(mp string) error {
	if mp == "" {
		return errors.New("Metrics Path must be specified")
	}
	if len(mp) > 200 {
		return errors.New("Metrics Path length must be 1..200")
	}
	return nil
}
