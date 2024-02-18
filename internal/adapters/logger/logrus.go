package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *logrus.TextFormatter) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetFormatter(config)
	return logger
}
