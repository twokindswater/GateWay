package db

import "errors"

var (
	errSerializer = errors.New("serializer error")

	errDecoding = errors.New("decoding error")

	errSetDataBase = errors.New("database set error")
	errGetDataBase = errors.New("database get error")
)
