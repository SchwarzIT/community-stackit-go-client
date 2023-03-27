package env

import "strings"

type Environment string

// Define API environments
const (
	Production Environment = "prod"
	QA         Environment = "qa"
	Dev        Environment = "dev"
)

func Parse(env string) Environment {
	if strings.EqualFold(env, "dev") {
		return Dev
	}
	if strings.EqualFold(env, "qa") {
		return QA
	}
	return Production
}

func (e Environment) IsProduction() bool {
	return e == Production
}

func (e Environment) IsQA() bool {
	return e == QA
}

func (e Environment) IsDev() bool {
	return e == Dev
}
