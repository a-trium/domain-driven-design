package main

import (
	"fmt"

	"github.com/a-trium/domain-driven-design/implementation-torres/service-gateway/internal/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/a-trium/domain-driven-design/implementation-torres/service-gateway/internal/pkg/config"
	_ "github.com/a-trium/domain-driven-design/implementation-torres/service-gateway/internal/domain/user"
)

func main() {

	// env
	env := config.GetEnvironment()
	fmt.Println(" #-- env : ", env)

	// Database
	db :=config.GetDatabase()
	//var repo user.UserRepository
	repo := user.NewUserRepository(db)
	repo.FineAll()


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
