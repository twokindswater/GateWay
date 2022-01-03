package web

import (
	"context"

	"firebase.google.com/go/messaging"
	"github.com/Gateway/internal/homebody/db"
	"github.com/gin-gonic/gin"
)

type (
	Web struct {
		engine *gin.Engine
		port   string
		db     *db.DB
		client *messaging.Client
	}

	Config struct {
		Port string `config:"port"`
	}
)

func Init(ctx context.Context, config Config, db *db.DB, client *messaging.Client) (*Web, error) {

	web := &Web{
		engine: gin.Default(),
		port:   config.Port,
		db:     db,
		client: client,
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

	// friend
	w.AddFriendHandler(ctx)
	w.GetAllFriendsHandler(ctx)
	w.GetFriendHandler(ctx)
	w.DeleteFriendHandler(ctx)

	// knock
	w.KnockHandler(ctx)
}
