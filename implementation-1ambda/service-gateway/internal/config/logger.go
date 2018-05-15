package config

import "go.uber.org/zap"

func GetLogger() *zap.SugaredLogger {
	var log *zap.Logger = nil

	if Env.IsTestMode() || Env.IsLocalMode() {
		log, _ = zap.NewDevelopment()
	} else {
		log, _ = zap.NewProduction()
	}

	logger := log.Sugar().With("service_name", Env.ServiceName, "service_id", Env.ServiceId, )

	return logger
}
