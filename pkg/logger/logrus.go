package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger(logFilePath ...string) (Logger, error) {
	var f *os.File
	var err error
	if len(logFilePath) > 0 {
		f, err = os.OpenFile(
			logFilePath[0],
			os.O_RDWR|os.O_CREATE|os.O_APPEND,
			0666,
		)
	} else {
		f = os.Stdout
	}
	if err != nil {
		return nil, fmt.Errorf(
			"err in file %s: %w", logFilePath, err)
	}

	logger := logrus.New()
	logger.SetOutput(f)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger, nil
}
