package config

import "go.uber.org/zap"

func GetLogger() *zap.SugaredLogger {

	var log *zap.Logger = nil

	if Env.IsProdMode() {
		log, _ = zap.NewProduction()
	} else {
		log, _ = zap.NewDevelopment()
	}

	logger := log.Sugar().With("service_name", Env.ServiceName, "service_id", Env.ServiceId, )

	return logger
}
