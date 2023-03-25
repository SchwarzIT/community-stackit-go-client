package env

import (
	"fmt"
	"os"
	"strings"
)

type EnvironmentURLs struct {
	// Base URL for each environment
	Prod string
	QA   string
	Dev  string

	// OverrideWith specifies an environment
	// variable name. When set, the value
	// it contains will override the base URL
	OverrideWith string
}

// URLs expects the package name and base URLs for prod, qa, dev
// the package name is used fot setting OverrideWith
// for example, for pkg=costs, OverrideWith will be
// STACKIT_COSTS_BASEURL
func URLs(pkg, prod, qa, dev string) EnvironmentURLs {
	return EnvironmentURLs{
		Prod:         prod,
		QA:           qa,
		Dev:          dev,
		OverrideWith: fmt.Sprintf("STACKIT_%s_BASEURL", strings.ToUpper(pkg)),
	}
}

func (eu EnvironmentURLs) GetURL(e Environment) string {
	url := os.Getenv(eu.OverrideWith)
	if url != "" {
		return url
	}
	if e.IsDev() {
		return eu.Dev
	}
	if e.IsQA() {
		return eu.QA
	}
	return eu.Prod
}
