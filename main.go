package main

import (
	"cloud.bsdrocker.com/CrowderSoup/di-example/config"
	"cloud.bsdrocker.com/CrowderSoup/di-example/logger"
	"cloud.bsdrocker.com/CrowderSoup/di-example/server"
)

func main() {
	config := config.LoadConfig()

	logger, _ := logger.NewZapLogger(config)
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	server := server.NewServer(config, sugar)

	sugar.Fatal(server.Run())
}
