package main

import (
	"context"
	"os"

	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/config"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/jessevdk/go-flags"
	"github.com/rs/cors"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/domain/user"
	"github.com/gorilla/sessions"
	)

func main() {
	env := config.Env
	logger := config.GetLogger().With("service_name", env.ServiceName, "service_id", env.ServiceId, )

	logger.Infow("Build Manifest",
		"build_date", env.BuildDate,
		"git_commit", env.GitCommit,
		"git_branch", env.GitBranch,
		"git_state", env.GitState,
		"version", env.Version,
		"host", env.Host,
		"port", env.RestPort,
	)

	swaggerSpec, err := loads.Analyzed(swagserver.FlatSwaggerJSON, "")
	if err != nil {
		logger.Fatalw("Failed to configure REST server", "error", err)
	}
	api := swagapi.NewGatewayAPI(swaggerSpec)
	server := swagserver.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			logger.Fatalw("Failed to parse command-line option for REST server", "error", err)
		}
	}

	server.Host = env.Host
	server.Port = env.RestPort
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.Logger = logger.Infof

	// configure session storage
	sessionSecret := "something-very-secret"
	sessionStore := sessions.NewCookieStore([]byte(sessionSecret))

	// configure database
	db := config.GetDatabase()

	// configure REST server handlers, middlewares
	logger.Info("Configure REST API handlers")

	userRepo := user.NewRepository(db)
	encryptor := user.NewEncryptor(0)
	authHandler := user.NewAuthHandler(userRepo, encryptor, sessionStore)

	authHandler.Configure(api)

	logger.Info("Configure REST API middleware")

	handler := api.Serve(nil)
	handler = cors.AllowAll().Handler(handler)
	server.SetHandler(handler)

	_, cancel := context.WithCancel(context.Background())

	api.ServerShutdown = func() {
		cancel()
	}

	if err := server.Serve(); err != nil {
		logger.Fatalw("Failed to start REST server", "error", err)
	}
}
