package main

import (
	"ExploriteGateway/pkg/config"
	"ExploriteGateway/pkg/di"
	"log"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start sever: ", diErr)
	} else {
		server.Start()
	}
}
