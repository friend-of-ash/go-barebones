package main

import (
	"pkg/config"
	"pkg/log"
	"pkg/router"
	"pkg/server"
)

func init() {
	config.LoadConfig()
	log.Infof("Config loaded for environment: %v", config.Conf.ENVIRONMENT)
}

func main() {
	log.Info("Hello, server")
	server.NewServer(router.GetInitialisedRouter())
}
