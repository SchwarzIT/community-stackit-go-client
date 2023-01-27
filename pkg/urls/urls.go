// urls is used to manage base urls
// for every STACKIT environment
package urls

import (
	"fmt"
	"os"
	"strings"

	"github.com/SchwarzIT/community-stackit-go-client/internal/common"
)

type ByEnvs struct {
	// Base URL for each environment
	Prod string
	QA   string
	Dev  string

	// OverrideWith specifies an environment
	// variable name. When set, the value
	// it contains will override the base URL
	OverrideWith string
}

// Init expects base URL strings for pkg, prod, qa, dev
// the package name is used fot setting OverrideWith
// for example, for pkg=costs, OverrideWith will be
// STACKIT_COSTS_BASEURL
func Init(pkg, prod, qa, dev string) ByEnvs {
	return ByEnvs{
		Prod:         prod,
		QA:           qa,
		Dev:          dev,
		OverrideWith: fmt.Sprintf("STACKIT_%s_BASEURL", strings.ToUpper(pkg)),
	}
}

func (e ByEnvs) GetURL(c common.Client) string {
	url := os.Getenv(e.OverrideWith)
	if url != "" {
		return url
	}

	switch c.GetEnvironment() {
	case common.ENV_DEV:
		return e.Dev
	case common.ENV_QA:
		return e.QA
	default:
		return e.Prod
	}
}
