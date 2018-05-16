package config

import "github.com/kelseyhightower/envconfig"

var (
	// These fields are populated by govvv
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
)

type Environment struct {
	Mode             string `default:"LOCAL"` // LOCAL TEST DEV PROD
	ServiceName      string `default:"service-gateway"`
	ServiceId        string `default:"0"`
	DatabaseProperty DatabaseProperty

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

//func (env *Environment) IsProdMode() bool {
//	return env.Mode == "PROD"
//}

func IsProdMode() bool {
	return env.Mode == "PROD"
}

func GetEnvironment() *Environment {
	return &env
}
