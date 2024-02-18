package app

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(config *logrus.TextFormatter) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(config) // Aquí se asume que config no es nil
	return logger
}
