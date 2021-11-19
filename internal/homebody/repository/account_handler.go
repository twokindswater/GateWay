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

type getAccountHeader struct {
	ID string `header:"id" binding:"required"`
}

func (r *repository) SetAccountHandler(ctx context.Context) {
	account := model.AccountInfo{}

	r.server.Client.Router.POST(setAccountPath, func(c *gin.Context) {

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with account info struct(%v)", c.Request, account)))

			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		// check account has previous.
		prevAccount, err := r.db.GetAccount(ctx, account.Id)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if prevAccount != nil {
			// replace account info to new own.
			err := r.db.SetAccount(ctx, account)
			if err != nil {
				logger.Error(err)
				c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
				return
			}

			c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
			return
		}

		// register first account info.
		err = r.db.SetAccount(ctx, account)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
		return
	})
}

func (r *repository) GetAccountHandler(ctx context.Context) {
	header := getAccountHeader{}

	r.server.Client.Router.GET(getAccountPath, func(c *gin.Context) {

		if err := c.ShouldBindHeader(&header); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched +
				fmt.Sprintf("expected(%v) actual(%v)", header, c.Request)})
			return
		}

		if header.ID == "" {
			logger.Error(fmt.Errorf("header.ID is null"))
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched +
				"header.id is empty"})
			return
		}

		ac, err := r.db.GetAccount(ctx, header.ID)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
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
		return
	})
}

func (r *repository) UpdateAccountHandler(ctx context.Context) {
	account := model.AccountInfo{}

	r.server.Client.Router.POST(updateAccountPath, func(c *gin.Context) {

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with account info struct(%v)", c.Request, account)))

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		preAccount, err := r.db.GetAccount(ctx, account.Id)
		if err != nil {
			err = r.db.SetAccount(ctx, account)
			if err != nil {
				logger.Error(err)
				c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			}
		}

		updateAccount(preAccount, &account)

		err = r.db.SetAccount(ctx, *preAccount)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
		}

		c.JSON(model.SuccessResponseCode, gin.H{
			"name":      preAccount.Name,
			"image":     preAccount.Image,
			"ssid":      preAccount.SSID,
			"bssid":     preAccount.BSSID,
			"street":    preAccount.Street,
			"initDate":  preAccount.InitDate,
			"latitude":  preAccount.Latitude,
			"longitude": preAccount.Longitude,
		})
		return
	})
}

func (r *repository) DeleteAccountHandler(ctx context.Context) {

	r.server.Client.Router.GET(deleteAccountPath, func(c *gin.Context) {
		header := getAccountHeader{}

		if err := c.ShouldBindHeader(&header); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched})
			return
		}

		if err := r.db.DeleteAccount(ctx, header.ID); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, model.SuccessResponse)
		return
	})
}

func updateAccount(prev *model.AccountInfo, new *model.AccountInfo) {
	if new.Name != "" {
		prev.Name = new.Name
	}

	if new.Image != "" {
		prev.Image = new.Image
	}

	if new.SSID != "" {
		prev.SSID = new.SSID
	}

	if new.BSSID != "" {
		prev.BSSID = new.BSSID
	}

	if new.Street != "" {
		prev.Street = new.Street
	}

	if new.InitDate != 0 {
		prev.InitDate = new.InitDate
	}

	if new.Latitude != 0 {
		prev.Latitude = new.Latitude
	}

	if new.Longitude != 0 {
		prev.Longitude = new.Longitude
	}
}
