package main

import (
	"context"
	"github.com/Gateway/internal/homebody/account"
	"github.com/Gateway/internal/homebody/config"
	"github.com/Gateway/internal/homebody/data"
	"github.com/Gateway/internal/homebody/db"
	"github.com/Gateway/internal/homebody/time"
	"github.com/Gateway/internal/homebody/web"
	"github.com/Gateway/pkg/banner"
	"github.com/Gateway/pkg/logger"
	"github.com/Gateway/pkg/serializer"
)

func main() {

	// initialize context.
	ctx := context.Background()

	// initialize logger.
	err := logger.Init(ctx)
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

	// initialize account handler.
	accountHandler, err := account.Init(server, database)
	if err != nil {
		panic(err)
	}

	timeHandler, err := time.Init(server, database)
	if err != nil {
		panic(err)
	}

	// register handler.
	accountHandler.AddHandler(ctx)
	timeHandler.AddHandler(ctx)

	// start message.
	banner.ShowBanner(data.HomeLongBanner)
	logger.Info("server start!\n")

	// web framework start.
	err = server.Client.Run()
	if err != nil {
		logger.Error(err)
	}

	// server done
	logger.Info("server done!")
}
