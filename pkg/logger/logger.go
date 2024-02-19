package log

import (
	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	logrus_log "github.com/HOCKNAS/demo-app/pkg/logger/logrus_log"
	zap_log "github.com/HOCKNAS/demo-app/pkg/logger/zap_log"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	UseLogger    string // "logrus" o "zap"
	LogrusConfig *logrus.TextFormatter
	ZapConfig    string
}

type Logger struct {
	Logger ports.Logger
}

func NewLogger(cfg LoggerConfig) *Logger {
	switch cfg.UseLogger {
	case "logrus":
		if cfg.LogrusConfig == nil {
			panic(domain.ErrLogrusConfigNotProvided.Error())
		}
		return &Logger{
			Logger: logrus_log.NewLogrusLogger(cfg.LogrusConfig),
		}
	case "zap":
		if cfg.ZapConfig == "" {
			panic(domain.ErrZapConfigNotProvided.Error())
		}
		return &Logger{
			Logger: zap_log.NewZapLogger(cfg.ZapConfig),
		}
	default:
		panic("Configuración de logger no soportada")
	}
}
