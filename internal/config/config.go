package config

import (
	"go.uber.org/fx"
)

// Config our applications config
type Config struct {
	HTTPAddress string
}

// NewConfig provides our application config
func NewConfig() *Config {
	return &Config{
		HTTPAddress: ":3000",
	}
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewConfig),
)
