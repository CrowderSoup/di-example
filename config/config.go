package config

import "github.com/koding/multiconfig"

// Config our app config
type Config struct {
	Address     string `default:":3000"`
	Environment string `default:"development"`
}

// LoadConfig loads and returns our app config
func LoadConfig() *Config {
	var config Config
	m := multiconfig.New()
	m.MustLoad(&config)

	return &config
}
