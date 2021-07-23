package data

import (
	"fmt"
	"net/http"
)

var (
	SuccessResponse     = fmt.Sprintf("success")
	SuccessResponseCode = http.StatusOK
	FailResponse        = fmt.Sprintf("server occured error")
	FailResponseCode    = http.StatusBadRequest
	HeaderIsNotMatched  = fmt.Sprintf("header is not matched")
)
