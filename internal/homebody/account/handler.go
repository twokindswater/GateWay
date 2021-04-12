package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/data"
	"github.com/Gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrEmptyUser        = errors.New("empty user name")
	ErrEmptyAccountInfo = errors.New("empty account info")
)

func (a *Account) AddHandler(ctx context.Context) {
	a.kakaoAuthHandler(ctx)
	a.getAccountHandler(ctx)
	a.SetAccountAPHandler(ctx)
}

func (a *Account) getAccountHandler(ctx context.Context) {

	a.server.Client.Router.GET("account/get/:user", func(c *gin.Context) {

		// get user id.
		id := c.Param("user")
		if len(id) == 0 {
			logger.Error(ErrEmptyUser)
			c.JSON(data.FailResponseCode, gin.H{"error": fmt.Sprintf("user is not defined")})
			return
		}

		ac, err := a.getAccount(ctx, id)
		if err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		// account valid check.
		if ac == nil {
			logger.Error(ErrEmptyAccountInfo)
			c.JSON(data.FailResponseCode, gin.H{"error": fmt.Sprintf("user(%s) info is empty", id)})
			return
		}

		c.JSON(data.SuccessResponseCode, gin.H{
			"id":    ac.Id,
			"image": ac.Image,
			"ssid":  ac.SSID,
			"bssid": ac.BSSID,
		})
	})
}

func (a *Account) SetAccountAPHandler(ctx context.Context) {

	account := &data.AccountInfo{}

	a.server.Client.Router.POST("/account/set", func(c *gin.Context) {

		// checking request body is matched with account info.
		if err := c.BindJSON(account); err != nil {
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with account info struct", c.Request)))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// get account info.
		prevAccount, err := a.getAccount(ctx, account.Id)
		if err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		// checking has account info.
		if prevAccount == nil {
			logger.Error(ErrEmptyAccountInfo)
			c.JSON(data.FailResponseCode, gin.H{"error": fmt.Sprintf("user(%s) info is empty", account.Id)})
			return
		}

		// update account ap info.
		prevAccount.SSID = account.SSID
		prevAccount.BSSID = account.BSSID

		// set updated account info.
		err = a.updateAccount(ctx, *prevAccount)
		if err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		c.JSON(data.SuccessResponseCode, gin.H{"status": data.SuccessResponse})
		return
	})
}
