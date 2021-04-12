package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Init(ctx context.Context) error {
	logger = logrus.New()
	logger.SetReportCaller(false)
	return nil
}

func Error(err error) {
	logger.Error(err)
}

func Info(info string) {
	logger.Info(info)
}
