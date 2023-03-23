// this file is used for validating cluster data and properties

package cluster

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/kubernetes/v1.0/cluster"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

// Validate validates the given cluster data (dry validation)
func Validate(
	clusterName string,
	clusterConfig cluster.Kubernetes,
	nodePools []cluster.Nodepool,
	maintenance *cluster.Maintenance,
	hibernation *cluster.Hibernation,
	extensions *cluster.Extension,
) error {
	if err := validate.SemVer(clusterConfig.Version); err != nil {
		return err
	}
	if err := ValidateClusterName(clusterName); err != nil {
		return err
	}
	if len(nodePools) == 0 {
		return errors.New("at least one node pool must be specified")
	}
	for _, np := range nodePools {
		if err := ValidateNodePool(np); err != nil {
			return err
		}
	}
	if err := ValidateMaintenance(maintenance); err != nil {
		return err
	}
	if err := ValidateHibernation(hibernation); err != nil {
		return err
	}
	if err := ValidateExtensions(extensions); err != nil {
		return err
	}
	return nil
}

// ValidateClusterName validates a given cluster name
func ValidateClusterName(name string) error {
	exp := `^[a-z0-9]{1}[a-z0-9-]{0,10}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid cluster name. valid name is of: %s", exp)
	}
	return nil
}

// ValidateNodePoolName validates a given pool name
func ValidateNodePoolName(name string) error {
	exp := `^[a-z0-9]{1}[a-z0-9-]{0,14}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid node pool name. valid name is of: %s", exp)
	}
	return nil
}

// ValidateNodePool validates a given node pool
func ValidateNodePool(np cluster.Nodepool) error {
	if err := ValidateNodePoolName(np.Name); err != nil {
		return err
	}
	if np.Machine.Type == "" {
		return errors.New("machine type must be specified")
	}
	if np.Machine.Image.Version == "" {
		return errors.New("machine image version must be specified")
	}
	if np.Minimum > np.Maximum {
		return errors.New("minimum value can't be larger than maximum")
	}
	if np.Minimum < 1 || np.Minimum > 100 {
		return errors.New("minimum value must be in the range of 1..100")
	}
	if np.Maximum < 1 || np.Maximum > 100 {
		return errors.New("maximum value must be in the range of 1..100")
	}
	if np.MaxSurge != nil {
		if *np.MaxSurge < 1 || *np.MaxSurge > 10 {
			return errors.New("max surge value must be in the range of 1..10")
		}
	}
	if np.Volume.Size < 20 || np.Volume.Size > 10240 {
		return errors.New("volume size value must be in the range of 20..10240")
	}
	if np.Taints != nil {
		for _, t := range *np.Taints {
			if err := ValidateTaint(t); err != nil {
				return err
			}
		}
	}
	if err := ValidateCRI(np.CRI); err != nil {
		return err
	}
	return nil
}

// ValidateTaint validates a given node pool taint
func ValidateTaint(t cluster.Taint) error {
	switch t.Effect {
	case cluster.NO_EXECUTE:
		fallthrough
	case cluster.NO_SCHEDULE:
		fallthrough
	case cluster.PREFER_NO_SCHEDULE:
	default:
		return fmt.Errorf("invalid taint effect '%s'", t.Effect)
	}

	if t.Key == "" {
		return errors.New("taint key is required")
	}
	return nil
}

// ValidateCRI validates the given cri struct
func ValidateCRI(c *cluster.CRI) error {
	if c == nil {
		return nil
	}
	if c.Name == nil {
		return nil
	}
	switch *c.Name {
	case cluster.CONTAINERD:
		fallthrough
	case cluster.DOCKER:
	default:
		return fmt.Errorf("invalid CRI name '%s'", string(*c.Name))
	}
	return nil
}

// ValidateMaintenance validates a given cluster maintenance
func ValidateMaintenance(m *cluster.Maintenance) error {
	if m == nil {
		return nil
	}
	if m.TimeWindow.End == "" {
		return errors.New("maintenance end time window is required")
	}
	if m.TimeWindow.Start == "" {
		return errors.New("maintenance start time window is required")
	}
	return nil
}

// ValidateHibernation validates a given cluster hibernation
func ValidateHibernation(h *cluster.Hibernation) error {
	if h == nil {
		return nil
	}
	for _, s := range h.Schedules {
		if s.End == "" {
			return errors.New("hibernation end time is required")
		}
		if s.Start == "" {
			return errors.New("hibernation start time is required")
		}
	}
	return nil
}

// ValidateExtensions validates a given cluster extensions
func ValidateExtensions(e *cluster.Extension) error {
	if e == nil {
		return nil
	}
	if e.Argus != nil {
		if e.Argus.Enabled && e.Argus.ArgusInstanceID == "" {
			return errors.New("argus instance ID is mandatory when Argus is enabled")
		}
	}
	return nil
}
