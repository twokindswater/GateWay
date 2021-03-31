package logger

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func Init() error {
	logger = logrus.New()
	logger.SetReportCaller(true)
	return nil
}

func Error(err error) {
	logger.Error("ERROR:%e\n", err)
}

func Info(info string) {
	logger.Info(info)
}
