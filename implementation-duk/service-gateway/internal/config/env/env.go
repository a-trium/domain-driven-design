package env

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config/db"
	"github.com/kelseyhightower/envconfig"
)

var (
	// These fields are populated by govvv
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
)

type Property struct {
	Mode             string `default:"LOCAL"` // LOCAL TEST DEV PROD
	ServiceName      string `default:"service-gateway"`
	ServiceId        string `default:"0"`
	Port             string `default:"9000"`
	DatabaseProperty db.Property

	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
}

var env Property

func init() {
	err := envconfig.Process("", &env)
	if err != nil {
		panic("Failed to get specification")
	}

	env.BuildDate = BuildDate
	env.GitCommit = GitCommit
	env.GitBranch = GitBranch
	env.GitState = GitState
	env.GitState = GitState
	env.Version = Version
}

//func (env *Environment) IsProdMode() bool {
//	return env.Mode == "PROD"
//}

func IsProdMode() bool {
	return env.Mode == "PROD"
}

func GetEnvironment() *Property {
	return &env
}
