package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/data"
	"github.com/Gateway/internal/homebody/db"
	"github.com/Gateway/internal/homebody/web"
	"github.com/Gateway/pkg/logger"
)

type auth struct {
	server *web.Web
	db     *db.DB
}

func Init(s *web.Web, db *db.DB) (*auth, error) {
	return &auth{
		server: s,
		db:     db,
	}, nil
}

func (a *auth) setAccount(ctx context.Context, account data.AccountInfo) error {

	// check kakao account has previous.
	prevAccount, err := a.db.GetAccount(ctx, account.Id)
	if err != nil {
		// fail to get account info.
		logger.Error(err)
		return errors.New(fmt.Sprintf("get account from db error (account=%v)", account.Id))
	}

	if prevAccount != nil {
		// previous account is exist.
		logger.Info(fmt.Sprintf("account already exist(%v)", account.Id))
		logger.Info(fmt.Sprintf("prevAccount info:%v", prevAccount))

		// replace account.
		return a.db.SetAccount(ctx, account)
	}

	return a.db.SetAccount(ctx, account)
}

func (a *auth) getAccount(ctx context.Context, id string) (*data.AccountInfo, error) {
	return a.db.GetAccount(ctx, id)
}
