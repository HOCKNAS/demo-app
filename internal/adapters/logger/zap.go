package logger

import (
	"github.com/HOCKNAS/demo-app/internal/core/ports"
	zaplog "github.com/HOCKNAS/demo-app/pkg/log/zap_log"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(config string) ports.Logger {

	logger := zaplog.NewLogger([]byte(config))

	return &zapLogger{
		logger: logger,
	}
}

func (l *zapLogger) toZapFields(args ...interface{}) []zap.Field {
	if len(args)%2 != 0 {
		l.logger.Warn("Número impar de argumentos pasados a toZapFields", zap.Int("argsCount", len(args)))
	}

	var fields []zap.Field
	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			l.logger.Warn("Clave no es una cadena", zap.Any("key", args[i]))
			continue
		}
		fields = append(fields, zap.Any(key, args[i+1]))
	}
	return fields
}

func (l *zapLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, l.toZapFields(args...)...)
}

func (l *zapLogger) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, l.toZapFields(args...)...)
}

func (l *zapLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, l.toZapFields(args...)...)
}

func (l *zapLogger) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, l.toZapFields(args...)...)
}
