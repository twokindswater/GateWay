package web

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/Gateway/internal/homebody/db"
	"github.com/gin-gonic/gin"
)

type Web struct {
	engine *gin.Engine
	port   string
	db     *db.DB
	fb     *firebase.App
}

func Init(ctx context.Context, port string, db *db.DB, fb *firebase.App) (*Web, error) {

	web := &Web{
		engine: gin.Default(),
		port:   port,
		db:     db,
		fb:     fb,
	}

	return web, nil
}

func (w *Web) Start(ctx context.Context) error {
	return w.engine.Run(w.port)
}

func (w *Web) AddHandler(ctx context.Context) {

	w.SetAccountHandler(ctx)
	w.GetAccountHandler(ctx)

	w.SetDayTimeHandler(ctx)
	w.GetDayTimeHandler(ctx)

}
