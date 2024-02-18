package app

import (
	"github.com/HOCKNAS/demo-app/internal/adapters/logger"
	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	UseLogger    string // "logrus" o "zap"
	LogrusConfig *logrus.TextFormatter
	ZapConfig    string
}

type Logger struct {
	Users ports.Logger
}

func NewLoggers(cfg LoggerConfig) *Logger {
	switch cfg.UseLogger {
	case "logrus":
		if cfg.LogrusConfig == nil {
			panic(domain.ErrLogrusConfigNotProvided.Error())
		}
		return &Logger{
			Users: logger.NewLogrusLogger(cfg.LogrusConfig),
		}
	case "zap":
		if cfg.ZapConfig == "" {
			panic(domain.ErrZapConfigNotProvided.Error())
		}
		return &Logger{
			Users: logger.NewZapLogger(cfg.ZapConfig),
		}
	default:
		panic("Configuración de logger no soportada")
	}
}
