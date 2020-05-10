package main

import (
	"cloud.bsdrocker.com/CrowderSoup/di-example/internal/config"
	"cloud.bsdrocker.com/CrowderSoup/di-example/internal/logging"
	"cloud.bsdrocker.com/CrowderSoup/di-example/internal/server"

	"go.uber.org/fx"
)

func main() {
	bundle := fx.Options(
		config.Module,
		logging.Module,
		server.Module,
	)
	app := fx.New(
		bundle,
	)

	app.Run()
}
