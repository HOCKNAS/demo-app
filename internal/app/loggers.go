package app

import (
	"github.com/HOCKNAS/demo-app/internal/adapters/logger"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Users ports.Logger
}

func NewLoggers(config *logrus.TextFormatter) *Logger {
	return &Logger{
		Users: logger.NewLogrusLogger(config),
	}
}
