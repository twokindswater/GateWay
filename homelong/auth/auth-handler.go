package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Auth) AddHandler() {
	a.server.Router.POST("/login/a", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	a.server.Router.GET("/login/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	})
}
