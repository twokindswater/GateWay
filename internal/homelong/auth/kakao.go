package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/HomeLongServer/internal/homelong/db"
	"github.com/HomeLongServer/internal/homelong/web"
	"github.com/HomeLongServer/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addKakaoAuthHandler(ctx context.Context, w *web.Web, d *db.DB) {

	var account *db.Account

	// Register kakao account.
	w.Client.Router.POST("/register/kakao", func(context *gin.Context) {

		// checking request can bind with kakao info.
		if err := context.ShouldBindJSON(account); err != nil {
			// request is not matched with kakao info struct.
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with kakao info struct", context.Request)))
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// check kakao account has previous.
		prevAccount, err := d.GetAccount(ctx, account.Id)
		if err != nil {
			// fail to get account info.
			logger.Error(err)
			context.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("server occured error")})
			return
		}

		if prevAccount != nil {
			// previous account is exist.
			// return http.StatusNotAcceptable status with account already exist error message.
			context.JSON(http.StatusNotAcceptable, gin.H{"error": fmt.Sprintf("Account already exist(%v)", account.Id)})
			return
		}

		// store new account.
		err = d.SetAccount(ctx, *account)
		if err != nil {
			logger.Error(err)
			context.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("server occured error")})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	})

	// Login kakao user.
	w.Client.Router.GET("/login/kakao", func(context *gin.Context) {

		// checking request can bind with kakao info.
		if err := context.ShouldBindJSON(account); err != nil {
			// request is not matched with kakao info struct.
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with kakao info struct", context.Request)))
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// check stored kakao account.
		storedAccount, err := d.GetAccount(ctx, account.Id)
		if err != nil {
			logger.Error(err)
			context.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("server occured error")})
			return
		}

		// check account's token is matched with stored account's token.
		if account.Token != storedAccount.Token {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("request kakao account is not valid")})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	})
}
