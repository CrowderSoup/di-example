package logging

import (
	"cloud.bsdrocker.com/CrowderSoup/di-example/internal/config"
	"go.uber.org/fx"
)

// Logger application logger
type Logger struct {
	Config *config.Config
}

// NewLogger returns new application logger
func NewLogger(cfg *config.Config) *Logger {
	return &Logger{
		Config: cfg,
	}
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewLogger),
)
