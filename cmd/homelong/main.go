package main

import (
	"context"
	"github.com/HomeLongServer/internal/homelong/auth"
	"github.com/HomeLongServer/internal/homelong/config"
	"github.com/HomeLongServer/internal/homelong/consts"
	"github.com/HomeLongServer/internal/homelong/db"
	"github.com/HomeLongServer/internal/homelong/web"
	"github.com/HomeLongServer/pkg/banner"
	"github.com/HomeLongServer/pkg/logger"
	"github.com/HomeLongServer/pkg/serializer"
)

func main() {

	// initialize context.
	ctx := context.Background()

	// initialize logger.
	err := logger.Init()
	if err != nil {
		panic(err)
	}

	// get configuration.
	cfg := config.GetConfig()

	// initialize serializer.
	serializer, err := serializer.Init(cfg.Serializer.Type)
	if err != nil {
		panic(err)
	}

	// initialize database.
	database, err := db.Init(cfg.DB.Type, cfg.DB.Address, serializer)
	if err != nil {
		panic(err)
	}

	// initialize api server.
	server, err := web.Init(cfg.Web.Port)
	if err != nil {
		panic(err)
	}

	// initialize auth handler.
	authHandler, err := auth.Init(server, database)
	if err != nil {
		panic(err)
	}

	// register auth handler.
	authHandler.AddHandler(ctx)

	// start message.
	banner.ShowBanner(consts.HomeLongBanner)
	logger.Info("server start!\n")

	// web framework start.
	err = server.Client.Run()
	if err != nil {
		logger.Error(err)
	}

	// server done
	logger.Info("server done!")
}
