package main

import (
	"context"
	"github.com/Gateway/internal/homebody/config"
	"github.com/Gateway/internal/homebody/db"
	"github.com/Gateway/internal/homebody/firebase"
	"github.com/Gateway/internal/homebody/model"
	"github.com/Gateway/internal/homebody/serializer"
	"github.com/Gateway/internal/homebody/web"
	"github.com/Gateway/pkg/banner"
	"github.com/Gateway/pkg/logger"
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
	cfg := config.GetConfig(ctx)

	// initialize serializer.
	serializer, err := serializer.Init(ctx, cfg.Serializer)
	if err != nil {
		panic(err)
	}

	// initialize database.
	database, err := db.Init(ctx, cfg.DB, serializer)
	if err != nil {
		panic(err)
	}

	fb, err := firebase.Init(ctx, cfg.Firebase)
	if err != nil {
		panic(err)
	}

	// initialize webServer.
	webServer, err := web.Init(ctx, cfg.Web.Port, database, fb)
	if err != nil {
		panic(err)
	}

	// start message.
	banner.ShowBanner(model.HomeLongBanner)
	logger.Info("webServer start!\n")

	// webServer start.
	err = webServer.Start(ctx)
	if err != nil {
		logger.Error(err)
	}

	// webServer done.
	logger.Info("webServer done!")
}
