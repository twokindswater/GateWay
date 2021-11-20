package web

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (w *Web) KakaoAuthHandler(ctx context.Context) {
	w.engine.GET("/callbacks/kakao/sign_in", func(c *gin.Context) {

		values := map[string]string{
			"grant_type": "authorization_code",
			"client_id":  "2ca0713904e3a50b093a9b16cd14c634",
		}
		json_data, err := json.Marshal(values)

		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post("kauth.kakao.com/oauth/token", "application/json",
			bytes.NewBuffer(json_data))

		fmt.Println(resp)

	})

	w.engine.POST("/callbacks/kakao/token", func(c *gin.Context) {

	})
}
