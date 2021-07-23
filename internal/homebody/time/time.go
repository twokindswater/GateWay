package time

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gateway/internal/homebody/db"
	"github.com/Gateway/internal/homebody/web"
	"github.com/Gateway/pkg/logger"
)

type Time struct {
	server *web.Web
	db     *db.DB
}

func Init(s *web.Web, db *db.DB) (*Time, error) {
	return &Time{
		server: s,
		db:     db,
	}, nil
}

func (t *Time) setTime(ctx context.Context, id string, date, time int) error {

	prevDayTime, err := t.GetDayTime(ctx, id, date)
	if err != nil {
		logger.Error(err)
		return errors.New(fmt.Sprintf("get account(%s) date(%d) time(%d)", id, date, time))
	}

	if prevDayTime != 0 {
		// add prev time and current time
		time = prevDayTime + time
	}

	return t.db.SetDayTime(ctx, id, date, time)
}

func (t *Time) GetDayTime(ctx context.Context, id string, date int) (int, error) {
	return t.db.GetDayTime(ctx, id, date)
}
