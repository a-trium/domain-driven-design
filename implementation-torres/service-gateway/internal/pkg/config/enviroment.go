package config

import (
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

type Environment struct {
   Debug bool   `default:"true"`
   Mode  string `default:"LOCAL"` // `LOCAL`, `TEST`, `DEV`, `PROD`

   RestPort      int    `default:"30001"`
   CorsAllowUrl  string `default:"localhost:8080"`
   ServiceName   string `default:"service-gateway"`
   ServiceId     string `default:"0"`
   MysqlHost     string `default:"localhost"`
   MysqlPort     string `default:"3306"`
   MysqlUserName string `default:"root"`
   MysqlPassword string `default:"root"`
   MysqlDatabase string `default:"application"`

   // copied from govvv injected values
   BuildDate string
   GitCommit string
   GitBranch string
   GitState  string
   Version   string
}

func (e *Environment) DebugEnabled() bool {
   return e.Debug
}

func (e *Environment) IsTestMode() bool {
   return e.Mode == "TEST"
}

func (e *Environment) IsLocalMode() bool {
   return e.Mode == "LOCAL"
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

func GetEnvironment() *Environment {
   return &env
}