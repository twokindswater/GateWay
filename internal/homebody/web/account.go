package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Gateway/internal/homebody/model"
	"github.com/Gateway/pkg/logger"
	"github.com/Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AccountHeader struct {
	ID string `header:"id" binding:"required"`
}

type FriendHeader struct {
	ID  string `header:"id" binding:"required"`
	FID string `header:"fid" binding:"omitempty"`
}

var (
	setAccountPath    = "/account/set"
	getAccountPath    = "/account/get"
	deleteAccountPath = "/account/delete"

	setLocationPath = "/location/set"
	setWifiPath     = "/wifi/set"

	setFriendPath    = "/friend/set"
	getFriendPath    = "/friend/get"
	getAllFriendPath = "/friend/get/all"
	deleteFrindPath  = "/friend/delete"
)

var (
	ErrEmptyAccount = errors.New("empty account")
)

func (w *Web) SetAccountHandler(ctx context.Context) {
	account := model.Account{}

	w.engine.POST(setAccountPath, func(c *gin.Context) {

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(err)

			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		err := w.db.SetAccount(ctx, account)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
		return
	})
}

func (w *Web) GetAccountHandler(ctx context.Context) {
	header := AccountHeader{}

	w.engine.GET(getAccountPath, func(c *gin.Context) {
		if err := c.ShouldBindHeader(&header); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched})
			return
		}

		if header.ID == "" {
			logger.Error(fmt.Errorf("header.ID is null"))
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched})
			return
		}

		account, err := w.db.GetAccount(ctx, header.ID)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if account == nil {
			logger.Error(ErrEmptyAccount)
			c.JSON(http.StatusBadGateway, gin.H{"error": ErrEmptyAccount.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{
			"name":      account.Name,
			"image":     account.Image,
			"ssid":      account.SSID,
			"bssid":     account.BSSID,
			"street":    account.Street,
			"initDate":  account.InitDate,
			"latitude":  account.Latitude,
			"longitude": account.Longitude,
			"friend":    account.Friends,
		})
		return
	})
}

func (w *Web) DeleteAccountHandler(ctx context.Context) {

	w.engine.GET(deleteAccountPath, func(c *gin.Context) {
		header := AccountHeader{}

		if err := c.ShouldBindHeader(&header); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched})
			return
		}

		if err := w.db.DeleteAccount(ctx, header.ID); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, model.SuccessResponse)
		return
	})
}

func (w *Web) SetLocationHandler(ctx context.Context) {
	account := model.Account{}

	w.engine.POST(setLocationPath, func(c *gin.Context) {

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(err)

			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		// check account has previous.
		prevAccount, err := w.db.GetAccount(ctx, account.Id)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if prevAccount != nil {

			prevAccount.Latitude = account.Latitude
			prevAccount.Longitude = account.Longitude
			prevAccount.Street = account.Street

			err := w.db.SetAccount(ctx, *prevAccount)
			if err != nil {
				logger.Error(err)
				c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
				return
			}

			c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
			return
		}

		logger.Error(fmt.Errorf("account info is empty(%s)", account.Id))
		c.JSON(model.FailResponseCode, gin.H{"error": fmt.Sprintf("account id(%s) is empty", account.Id)})
	})
}

func (w *Web) SetWifiHandler(ctx context.Context) {
	account := model.Account{}

	w.engine.POST(setWifiPath, func(c *gin.Context) {

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(err)

			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		// check account has previous.
		prevAccount, err := w.db.GetAccount(ctx, account.Id)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if prevAccount != nil {

			prevAccount.SSID = account.SSID
			prevAccount.BSSID = account.BSSID

			err := w.db.SetAccount(ctx, *prevAccount)
			if err != nil {
				logger.Error(err)
				c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
				return
			}

			c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
			return
		}

		logger.Error(fmt.Errorf("account info is empty(%s)", account.Id))
		c.JSON(model.FailResponseCode, gin.H{"error": fmt.Sprintf("account id(%s) is empty", account.Id)})
	})
}

func (w *Web) GetAllFriendsHandler(ctx context.Context) {
	friendHeader := &FriendHeader{}

	w.engine.GET(getAllFriendPath, func(c *gin.Context) {
		if err := c.ShouldBindHeader(&friendHeader); err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		account, err := w.db.GetAccount(ctx, friendHeader.ID)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if account == nil {
			logger.Error(ErrEmptyAccount)
			c.JSON(http.StatusBadGateway, gin.H{"error": ErrEmptyAccount.Error()})
			return
		}

		//var body []map[string]interface{}
		var friends []model.Friend

		for _, fid := range account.Friends {
			account, err := w.db.GetAccount(ctx, fid)
			if err != nil {
				logger.Error(err)
				c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
				return
			}
			// ToDO : make AtHome DB and manage home user
			friend := model.Friend{
				Id:     account.Id,
				Name:   account.Name,
				Image:  account.Image,
				AtHome: account.AtHome,
			}

			friends = append(friends, friend)
		}

		c.JSON(model.SuccessResponseCode, friends)

		return
	})
}
func (w *Web) AddFriendHandler(ctx context.Context) {
	friendHeader := &FriendHeader{}

	w.engine.POST(setFriendPath, func(c *gin.Context) {
		if err := c.ShouldBindHeader(&friendHeader); err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		account, err := w.db.GetAccount(ctx, friendHeader.ID)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if account == nil {
			logger.Error(ErrEmptyAccount)
			c.JSON(http.StatusBadGateway, gin.H{"error": ErrEmptyAccount.Error()})
			return
		}

		account.Friends = append(account.Friends, friendHeader.FID)

		if err := w.db.SetAccount(ctx, *account); err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})
		return
	})
}

func (w *Web) GetFriendHandler(ctx context.Context) {
	friendHeader := &FriendHeader{}

	w.engine.GET(getFriendPath, func(c *gin.Context) {
		if err := c.ShouldBindHeader(&friendHeader); err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		account, err := w.db.GetAccount(ctx, friendHeader.ID)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if account == nil {
			logger.Error(ErrEmptyAccount)
			c.JSON(http.StatusBadGateway, gin.H{"error": ErrEmptyAccount.Error()})
			return
		}

		// ToDO : make AtHome DB and manage home user
		friend := model.Friend{Id: account.Id, Name: account.Name,
			Image: account.Image, AtHome: account.AtHome}
		c.JSON(model.SuccessResponseCode, friend)

		return
	})
}

func (w *Web) DeleteFriendHandler(ctx context.Context) {
	friendHeader := FriendHeader{}

	w.engine.GET(deleteFrindPath, func(c *gin.Context) {
		if err := c.ShouldBindHeader(&friendHeader); err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		account, err := w.db.GetAccount(ctx, friendHeader.ID)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if account == nil {
			logger.Error(ErrEmptyAccount)
			c.JSON(http.StatusBadGateway, gin.H{"error": ErrEmptyAccount.Error()})
			return
		}

		account.Friends = utils.Remove(account.Friends, friendHeader.FID)

		err = w.db.SetAccount(ctx, *account)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(model.SuccessResponseCode, gin.H{"status": model.SuccessResponse})

		return
	})
}
