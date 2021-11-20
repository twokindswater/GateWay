package db

import (
	"context"
	"github.com/Gateway/pkg/database"
	"github.com/Gateway/pkg/serializer"
)

type DB struct {
	Client     database.DataBase
	serializer serializer.Serializer
}

func Init(ctx context.Context, config Config, serializer serializer.Serializer) (*DB, error) {
	client, err := database.Init(config.Type, config.Address)
	if err != nil {
		return nil, err
	}
	return &DB{
		Client:     client,
		serializer: serializer,
	}, nil
}

func (db *DB) Clear(ctx context.Context) error {
	return db.Client.Clear(ctx)
}
