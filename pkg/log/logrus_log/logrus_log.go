package logruslog

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	return logrus.New()
}
