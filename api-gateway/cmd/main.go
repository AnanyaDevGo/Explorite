package main

import (
	logging "ExploriteGateway/Logging"
	_ "ExploriteGateway/cmd/docs"
	"ExploriteGateway/pkg/config"
	"ExploriteGateway/pkg/di"
	"log"
)

// @title Go + Gin social media API Explorite
// @version 1.0.0
// @description Explorite is a social media platform
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
// @query.collection.format multi

func main() {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/explorite_gateway.log")
	defer logrusLogFile.Close()

	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		logrusLogger.Error("Failed to load config: ", cfgErr)
		log.Fatal("cannot load config: ", cfgErr)
	}
	server, diErr := di.InitializeAPI(cfg)
	if diErr != nil {
		logrusLogger.Fatal("Cannot start server: ", diErr)
		log.Fatal("cannot start server: ", diErr)
	}
	logrusLogger.Info("explorite started")
	server.Start()
}
