package db

import (
	"context"
	"fmt"
	"github.com/Gateway/internal/homebody/model"
	"github.com/Gateway/pkg/logger"
)

func (db *DB) SetDayTime(ctx context.Context, id string, date, time int) error {

	// encoding time.
	b, err := db.serializer.Encode(ctx, time)
	if err != nil {
		logger.Error(err)
		return errSerializer
	}

	// get account day time path.
	p := GetAccountDayTimePath(id, date)

	// set account day time.
	err = db.Client.Set(ctx, p, b)
	if err != nil {
		logger.Error(err)
		return errSetDataBase
	}

	return nil
}

func (db *DB) GetDayTime(ctx context.Context, id string, time int) (int, error) {

	// get account day time path.
	p := GetAccountDayTimePath(id, time)

	// get account day time.
	b, err := db.Client.Get(ctx, p)
	if err != nil {
		logger.Error(err)
		return 0, errGetDataBase
	}

	// has no data.
	if b == nil {
		return 0, nil
	}

	// initialize day time
	var minute int

	// decoding account day time
	err = db.serializer.Decode(ctx, b, &minute)
	if err != nil {
		logger.Error(err)
		return 0, errDecoding
	}

	return minute, nil
}

// h/t/{{id}}/{{date}}
func GetAccountDayTimePath(id string, date int) string {
	return fmt.Sprintf("%s%s%s%s%s%s%d", model.ServicePrefix, model.Delimiter, model.TimePrefix,
		model.Delimiter, id, model.Delimiter, date)
}
