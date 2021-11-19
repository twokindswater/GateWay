package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	SuccessResponse     = fmt.Sprintf("success")
	SuccessResponseCode = http.StatusOK
	FailResponseCode    = http.StatusBadRequest
	HeaderIsNotMatched  = fmt.Sprintf("header is not matched")
	FailResponse        = gin.H{"error": "server occured error"}
)
