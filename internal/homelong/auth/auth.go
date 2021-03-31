package auth

import (
	"context"
	"github.com/HomeLongServer/internal/homelong/db"
	"github.com/HomeLongServer/internal/homelong/web"
)

type Auth struct {
	server *web.Web
	db     *db.DB
}

func Init(s *web.Web, db *db.DB) (*Auth, error) {
	return &Auth{
		server: s,
		db:     db,
	}, nil
}

func (a *Auth) AddHandler(ctx context.Context) {
	addKakaoAuthHandler(ctx, a.server, a.db)

}
