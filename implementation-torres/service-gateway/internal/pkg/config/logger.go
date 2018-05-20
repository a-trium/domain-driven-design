package config

import "go.uber.org/zap"

func GetLogger() *zap.SugaredLogger {
	var log *zap.Logger = nil

	if env.IsTestMode() || env.IsLocalMode() {
		log, _ = zap.NewDevelopment()
	} else {
		log, _ = zap.NewProduction()
	}

	logger := log.Sugar().With("service_name", env.ServiceName, "service_id", env.ServiceId, )

	return logger
}
