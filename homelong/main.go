package main

import (
	"fmt"
	"github.com/HomeLongServer/homelong/auth"
	"github.com/HomeLongServer/homelong/config"
	"github.com/HomeLongServer/pkg/banner"
	"github.com/HomeLongServer/pkg/database"
	"github.com/HomeLongServer/pkg/logger"
	"github.com/HomeLongServer/pkg/webserver"
)

func main() {

	// initialize context
	//ctx := context.Background()

	// initialize logger
	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}

	// get configuration
	cfg := config.GetConfig()

	// initialize database
	db, err := database.InitDB(cfg.DBCfg.Type, cfg.DBCfg.Address)
	if err != nil {
		logger.Error(err)
	}

	// initialize api server
	server, err := webServer.InitWebServer(cfg.WebServer.Port)
	if err != nil {
		logger.Error(err)
	}

	// initialize auth handler
	authHandler, err := auth.InitAuth(server, db)
	if err != nil {
		logger.Error(err)
	}
	authHandler.AddHandler()

	// start message
	banner.ShowBanner(banner.HomeLongBanner)
	fmt.Printf("server start!\n")

	err = server.Run()
	if err != nil {
		logger.Error(err)
	}

	// server done
	logger.Info("server done!")
}
