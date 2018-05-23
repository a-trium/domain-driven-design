package config
//
//import "go.uber.org/zap"
//
//func GetLogger(env Environment) *zap.SugaredLogger {
//
//	var log *zap.Logger = nil
//
//	if IsProdMode() {
//		log, _ = zap.NewProduction()
//	} else {
//		log, _ = zap.NewDevelopment()
//	}
//
//	logger := log.Sugar().With("service_name", env.ServiceName, "service_id", env.ServiceId, )
//
//	return logger
//}
