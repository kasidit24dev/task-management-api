package main

import (
	"task-management-api/config"
	_ "task-management-api/docs"
	log "task-management-api/logger"
	"task-management-api/server"
)

func main() {

	cfg := config.LoadConfigs()
	logger := log.InitLogger(cfg)

	server.NewEchoServer(cfg, logger).Start()

}
