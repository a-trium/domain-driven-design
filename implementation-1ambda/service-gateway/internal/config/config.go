package config

import (
	"github.com/kelseyhightower/envconfig"
)

const RestPort = 30001

var (
	// These fields are populated by govvv
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
)

type Environment struct {
	debug bool   `default:"true"`
	mode  string `default:"LOCAL"` // `LOCAL`, `TEST`, `DEV`, `PROD`

	RestPort      int    `default:"30001"`
	CorsAllowUrl  string `default:"localhost:8080"`
	ServiceName   string `default:"service-gateway"`
	ServiceId     string `default:"0"`
	MysqlHost     string `default:"localhost"`
	MysqlPort     string `default:"33306"`
	MysqlUserName string `default:"root"`
	MysqlPassword string `default:"root"`
	MysqlDatabase string `default:"ddd"`

	// copied from govvv injected values
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
}

func (e *Environment) DebugEnabled() bool {
	return e.debug
}

func (e *Environment) IsTestMode() bool {
	return e.mode == "TEST"
}

func (e *Environment) IsLocalMode() bool {
	return e.mode == "LOCAL"
}

var Env Environment

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic("Failed to get specification")
	}

	Env.BuildDate = BuildDate
	Env.GitCommit = GitCommit
	Env.GitBranch = GitBranch
	Env.GitState = GitState
	Env.GitState = GitState
	Env.Version = Version
}
