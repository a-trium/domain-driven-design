package config

import (
	"go.uber.org/zap"
)

func GetLogger(env *Environment) *zap.SugaredLogger {
	var log *zap.Logger = nil

	if env.IsProd() {
		log, _ = zap.NewProduction()
	} else {
		log, _ = zap.NewDevelopment()
	}

	return log.Sugar().With("service_name", env.ServiceName, "service_id", env.ServiceId, )
}