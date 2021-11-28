package web

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

func (w *Web) SetDayTimeHandler(ctx context.Context) {

	timeInfo := &model.DayTimeInfo{}

	w.engine.POST("/time/day/set", func(c *gin.Context) {

		if err := c.BindJSON(timeInfo); err != nil {
			// request is not matched with day time info struct.
			logger.Error(errors.New(fmt.Sprintf("unmarshal failed : request info(%+v)"+
				" is not matched with timeInfo info struct(%v)", c.Request, timeInfo)))
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
		}

		// store day time.
		err := w.setTime(ctx, timeInfo.Id, timeInfo.Date, timeInfo.Time)
		if err != nil {
			logger.Error(err)
			c.JSON(model.FailResponseCode, gin.H{"error": err.Error()})
			return
		}

		c.JSON(model.SuccessResponseCode, gin.H{"error": err.Error()})
	})
}

func (w *Web) GetDayTimeHandler(ctx context.Context) {
	header := getDayTimeHeader{}

	w.engine.GET("/time/day/get/:user/:date", func(c *gin.Context) {

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
		dayTime, err := w.GetDayTime(ctx, header.ID, header.Date)
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

func (w *Web) setTime(ctx context.Context, id string, date, time int) error {

	prevDayTime, err := w.GetDayTime(ctx, id, date)
	if err != nil {
		logger.Error(err)
		return errors.New(fmt.Sprintf("get account(%s) date(%d) time(%d)", id, date, time))
	}

	if prevDayTime != 0 {
		// add prev time and current time
		time = prevDayTime + time
	}

	return w.db.SetDayTime(ctx, id, date, time)
}

func (w *Web) GetDayTime(ctx context.Context, id string, date int) (int, error) {
	return w.db.GetDayTime(ctx, id, date)
}
