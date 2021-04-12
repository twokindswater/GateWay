package db

import (
	"context"
	"fmt"
	"github.com/Gateway/internal/homebody/data"
	"github.com/Gateway/pkg/logger"
)

func (d *DB) SetAccount(ctx context.Context, a data.AccountInfo) error {

	// encoding account.
	b, err := d.serializer.Encode(ctx, a)
	if err != nil {
		logger.Error(err)
		return errSerializer
	}

	// get account path.
	p := GetAccountPath(a.Id)

	// set account.
	err = d.Client.Set(ctx, p, b)
	if err != nil {
		logger.Error(err)
		return errSetDataBase
	}

	return nil
}

func (d *DB) GetAccount(ctx context.Context, id string) (*data.AccountInfo, error) {

	// get account path.
	p := GetAccountPath(id)

	// get account.
	b, err := d.Client.Get(ctx, p)
	if err != nil {
		logger.Error(err)
		return nil, errGetDataBase
	}

	// has no data.
	if b == nil {
		return nil, nil
	}

	// initialize account.
	a := &data.AccountInfo{}

	// decoding account.
	err = d.serializer.Decode(ctx, b, a)
	if err != nil {
		logger.Error(err)
		return nil, errDecoding
	}

	return a, nil
}

// h/a/{{id}}
func GetAccountPath(id string) string {
	return fmt.Sprintf("%s%s%s%s%s", data.ServicePrefix, data.Delimiter, data.AccountPrefix, data.Delimiter, id)
}
