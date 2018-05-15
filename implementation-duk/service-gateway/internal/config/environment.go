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

type DatabaseProperty struct {
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	UserName string `default:"root"`
	Password string `default:"root"`
	Database string `default:"application"`
}

type Environment struct {
	Mode       string `default:"LOCAL"` // LOCAL TEST DEV PROD
	ServiceNam string `default:"service-gateway"`
	Database   DatabaseProperty

	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
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
