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

func (r *repository) setTime(ctx context.Context, id string, date, time int) error {

	prevDayTime, err := r.GetDayTime(ctx, id, date)
	if err != nil {
		logger.Error(err)
		return errors.New(fmt.Sprintf("get account(%s) date(%d) time(%d)", id, date, time))
	}

	if prevDayTime != 0 {
		// add prev time and current time
		time = prevDayTime + time
	}

	return r.db.SetDayTime(ctx, id, date, time)
}

func (r *repository) GetDayTime(ctx context.Context, id string, date int) (int, error) {
	return r.db.GetDayTime(ctx, id, date)
}

func (r *repository) SetDayTimeHandler(ctx context.Context) {

	timeInfo := &model.DayTimeInfo{}

	r.server.Client.Router.POST("/time/day/set", func(c *gin.Context) {

		if err := c.BindJSON(timeInfo); err != nil {
			// request is not matched with day time info struct.
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with kakao info struct", c.Request)))
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
		}

		// store day time.
		err := r.setTime(ctx, timeInfo.Id, timeInfo.Date, timeInfo.Time)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{"error": model.FailResponse})
	})
}

func (r *repository) GetDayTimeHandler(ctx context.Context) {

	r.server.Client.Router.GET("/time/day/get/:user/:date", func(c *gin.Context) {

		// get user id.
		id := c.Param("user")
		if len(id) == 0 {
			logger.Error(ErrEmptyUser)
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
			return
		}

		// get date.
		date := c.Param("date")
		if len(date) == 0 {
			logger.Error(ErrEmptyDate)
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
			return
		}

		// convert string to int.
		dateInt, err := strconv.Atoi(date)
		if err != nil {
			logger.Error(errors.New(fmt.Sprintf("convert fail : date(%s) string to int", date)))
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
			return
		}

		// get day time.
		dayTime, err := r.GetDayTime(ctx, id, dateInt)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": model.FailResponse})
			return
		}

		// convert int to string.
		dayTimeString := strconv.Itoa(dayTime)

		c.String(model.SuccessResponseCode, dayTimeString)

		return
	})

}
