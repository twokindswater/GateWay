package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/model"
	"github.com/Gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type getDayTimeHeader struct {
	ID   string `header:"id" binding:"required"`
	Date int    `header:"date" binding:"required"`
}

func (r *repository) SetDayTimeHandler(ctx context.Context) {

	timeInfo := &model.DayTimeInfo{}

	r.server.Client.Router.POST("/time/day/set", func(c *gin.Context) {

		if err := c.BindJSON(timeInfo); err != nil {
			// request is not matched with day time info struct.
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with timeInfo info struct(%v)", c.Request, timeInfo)))
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
		}

		// store day time.
		err := r.setTime(ctx, timeInfo.Id, timeInfo.Date, timeInfo.Time)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{"error": err.Error()})
	})
}

func (r *repository) GetDayTimeHandler(ctx context.Context) {
	header := getDayTimeHeader{}

	r.server.Client.Router.GET("/time/day/get/:user/:date", func(c *gin.Context) {

		// get user id.
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

		if header.Date == 0 {
			logger.Error(fmt.Errorf("header.Date is zero"))
			c.JSON(model.FailResponseCode, gin.H{"error": model.HeaderIsNotMatched +
				"header.date has invalid value"})
			return
		}

		// get day time.
		dayTime, err := r.GetDayTime(ctx, header.ID, header.Date)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		// convert int to string.
		dayTimeString := strconv.Itoa(dayTime)

		c.String(model.SuccessResponseCode, dayTimeString)

		return
	})

}
