package db

import (
	"context"
	"fmt"
	"github.com/Gateway/internal/homebody/model"
	"github.com/Gateway/pkg/logger"
)

func (db *DB) SetAccount(ctx context.Context, a model.Account) error {

	// encoding account.
	b, err := db.serializer.Encode(ctx, a)
	if err != nil {
		logger.Error(err)
		return errSerializer
	}

	// get account path.
	p := GetAccountPath(a.Id)

	// set account.
	err = db.Client.Set(ctx, p, b)
	if err != nil {
		logger.Error(err)
		return errSetDataBase
	}

	return nil
}

func (db *DB) GetAccount(ctx context.Context, id string) (*model.Account, error) {

	// get account path.
	p := GetAccountPath(id)

	// get account.
	b, err := db.Client.Get(ctx, p)
	if err != nil {
		logger.Error(err)
		return nil, errGetDataBase
	}

	// has no data.
	if b == nil {
		return nil, nil
	}

	// initialize account.
	a := &model.Account{}

	// decoding account.
	err = db.serializer.Decode(ctx, b, a)
	if err != nil {
		logger.Error(err)
		return nil, errDecoding
	}

	return a, nil
}

func (db *DB) DeleteAccount(ctx context.Context, id string) error {

	path := GetAccountPath(id)

	return db.Client.Del(ctx, path)
}

// h/a/{{id}}
func GetAccountPath(id string) string {
	return fmt.Sprintf("%s%s%s%s%s", model.ServicePrefix, model.Delimiter, model.AccountPrefix, model.Delimiter, id)
}
