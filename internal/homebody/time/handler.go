package time

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/data"
	"github.com/Gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrEmptyUser = errors.New("empty user name")
	ErrEmptyDate = errors.New("empty date")
)

func (t *Time) AddHandler(ctx context.Context) {
	t.SetDayTimeHandler(ctx)
	t.GetDayTimeHandler(ctx)
}

func (t *Time) SetDayTimeHandler(ctx context.Context) {

	timeInfo := &data.DayTimeInfo{}

	t.server.Client.Router.POST("/time/day/set", func(c *gin.Context) {

		if err := c.BindJSON(timeInfo); err != nil {
			// request is not matched with day time info struct.
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with kakao info struct", c.Request)))
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
		}

		// store day time.
		err := t.SetDayTime(ctx, timeInfo.Id, timeInfo.Date, timeInfo.Time)
		if err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		c.JSON(data.SuccessResponseCode, gin.H{"error": data.FailResponse})
	})
}

func (t *Time) GetDayTimeHandler(ctx context.Context) {

	t.server.Client.Router.GET("/time/day/get/:user/:date", func(c *gin.Context) {

		// get user id.
		id := c.Param("user")
		if len(id) == 0 {
			logger.Error(ErrEmptyUser)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		// get date.
		date := c.Param("date")
		if len(date) == 0 {
			logger.Error(ErrEmptyDate)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		// convert string to int.
		dateInt, err := strconv.Atoi(date)
		if err != nil {
			logger.Error(errors.New(fmt.Sprintf("convert fail : date(%s) string to int", date)))
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		// get day time.
		dayTime, err := t.GetDayTime(ctx, id, dateInt)
		if err != nil {
			logger.Error(err)
			c.JSON(data.FailResponseCode, gin.H{"error": data.FailResponse})
			return
		}

		// convert int to string.
		dayTimeString := strconv.Itoa(dayTime)

		c.String(data.SuccessResponseCode, dayTimeString)

		return
	})

}
