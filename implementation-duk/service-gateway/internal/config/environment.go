package config

import (
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/config/db"
	"github.com/kelseyhightower/envconfig"
	"strings"
)

var (
	// These fields are populated by govvv
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
)

type Environment struct {
	Debug            bool   `default:"true"`
	Mode             string `default:"LOCAL"` // LOCAL TEST DEV PROD
	ServiceName      string `default:"service-gateway"`
	ServiceId        string `default:"0"`
	Port             string `default:"9000"`
	DatabaseProperty db.DatabaseProperty

	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
}

var env Environment

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

func (env *Environment) IsProd() bool {
	return strings.EqualFold(env.Mode, "PROD")
}

func (env *Environment) IsDev() bool {
	return strings.EqualFold(env.Mode, "DEV")
}

func (env *Environment) IsLocal() bool {
	return strings.EqualFold(env.Mode, "LOCAL")
}

func (env *Environment) IsTest() bool {
	return strings.EqualFold(env.Mode, "TEST")
}

func (e *Environment) IsDebugging() bool {
	return e.Debug
}

func GetEnvironment() *Environment {
	return &env
}
