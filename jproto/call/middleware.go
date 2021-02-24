package call

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

//获取日志
func GetZapLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Error("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	//grpcZap.ReplaceGrpcLoggerV2(logger)
	return logger
}
