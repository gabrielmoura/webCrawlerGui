package log

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()
	Logger = logger
	return nil
}
