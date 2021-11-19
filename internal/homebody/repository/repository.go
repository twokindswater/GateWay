package repository

import (
	"context"
	"github.com/Gateway/internal/homebody/db"
	"github.com/Gateway/internal/homebody/web"
)

type repository struct {
	server *web.Web
	db     *db.DB
}

func Init(s *web.Web, db *db.DB) (*repository, error) {
	return &repository{
		server: s,
		db:     db,
	}, nil
}

func (r *repository) AddHandler(ctx context.Context) {

	r.SetAccountHandler(ctx)
	r.GetAccountHandler(ctx)

	r.SetDayTimeHandler(ctx)
	r.GetDayTimeHandler(ctx)

}
