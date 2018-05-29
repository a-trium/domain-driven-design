package config

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func GetLogger(env *Environment) *Logger {
	var log *zap.Logger = nil

	if env.IsProd() {
		log, _ = zap.NewProduction()
	} else {
		log, _ = zap.NewDevelopment()
	}

	return &Logger{log.Sugar().With("service_name", env.ServiceName, "service_id", env.ServiceId, )}
}
