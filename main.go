package main

import (
	"cloud.bsdrocker.com/CrowderSoup/di-example/config"
	"cloud.bsdrocker.com/CrowderSoup/di-example/logger"
	"cloud.bsdrocker.com/CrowderSoup/di-example/server"
	"go.uber.org/fx"
)

func main() {
	bundle := fx.Options(
		config.Module,
		logger.Module,
		server.Module,
	)
	app := fx.New(
		bundle,
		fx.Invoke(logger.InvokeLogger),
		fx.Invoke(server.Run),
	)

	app.Run()

	<-app.Done()
}
