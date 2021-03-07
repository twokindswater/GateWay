package auth

import (
	"github.com/HomeLongServer/pkg/database"
	"github.com/HomeLongServer/pkg/webserver"
)

type Auth struct {
	server *webServer.WebServer
	db     database.DataBase
}

func InitAuth(s *webServer.WebServer, db database.DataBase) (*Auth, error) {
	return &Auth{
		server: s,
		db:     db,
	}, nil
}

type UserAccount struct {
	ID    string
	Email string
	Token string
}
