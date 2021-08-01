package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/model"
	"github.com/Gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrEmptyUser        = errors.New("empty user name")
	ErrEmptyAccountInfo = errors.New("empty account info")
	ErrEmptyDate        = errors.New("empty date")
)

var (
	setAccountPath = "/account/set"
	getAccountPath = "/account/get"
)

type getAccountHeader struct {
	ID string `header:"id" binding:"required"`
}

func (r *repository) setAccount(ctx context.Context, account model.AccountInfo) error {

	// check kakao account has previous.
	prevAccount, err := r.db.GetAccount(ctx, account.Id)
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
		return r.db.SetAccount(ctx, account)
	}

	return r.db.SetAccount(ctx, account)
}

func (r *repository) getAccount(ctx context.Context, id string) (*model.AccountInfo, error) {
	return r.db.GetAccount(ctx, id)
}

func (r *repository) setAccountHandler(ctx context.Context) {
	account := model.AccountInfo{}

	r.server.Client.Router.POST(setAccountPath, func(c *gin.Context) {

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
		prevAccount, err := r.getAccount(ctx, account.Id)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": model.FailResponse})
			return
		}

		if prevAccount != nil {
			// replace account info to new own.
			err := r.setAccount(ctx, account)
			if err != nil {
				logger.Error(err)
				c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
				return
			}

			c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
			return
		}

		// register first account info.
		err = r.setAccount(ctx, account)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
		return

	})
}

func (r *repository) getAccountHandler(ctx context.Context) {

	r.server.Client.Router.GET(getAccountPath, func(c *gin.Context) {

		header := getAccountHeader{}

		if err := c.ShouldBindHeader(&header); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched})
			return
		}

		logger.Info(fmt.Sprintf("header:%v", c.Request.Header))
		logger.Info(fmt.Sprintf("body:%v", c.Request.Body))

		ac, err := r.getAccount(ctx, string(header.ID))
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
			return
		}

		// account valid check.
		if ac == nil {
			logger.Error(fmt.Errorf("account info is empty(%s)", header.ID))
			c.JSON(model.FailResponseCode, gin.H{"error": fmt.Sprintf("account id(%s) is empty", header.ID)})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{
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
