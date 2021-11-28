package web

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/Gateway/internal/homebody/db"
	"github.com/gin-gonic/gin"
)

type (
	Web struct {
		engine *gin.Engine
		port   string
		db     *db.DB
		fb     *firebase.App
	}

	Config struct {
		Port string `config:"port"`
	}
)

func Init(ctx context.Context, config Config, db *db.DB, fb *firebase.App) (*Web, error) {

	web := &Web{
		engine: gin.Default(),
		port:   config.Port,
		db:     db,
		fb:     fb,
	}

	return web, nil
}

func (w *Web) Start(ctx context.Context) error {
	w.AddHandler(ctx)

	return w.engine.Run(w.port)
}

func (w *Web) AddHandler(ctx context.Context) {

	w.SetAccountHandler(ctx)
	w.GetAccountHandler(ctx)

	w.SetLocationHandler(ctx)
	w.SetWifiHandler(ctx)

	w.SetDayTimeHandler(ctx)
	w.GetDayTimeHandler(ctx)
}
