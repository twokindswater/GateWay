package db

import (
	"github.com/HomeLongServer/pkg/database"
	"github.com/HomeLongServer/pkg/serializer"
)

type DB struct {
	Client     database.DataBase
	serializer serializer.Serializer
}

func Init(dbType, address string, serializer serializer.Serializer) (*DB, error) {
	client, err := database.Init(dbType, address)
	if err != nil {
		return nil, err
	}
	return &DB{
		Client:     client,
		serializer: serializer,
	}, nil
}

type Account struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
