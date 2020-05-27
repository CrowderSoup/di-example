package logger

import (
	"cloud.bsdrocker.com/CrowderSoup/di-example/config"
	"go.uber.org/zap"
)

// NewZapLogger returns a new zap logger
func NewZapLogger(config *config.Config) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	if config.Environment == "development" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	return logger, err
}
