package db

import (
	"context"
	"fmt"
	"github.com/HomeLongServer/internal/homebody/consts"
)

func (d *DB) GetAccount(ctx context.Context, id string) (*Account, error) {

	// get account path.
	p := GetAccountPath(id)

	// get account.
	b, err := d.Client.Get(ctx, p)
	if err != nil {
		return nil, err
	}

	// initialize account.
	a := &Account{}

	// decoding account.
	err = d.serializer.Decode(ctx, b, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (d *DB) SetAccount(ctx context.Context, a Account) error {

	// encoding account.
	b, err := d.serializer.Encode(ctx, a)
	if err != nil {
		return err
	}

	// get account path.
	p := GetAccountPath(a.Id)

	// set account.
	err = d.Client.Set(ctx, p, b)
	if err != nil {
		return err
	}

	return nil
}

func GetAccountPath(id string) string {
	return fmt.Sprintf("%s%s%s%s%s", consts.ServicePrefix, consts.Delimiter, consts.AccountPrefix, consts.Delimiter, id)
}
