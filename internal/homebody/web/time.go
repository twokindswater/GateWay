package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/pkg/logger"
)

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
