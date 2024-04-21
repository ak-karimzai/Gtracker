// Package logger provides a logger implementation using Logrus.
package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a new logger instance. If a log file path is provided,
// the logs will be written to that file. Otherwise, logs will be written to stdout.
// It returns a Logger instance and an error if any.
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
