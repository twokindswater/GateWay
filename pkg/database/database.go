package database

import (
	"context"
	"errors"
	"github.com/HomeLongServer/pkg/database/redis"
)

// database type.
const (
	Redis = "redis"
)

var (
	errUndefinedDBType = errors.New("undefined database type or wrong database type")
)

type DataBase interface {
	Get(context.Context, string) ([]byte, error)
	Set(context.Context, string, []byte) error
}

func Init(dbType, address string) (DataBase, error) {
	var dataBase DataBase

	switch dbType {
	case Redis:
		// initialize redis database.
		dataBase = redis.InitRedis(address)
	default:
		return nil, errUndefinedDBType
	}

	return dataBase, nil
}
