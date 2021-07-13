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

func (a *auth) kakaoHandler(ctx context.Context) {
	account := &data.AccountInfo{}

	a.server.Client.Router.POST("/login/kakao", func(c *gin.Context) {

		// checking request body is matched with account info.
		if err := c.BindJSON(account); err != nil {
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with account info struct", c.Request)))
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

		if prevAccount == nil {
			// register first account info.
			err := a.setAccount(ctx, *account)
			if err != nil {
				logger.Error(err)
				c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
				return
			}

			c.JSON(data.SuccessResponseCode, gin.H{"status": data.SuccessResponse})
			return
		}

		// login success.
		c.JSON(data.SuccessResponseCode, gin.H{"status": data.SuccessResponse})
		return
	})
}
