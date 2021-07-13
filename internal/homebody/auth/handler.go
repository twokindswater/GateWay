package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/data"
	"github.com/Gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	ErrEmptyUser        = errors.New("empty user name")
	ErrEmptyAccountInfo = errors.New("empty account info")
)

func (a *auth) AddHandler(ctx context.Context) {

	// set account info handler.
	a.kakaoHandler(ctx)
	a.facebookHandler(ctx)

	// get account info handler.
	a.getAccountHandler(ctx)
}

// get account info.
func (a *auth) getAccountHandler(ctx context.Context) {

	a.server.Client.Router.GET("account/info/get/:user", func(c *gin.Context) {

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
			"name":      ac.Name,
			"image":     ac.Image,
			"ssid":      ac.SSID,
			"bssid":     ac.BSSID,
			"street":    ac.Street,
			"latitude":  ac.Latitude,
			"longitude": ac.Longitude,
		})
	})
}
