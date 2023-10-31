package log

import (
	"log"

	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

func Init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any

	sugarLogger = logger.Sugar()
	sugarLogger.Debug("log package - init")
}

func GetLogger() *zap.SugaredLogger {
	sugarLogger.Debug("log package - GetLogger")
	return sugarLogger
}

func GetDesugaredLogger() *zap.Logger {
	sugarLogger.Debug("log package - GetDesugaredLogger")
	return sugarLogger.Desugar()
}
