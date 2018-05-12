package config

import (
   "time"
)


var (
   // These fields are populated by govvv
   BuildDate  string
   GitCommit  string
   GitBranch  string
   GitState   string
   GitSummary string
   Version    string
   Started   = time.Now().UTC().Format(time.RFC3339)
)

//var (
// Version   = "undefined"
// BuildTime = "undefined"
// GitHash   = "undefined"
// Started   = time.Now().UTC().Format(time.RFC3339)
//)

type Flag struct {
   BuildDate  string
   GitCommit  string
   GitBranch  string
   GitState   string
   GitSummary string
   Version    string
   Started   string
}

func GetFlag() Flag {
   return Flag{
      Version:   Version,
      BuildDate: BuildDate,
      GitCommit: GitCommit,
      Started:   Started,
   }
}
