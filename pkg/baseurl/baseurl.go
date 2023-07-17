package baseurl

import (
	"fmt"
	"os"
	"strings"
)

type BaseURL struct {
	// Base URL
	BaseURL string

	// OverrideWith specifies an environment
	// variable name. When set, the value
	// it contains will override the base URL
	OverrideWith string
}

// New expects the package name and base URL
// for example, for pkg=costs, OverrideWith will be
// STACKIT_COSTS_BASEURL
func New(pkg, baseURL string) BaseURL {
	return BaseURL{
		BaseURL:      baseURL,
		OverrideWith: fmt.Sprintf("STACKIT_%s_BASEURL", strings.ToUpper(pkg)),
	}
}

// Get returns the base URL
// if the override environment variable is set, it is returned instead
func (eu BaseURL) Get() string {
	url := os.Getenv(eu.OverrideWith)
	if url != "" {
		return url
	}
	return eu.BaseURL
}

// GetOverrideName returns the name of the environment variable
// that can be used to override the base URL
func (eu BaseURL) GetOverrideName() string {
	return eu.OverrideWith
}
