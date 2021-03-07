package database

import (
	"errors"
	"github.com/HomeLongServer/pkg/database/redis"
)

const (
	Redis = "redis"
)

var (
	errUndefinedDBType = errors.New("undefined database type")
)

type DataBase interface {
}

var dataBase DataBase

func GetDataBase() DataBase {
	return dataBase
}

func InitDB(dbType, address string) (DataBase, error) {
	switch dbType {
	case Redis:
		dataBase = redis.InitRedis(address)
	default:
		return nil, errUndefinedDBType
	}
	return dataBase, nil
}
