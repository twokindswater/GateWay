package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/data"
	"github.com/Gateway/internal/homebody/db"
	"github.com/Gateway/internal/homebody/web"
	"github.com/Gateway/pkg/logger"
)

type Account struct {
	server *web.Web
	db     *db.DB
}

func Init(s *web.Web, db *db.DB) (*Account, error) {
	return &Account{
		server: s,
		db:     db,
	}, nil
}

func (a *Account) setAccount(ctx context.Context, account data.AccountInfo) error {

	// check kakao account has previous.
	prevAccount, err := a.db.GetAccount(ctx, account.Id)
	if err != nil {
		// fail to get account info.
		logger.Error(err)
		return errors.New(fmt.Sprintf("get account error (account=%v)", account.Id))
	}

	if prevAccount != nil {
		// previous account is exist.
		return errors.New(fmt.Sprintf("account already exist(%v)", account.Id))
	}

	return a.db.SetAccount(ctx, account)
}

func (a *Account) getAccount(ctx context.Context, id string) (*data.AccountInfo, error) {
	return a.db.GetAccount(ctx, id)
}

func (a *Account) updateAccount(ctx context.Context, account data.AccountInfo) error {
	return a.db.SetAccount(ctx, account)
}
