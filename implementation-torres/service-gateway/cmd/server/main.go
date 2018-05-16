package main

import (
	"fmt"

	"github.com/a-trium/domain-driven-design/implementation-torres/service-gateway/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {

	// env
	env := config.Env
	fmt.Println(" #-- env : ", env)

	// zap - logger
	fmt.Println(" #-- logger")
	logger := config.GetLogger().With("service_name", env.ServiceName, "service_id", env.ServiceId, )

	logger.Infow("Build Manifest",
		"build_date", env.BuildDate,
		"git_commit", env.GitCommit,
		"git_branch", env.GitBranch,
		"git_state", env.GitState,
		"version", env.Version,
	)

	// errors
	err := errors.New("Error Msg")
	fmt.Println(" #-- errors : ", err)

	// gin
	fmt.Println(" #-- starting Gin REST API")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}