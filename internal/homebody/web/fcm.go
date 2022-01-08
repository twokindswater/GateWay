package web

import (
	"context"
	"net/http"

	"firebase.google.com/go/messaging"
	"github.com/Gateway/internal/homebody/model"
	"github.com/Gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

type KnockHeader struct {
	ID   string `header:"id" binding:"required"`
	FID  string `header:"fid" binding:"required"`
	TIME string `header:"time" binding:"required"`
}

var (
	sendKnockPath = "/knock/set"
)

func (w *Web) KnockHandler(ctx context.Context) {
	w.engine.GET(sendKnockPath, func(c *gin.Context) {
		header := KnockHeader{}

		if err := c.ShouldBindHeader(&header); err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched})
			return
		}

		// This registration token comes from the client FCM SDKs.
		account, err := w.db.GetAccount(ctx, header.FID)
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

		registrationToken := account.Token
		logger.Info("friend id: " + account.Id + " token: " + registrationToken)

		// See documentation on defining a message payload.
		message := &messaging.Message{
			Data: map[string]string{
				"senderId": header.ID,
				"sentTime": header.TIME,
			},
			Token: registrationToken,
		}
		// Send a message to the device corresponding to the provided
		// registration token.
		response, err := w.client.Send(ctx, message)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}
		c.JSON(model.SuccessResponseCode, model.SuccessResponse)
		logger.Info("Successfully sent message: " + response)

	})
}
