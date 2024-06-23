package log

import "go.uber.org/zap"
import log2 "log"

var Logger *zap.Logger

func InitLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log2.Println(err)
		}
	}(logger)
	Logger = logger
	return nil
}
