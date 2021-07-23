package auth

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

var (
	setAccountPath = "/account/set"
	getAccountPath = "/account/get"
)

type getAccountHeader struct {
	ID string `header:"id" binding:"required"`
}

func (a *auth) AddHandler(ctx context.Context) {

	a.setAccountHandler(ctx)
	a.getAccountHandler(ctx)
}

func (a *auth) setAccountHandler(ctx context.Context) {
	account := data.AccountInfo{}

	a.server.Client.Router.POST(setAccountPath, func(c *gin.Context) {

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with account info struct", c.Request)))
			buf := make([]byte, 100)
			n := c.Request.Body
			logger.Error(errors.New(fmt.Sprintf("n:%v", n)))
			nn, e := n.Read(buf)
			logger.Error(errors.New(fmt.Sprintf("nn:%d, e:%v", nn, e)))
			logger.Error(errors.New(fmt.Sprintf("buf:%v", buf)))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// check account has previous.
		prevAccount, err := a.getAccount(ctx, account.Id)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": data.FailResponse})
			return
		}

		if prevAccount != nil {
			// replace account info to new own.
			err := a.setAccount(ctx, account)
			if err != nil {
				logger.Error(err)
				c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
				return
			}

			c.JSON(data.SuccessResponseCode, gin.H{"status": data.SuccessResponse})
			return
		}

		// register first account info.
		err = a.setAccount(ctx, account)
		if err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		c.JSON(data.SuccessResponseCode, gin.H{"status": data.SuccessResponse})
		return

	})
}

func (a *auth) getAccountHandler(ctx context.Context) {

	a.server.Client.Router.GET(getAccountPath, func(c *gin.Context) {

		header := getAccountHeader{}

		if err := c.ShouldBindHeader(&header); err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.HeaderIsNotMatched})
			return
		}

		logger.Info(fmt.Sprintf("header:%v", c.Request.Header))
		logger.Info(fmt.Sprintf("body:%v", c.Request.Body))

		ac, err := a.getAccount(ctx, string(header.ID))
		if err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		// account valid check.
		if ac == nil {
			logger.Error(fmt.Errorf("account info is empty(%s)", header.ID))
			c.JSON(data.FailResponseCode, gin.H{"error": fmt.Sprintf("account id(%s) is empty", header.ID)})
			return
		}

		c.JSON(data.SuccessResponseCode, gin.H{
			"name":      ac.Name,
			"image":     ac.Image,
			"ssid":      ac.SSID,
			"bssid":     ac.BSSID,
			"street":    ac.Street,
			"initDate":  ac.InitDate,
			"latitude":  ac.Latitude,
			"longitude": ac.Longitude,
		})
	})
}
